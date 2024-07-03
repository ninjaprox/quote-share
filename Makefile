# Variables
DOCKER_REPO := vinhis.me
IMAGE_NAME := quote-share
VERSION := 0.1.0

# Supported architectures for multi-arch build
ARCHITECTURES := amd64 arm64

# Docker buildx builder name
BUILDER_NAME := multiarch-builder

# Get the current architecture
CURRENT_ARCH := $(shell uname -m | sed 's/x86_64/amd64/' | sed 's/aarch64/arm64/')

.PHONY: all build build-local push clean

# Default target: build for current architecture
all: build-local

# Build for current architecture
build-local:
	docker build -t $(DOCKER_REPO)/$(IMAGE_NAME):$(VERSION)-$(CURRENT_ARCH) .

# Create a multi-architecture builder if it doesn't exist
create-builder:
	docker buildx create --name $(BUILDER_NAME) --use || true

# Build multi-architecture Docker images
build-multi: create-builder
	docker buildx build --platform $(shell echo $(ARCHITECTURES) | sed 's/ /,/g') \
		-t $(DOCKER_REPO)/$(IMAGE_NAME):$(VERSION) \
		--push \
		.

# Push the multi-architecture Docker images
push: build-multi

# Remove the Docker buildx builder
clean:
	docker buildx rm $(BUILDER_NAME) || true
