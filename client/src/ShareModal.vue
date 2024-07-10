<template>
  <Teleport to="body">
    <div v-if="isOpen" class="modal-overlay" @click="closeModal" @keydown.esc="closeModal" tabindex="0">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h2>Share quote</h2>
          <button @click="closeModal" class="close-button">&times;</button>
        </div>
        <div class="quote-container">
          <img
            src="https://quote.vinhis.me/v1/quote?text=xu%E1%BA%A5t%20s%E1%BA%AFc&image=https://cdn.brvn.vn/news/1280px/2024/LG-MONQ-COVER_1720517326.jpg"
            alt="" class="quote-image">
          <blockquote>
            {{ quote }}
          </blockquote>
          <p class="author">{{ author }}</p>
          <p class="source">{{ source }}</p>
        </div>
        <div class="modal-footer">
          <div class="background-options">
            <span>Background</span>
            <div class="color-options">
              <button v-for="color in backgroundColors" :key="color"
                :class="['color-option', { active: selectedBackground === color }]" :style="{ backgroundColor: color }"
                @click="selectBackground(color)">
              </button>
            </div>
          </div>
          <div class="aspect-ratio">
            <span>Aspect ratio</span>
            <div class="ratio-options">
              <button :class="['ratio-option', { active: selectedRatio === 'vertical' }]"
                @click="selectRatio('vertical')">
                Vertical
              </button>
              <button :class="['ratio-option', { active: selectedRatio === 'square' }]" @click="selectRatio('square')">
                Square
              </button>
            </div>
          </div>
          <div class="share-options">
            <button v-for="option in shareOptions" :key="option" class="share-option">
              <img :src="getShareIcon(option)" :alt="option">
            </button>
          </div>
          <button class="download-button">
            <!-- <img src="path-to-download-icon.svg" alt="Download"> -->
            Download
          </button>
        </div>
      </div>
    </div>
  </Teleport>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';

const props = defineProps({
  isOpen: {
    type: Boolean,
    required: true,
  },
  quote: {
    type: String,
    required: true,
  },
  author: {
    type: String,
    required: true,
  },
  source: {
    type: String,
    required: true,
  },
});

const emit = defineEmits(['close']);

const backgroundColors = ['#808080', '#8A2BE2', '#FFFFFF'];
const selectedBackground = ref(backgroundColors[0]);
const selectedRatio = ref('vertical');
const shareOptions = ['link', 'twitter', 'facebook', 'linkedin'];

const selectBackground = (color) => {
  selectedBackground.value = color;
};

const selectRatio = (ratio) => {
  selectedRatio.value = ratio;
};

const closeModal = () => {
  emit('close');
};

const handleEscape = (e) => {
  if (e.key === 'Escape' && props.isOpen) {
    closeModal();
  }
};

const getShareIcon = (option) => {
  return `path-to-${option}-icon.svg`;
};

onMounted(() => {
  document.addEventListener('keydown', handleEscape);
});

onUnmounted(() => {
  document.removeEventListener('keydown', handleEscape);
});
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  overflow-y: auto;
}

.modal-content {
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  max-width: 50%;
  max-height: 90%;
  overflow-y: auto;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.close-button {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
}

.quote-container {
  margin: 20px 0;
}

.modal-footer {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.background-options,
.aspect-ratio {
  display: flex;
  align-items: center;
  gap: 10px;
}

.color-options,
.ratio-options {
  display: flex;
  gap: 5px;
}

.color-option,
.ratio-option {
  width: 30px;
  height: 30px;
  border: 1px solid #ccc;
  border-radius: 50%;
  cursor: pointer;
}

.ratio-option {
  border-radius: 4px;
  padding: 5px 10px;
}

.active {
  border: 2px solid #000;
}

.share-options {
  display: flex;
  gap: 10px;
}

.share-option {
  background: none;
  border: none;
  cursor: pointer;
}

.download-button {
  background-color: #FF6600;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 5px;
}

.quote-image {
  width: 100%;
  height: auto;
}
</style>
