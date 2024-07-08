import debounce from 'lodash/debounce'
class SelectionPopup {
    constructor(options = {}) {
        this.options = {
            buttonTexts: ['Button 1', 'Button 2', 'Button 3'],
            buttonCallbacks: [() => { }, () => { }, () => { }],
            popupClass: 'selection-popup',
            debounceTime: 500, // Debounce time in milliseconds
            ...options
        };
        this.popup = null;
        this.init();
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
        this.popup = document.createElement('div');
        this.popup.className = this.options.popupClass;
        this.popup.style.display = 'none';
        this.popup.style.position = 'absolute';
        this.popup.style.zIndex = '1000';

        this.options.buttonTexts.forEach((text, index) => {
            const button = document.createElement('button');
            button.textContent = text;
            button.addEventListener('click', () => {
                this.options.buttonCallbacks[index]();
                this.hidePopup();
            });
            this.popup.appendChild(button);
        });

        document.body.appendChild(this.popup);
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
        const range = selection.getRangeAt(0);
        const rect = range.getBoundingClientRect();

        this.popup.style.left = `${rect.left + window.scrollX}px`;
        this.popup.style.top = `${rect.bottom + window.scrollY}px`;
        this.popup.style.display = 'block';
    }

    hidePopup() {
        this.popup.style.display = 'none';
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
