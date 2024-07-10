<script setup>

import { ref } from 'vue';
import ShareModal from './ShareModal.vue';

const props = defineProps(['visible', 'quoteText', 'popupStyle', 'onClickHandler'])
const { visible = false, quoteText, popupStyle, onClickHandler } = props

const modalOpen = ref(false);
const quoteAuthor = ref("Vivianne Nguyen");
const quoteSource = ref("Nam thời đại của loài người");

const openModal = () => {
    modalOpen.value = true;
};

const closeModal = () => {
    modalOpen.value = false;
};

</script>

<template>
    <div v-if="visible && !modalOpen" :style="popupStyle">
        <button @click="openModal">Share</button>
    </div>
    <ShareModal :isOpen="modalOpen" :quote="quoteText" :author="quoteAuthor" :source="quoteSource"
        @close="closeModal" />
</template>
