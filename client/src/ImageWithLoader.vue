<script setup>
import { Icon } from "@iconify/vue";
import { computed, ref, watch } from "vue";

// eslint-disable-next-line vue/require-prop-types
const props = defineProps(["src", "size"]);
const isLoading = ref(true);

watch(
  () => props.src,
  () => {
    isLoading.value = true;
  }
);

const intrinsicSize = computed(() =>
  props.size === "rectangle"
    ? {
        width: "1080px",
        height: "1350px",
      }
    : {
        width: "1080px",
        height: "1080px",
      }
);
const handleImageLoad = () => {
  isLoading.value = false;
};
</script>

<template>
  <div class="relative">
    <img
      :src="props.src"
      alt="quote image"
      class="w-full h-auto my-2.5 drop-shadow rounded border"
      :width="intrinsicSize.width"
      :height="intrinsicSize.height"
      @load="handleImageLoad"
    />
    <Icon
      v-if="isLoading"
      class="absolute top-[50%] left-[50%] translate-x-[-50%] translate-y-[-50%]"
      icon="svg-spinners:blocks-wave"
      width="48"
      height="48"
    />
  </div>
</template>
