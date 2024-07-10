<script setup>

import { ref } from 'vue';
import ShareModal from './ShareModal.vue';

const props = defineProps(['visible', 'popupStyle', 'onClickHandler'])
const { visible = false, popupStyle, onClickHandler } = props

const modalOpen = ref(false);
const quoteText = ref("Tín ngưỡng Hy Lạp cổ đại cho rằng xuất phát điểm của Thời đại Vàng là ở vùng đất Arcadia, một vùng nông thôn nghèo khó nhưng mang vẻ đẹp yên bình. Nơi này vốn là quê nhà của thần mục đồng Pan.");
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
