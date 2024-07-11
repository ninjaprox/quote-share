import debounce from 'lodash/debounce';
import { createApp, ref } from 'vue'
import SharePopover from './SharePopover.vue'
class SelectionPopup {
    constructor(options = {}) {
        this.options = {
            debounceTime: 500,
            ...options
        };
        this.imageSrc = ref('')
        this.quoteText = ref('')
        this.isOpen = ref(false)
        this.popoverStyle = ref({
            position: 'absolute',
            zIndex: '1000',
            left: '0px',
            top: '0px',
            display: 'none',
        })
        this.vueApp = null
    }

    init() {
        this.createPopup();
        this.debouncedShowPopup = debounce(
            this.showPopup.bind(this),
            this.options.debounceTime
        );
        document.addEventListener('selectionchange', () => this.handleSelectionChange());
    }

    createPopup() {
        const popupContainer = document.createElement('div')

        document.body.appendChild(popupContainer)
        this.vueApp = createApp(SharePopover, {
            open: this.isOpen,
            imageSrc: this.imageSrc,
            quoteText: this.quoteText,
            popoverStyle: this.popoverStyle,
        })
        this.vueApp.mount(popupContainer)
    }

    handleSelectionChange() {
        const selection = window.getSelection();

        if (selection.toString().trim().length > 0) {
            this.debouncedShowPopup(selection);
        } else {
            this.hidePopup();
        }
    }

    showPopup(selection) {
        const imageSrc = document.querySelector(this.options.imageSelector)?.src
        const range = selection.getRangeAt(0)
        const rect = range.getBoundingClientRect()

        this.popoverStyle.value = {
            ...this.popoverStyle.value,
            left: `${rect.left + window.scrollX}px`,
            top: `${rect.bottom + window.scrollY}px`,
            display: 'block',
        }
        this.imageSrc.value = imageSrc
        this.quoteText.value = selection.toString().trim()
        this.isOpen.value = true
    }

    hidePopup() {
        this.popoverStyle.value = {
            ...this.popoverStyle.value,
            display: 'none',
        }
        this.isOpen.value = false
    }
}

// Export the library
if (typeof module !== 'undefined' && typeof module.exports !== 'undefined') {
    module.exports = SelectionPopup;
} else if (typeof exports !== 'undefined') {
    exports.default = SelectionPopup;
} else {
    window.SelectionPopup = SelectionPopup;
}

export default SelectionPopup;
