import debounce from 'lodash/debounce';
import { createApp, ref } from 'vue'
import SelectionPopupUI from './SelectionPopupUI.vue'
class SelectionPopup {
    constructor(options = {}) {
        this.options = {
            buttonTexts: ['Button 1', 'Button 2', 'Button 3'],
            buttonCallbacks: [() => { }, () => { }, () => { }],
            popupClass: 'selection-popup',
            debounceTime: 500, // Debounce time in milliseconds
            ...options
        };
        this.isVisible = ref(false)
        this.quoteText = ref('')
        this.popupStyle = {
            position: 'absolute',
            zIndex: '1000',
            left: '0px',
            top: '0px',
        }
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
        this.vueApp = createApp(SelectionPopupUI, {
            quoteText: this.quoteText,
            visible: this.isVisible,
            popupStyle: this.popupStyle,
            onClickHandler: () => console.log('hello')
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
        const range = selection.getRangeAt(0)
        const rect = range.getBoundingClientRect()

        this.popupStyle.left = `${rect.left + window.scrollX}px`
        this.popupStyle.top = `${rect.bottom + window.scrollY}px`
        this.quoteText.value = selection.toString().trim()
        this.isVisible.value = true
    }

    hidePopup() {
        this.isVisible.value = false
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
