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
  RadioGroupIndicator,
  RadioGroupItem,
  RadioGroupRoot,
} from "radix-vue";
import { computed, ref } from "vue";

// eslint-disable-next-line vue/require-prop-types
const props = defineProps(["open", "image", "quote", "triggerStyle"]);
const background = ref("image");
const size = ref("square");

const baseURL = import.meta.env.DEV
  ? new URL("/v1/quote", "http://localhost:8080")
  : new URL("/v1/quote", "https://quote.vinhis.me");

const imageParam = computed(() => {
  const { image } = props;

  return background.value === "image" ? image : background.value;
});
const sizeParam = computed(() =>
  size.value === "rectangle" ? "portrait" : "square"
);
const imageSrc = computed(() => {
  const { quote } = props;
  const imageURL = new URL(baseURL);

  imageURL.searchParams.set("text", quote);
  imageURL.searchParams.set("image", imageParam.value);
  imageURL.searchParams.set("size", sizeParam.value);

  return imageURL.href;
});
const downloadHref = computed(() => {
  const { quote } = props;
  const downloadURL = new URL(baseURL);

  downloadURL.searchParams.set("text", quote);
  downloadURL.searchParams.set("image", imageParam.value);
  downloadURL.searchParams.set("size", sizeParam.value);
  downloadURL.searchParams.set("download", "true");

  return downloadURL.href;
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
        class="overflow-auto data-[state=open]:animate-contentShow fixed top-[50%] left-[50%] max-h-[85vh] w-[90vw] max-w-[450px] translate-x-[-50%] translate-y-[-50%] rounded-[6px] bg-white p-[25px] shadow-[hsl(206_22%_7%_/_35%)_0px_10px_38px_-10px,_hsl(206_22%_7%_/_20%)_0px_10px_20px_-15px] focus:outline-none z-[100]"
      >
        <DialogTitle class="text-mauve12 m-0 text-[17px] font-semibold"
          >Share quote</DialogTitle
        >
        <div class="relative group mb-4">
          <a :href="downloadHref" download>
            <img
              :src="imageSrc"
              alt="quote image"
              class="w-full h-auto my-2.5 drop-shadow rounded border"
              width="1000px"
              height="1000px"
            />
          </a>
          <Icon
            icon="lucide:download"
            class="w-8 h-8 absolute bottom-0 right-0 m-2 hidden group-hover:block"
            :class="[background === 'white' ? 'text-black' : 'text-white']"
          />
        </div>
        <div class="mb-2">
          Background
          <RadioGroupRoot
            v-model="background"
            class="flex flex-row gap-2.5"
            default-value="image"
            aria-label="Background"
            orientation="horizontal"
          >
            <div class="flex items-center">
              <RadioGroupItem
                id="background-image"
                class="bg-white w-[25px] h-[25px] rounded-full shadow-[0_2px_10px] shadow-blackA7 hover:bg-green3 focus:shadow-[0_0_0_2px] focus:shadow-black outline-none cursor-default"
                value="image"
              >
                <RadioGroupIndicator
                  class="flex items-center justify-center w-full h-full relative after:content-[''] after:block after:w-[11px] after:h-[11px] after:rounded-[50%] after:bg-grass11"
                />
              </RadioGroupItem>
              <label
                class="text-[15px] leading-none pl-[15px]"
                for="background-image"
              >
                Image
              </label>
            </div>
            <div class="flex items-center">
              <RadioGroupItem
                id="background-red"
                class="bg-white w-[25px] h-[25px] rounded-full shadow-[0_2px_10px] shadow-blackA7 hover:bg-green3 focus:shadow-[0_0_0_2px] focus:shadow-black outline-none cursor-default"
                value="red"
              >
                <RadioGroupIndicator
                  class="flex items-center justify-center w-full h-full relative after:content-[''] after:block after:w-[11px] after:h-[11px] after:rounded-[50%] after:bg-grass11"
                />
              </RadioGroupItem>
              <label
                class="text-[15px] leading-none pl-[15px]"
                for="background-red"
              >
                Red
              </label>
            </div>
            <div class="flex items-center">
              <RadioGroupItem
                id="background-white"
                class="bg-white w-[25px] h-[25px] rounded-full shadow-[0_2px_10px] shadow-blackA7 hover:bg-green3 focus:shadow-[0_0_0_2px] focus:shadow-black outline-none cursor-default"
                value="white"
              >
                <RadioGroupIndicator
                  class="flex items-center justify-center w-full h-full relative after:content-[''] after:block after:w-[11px] after:h-[11px] after:rounded-[50%] after:bg-grass11"
                />
              </RadioGroupItem>
              <label
                class="text-[15px] leading-none pl-[15px]"
                for="background-white"
              >
                White
              </label>
            </div>
          </RadioGroupRoot>
        </div>
        <div>
          Aspect radio
          <RadioGroupRoot
            v-model="size"
            class="flex flex-row gap-2.5"
            default-value="image"
            aria-label="Aspect ratio"
            orientation="horizontal"
          >
            <div class="flex items-center">
              <RadioGroupItem
                id="size-square"
                class="bg-white w-[25px] h-[25px] rounded-full shadow-[0_2px_10px] shadow-blackA7 hover:bg-green3 focus:shadow-[0_0_0_2px] focus:shadow-black outline-none cursor-default"
                value="square"
              >
                <RadioGroupIndicator
                  class="flex items-center justify-center w-full h-full relative after:content-[''] after:block after:w-[11px] after:h-[11px] after:rounded-[50%] after:bg-grass11"
                />
              </RadioGroupItem>
              <label
                class="text-[15px] leading-none pl-[15px]"
                for="size-square"
              >
                Square
              </label>
            </div>
            <div class="flex items-center">
              <RadioGroupItem
                id="size-rectangle"
                class="bg-white w-[25px] h-[25px] rounded-full shadow-[0_2px_10px] shadow-blackA7 hover:bg-green3 focus:shadow-[0_0_0_2px] focus:shadow-black outline-none cursor-default"
                value="rectangle"
              >
                <RadioGroupIndicator
                  class="flex items-center justify-center w-full h-full relative after:content-[''] after:block after:w-[11px] after:h-[11px] after:rounded-[50%] after:bg-grass11"
                />
              </RadioGroupItem>
              <label
                class="text-[15px] leading-none pl-[15px]"
                for="size-rectangle"
              >
                Rectangle
              </label>
            </div>
          </RadioGroupRoot>
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
