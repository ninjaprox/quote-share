<script setup>
import { ref, computed, onMounted, onUnmounted } from "vue";
import {
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogOverlay,
  DialogPortal,
  DialogRoot,
  DialogTitle,
  DialogTrigger,
} from "radix-vue";

const props = defineProps(["open", "image", "quote", "triggerStyle"]);

const imageSrc = computed(() => {
  const { image, quote } = props;
  const imageURL = new URL("http://localhost:8080/v1/quote");

  imageURL.searchParams.set("text", quote);
  imageURL.searchParams.set("image", image);

  return imageURL.href;
});
</script>

<template>
  <DialogRoot>
    <DialogTrigger :style="triggerStyle">Share</DialogTrigger>
    <DialogPortal>
      <DialogOverlay class="overlay" />
      <DialogContent class="dialog-content">
        <DialogTitle>Share quote</DialogTitle>
        <DialogDescription>Description</DialogDescription>
        <img :src="imageSrc" alt="" class="quote-image" />
        <blockquote>
          {{ quote }}
        </blockquote>
        <DialogClose>Close</DialogClose>
      </DialogContent>
    </DialogPortal>
  </DialogRoot>
</template>

<style scoped>
.dialog-content {
  position: fixed;
  top: 0;
  left: 0;
  /* background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  overflow-y: auto; */

  background-color: white;
  padding: 20px;
  border-radius: 8px;
  max-width: 50%;
  max-height: 90%;
  overflow-y: auto;
}

.quote-image {
  width: 100%;
  height: auto;
}
</style>
