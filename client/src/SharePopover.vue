<script setup>
import "./style.css";

import ShareDialog from "./ShareDialog.vue";

// eslint-disable-next-line vue/require-prop-types
const props = defineProps(["open", "imageSrc", "quoteText", "popoverStyle"]);
const { imageSrc, quoteText, popoverStyle } = props;
</script>

<template>
  <ShareDialog
    :image="imageSrc"
    :quote="quoteText"
    :trigger-style="popoverStyle"
  />
</template>
