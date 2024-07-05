FROM scratch

COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY fonts fonts
COPY quote-share /

ENTRYPOINT ["/quote-share"]
