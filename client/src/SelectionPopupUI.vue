<script setup>
const props = defineProps(['visible', 'popupStyle', 'onClickHandler'])
const { visible = false, popupStyle, onClickHandler } = props

</script>

<template>
    <div v-if="visible" :style="popupStyle">
        <button @click="onClickHandler">Share</button>
    </div>
</template>
