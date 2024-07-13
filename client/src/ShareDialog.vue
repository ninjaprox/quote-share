<script setup>
import { Icon } from "@iconify/vue";
import {
  DialogClose,
  DialogContent,
  DialogOverlay,
  DialogPortal,
  DialogRoot,
  DialogTitle,
  DialogTrigger,
} from "radix-vue";
import { computed } from "vue";

// eslint-disable-next-line vue/require-prop-types
const props = defineProps(["open", "image", "quote", "triggerStyle"]);

const imageSrc = computed(() => {
  const { image, quote } = props;
  const imageURL = new URL("http://localhost:8080/v1/quote");

  imageURL.searchParams.set("text", quote);
  imageURL.searchParams.set("image", image);

  return imageURL.href;
});
const downloadHref = computed(() => {
  const { image, quote } = props;
  const imageURL = new URL("http://localhost:8080/v1/quote");

  imageURL.searchParams.set("text", quote);
  imageURL.searchParams.set("image", image);
  imageURL.searchParams.set("download", "true");

  return imageURL.href;
});
</script>

<template>
  <DialogRoot>
    <DialogTrigger
      class="border p-1 px-2 bg-white rounded text-grass11 border-grass11"
      :style="triggerStyle"
      >Share</DialogTrigger
    >
    <DialogPortal>
      <DialogOverlay
        class="bg-blackA9 data-[state=open]:animate-overlayShow fixed inset-0 z-30"
      />
      <DialogContent
        class="data-[state=open]:animate-contentShow fixed top-[50%] left-[50%] max-h-[85vh] w-[90vw] max-w-[450px] translate-x-[-50%] translate-y-[-50%] rounded-[6px] bg-white p-[25px] shadow-[hsl(206_22%_7%_/_35%)_0px_10px_38px_-10px,_hsl(206_22%_7%_/_20%)_0px_10px_20px_-15px] focus:outline-none z-[100]"
      >
        <DialogTitle class="text-mauve12 m-0 text-[17px] font-semibold"
          >Share quote</DialogTitle
        >
        <div class="relative group">
          <a :href="downloadHref" download>
            <img
              :src="imageSrc"
              alt="quote image"
              class="w-full h-auto my-2.5"
              width="1000px"
              height="1000px"
            />
          </a>
          <Icon
            icon="lucide:download"
            class="w-8 h-8 absolute bottom-0 right-0 text-white m-2 hidden group-hover:block"
          />
        </div>
        <DialogClose
          class="text-grass11 hover:bg-green4 focus:shadow-green7 absolute top-[10px] right-[10px] inline-flex h-[25px] w-[25px] appearance-none items-center justify-center rounded-full focus:shadow-[0_0_0_2px] focus:outline-none"
          aria-label="Close"
        >
          <Icon icon="lucide:x" />
        </DialogClose>
      </DialogContent>
    </DialogPortal>
  </DialogRoot>
</template>

<style scoped></style>
