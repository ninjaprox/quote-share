<script setup>
import { ref } from "vue";
import ShareDialog from "./ShareDialog.vue";

// eslint-disable-next-line vue/require-prop-types
const props = defineProps(["open", "imageSrc", "quoteText", "popoverStyle"]);
const { open = false, imageSrc, quoteText, popoverStyle } = props;
</script>

<template>
  <ShareDialog
    :open
    :image="imageSrc"
    :quote="quoteText"
    :trigger-style="popoverStyle"
  />
</template>
