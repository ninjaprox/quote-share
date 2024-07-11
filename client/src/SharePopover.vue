<script setup>
import { ref } from "vue";
import ShareDialog from "./ShareDialog.vue";

const props = defineProps(["open", "imageSrc", "quoteText", "popoverStyle"]);
const { open = false, imageSrc, quoteText, popoverStyle } = props;
</script>

<template>
  <ShareDialog
    :open
    :image="imageSrc"
    :quote="quoteText"
    :triggerStyle="popoverStyle"
  />
</template>
