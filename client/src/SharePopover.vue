<script setup>
import { ref } from "vue";
import {
  PopoverAnchor,
  PopoverArrow,
  PopoverClose,
  PopoverContent,
  PopoverPortal,
  PopoverRoot,
  PopoverTrigger,
} from "radix-vue";
import ShareModal from "./ShareModal.vue";

const props = defineProps(["open", "imageSrc", "quoteText", "popupStyle"]);
const { open = false, imageSrc, quoteText, popupStyle } = props;

const modalOpen = ref(false);

const openModal = () => {
  modalOpen.value = true;
};

const closeModal = () => {
  modalOpen.value = false;
};

console.log({ popupStyle });
</script>

<template>
  <PopoverRoot :open="open && !modalOpen">
    <PopoverAnchor :style="popupStyle" />
    <PopoverPortal>
      <PopoverContent>
        <button @click="openModal">Share</button>
      </PopoverContent>
    </PopoverPortal>
  </PopoverRoot>
  <ShareModal
    :isOpen="modalOpen"
    :image="imageSrc"
    :quote="quoteText"
    @close="closeModal"
  />
</template>
