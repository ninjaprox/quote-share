import debounce from "lodash/debounce";
import { createApp, ref } from "vue";
import SharePopover from "./SharePopover.vue";

class QuoteShare {
  constructor(options = {}) {
    this.options = {
      debounceTime: 500,
      ...options,
    };
    this.isEnabled = true;
    this.imageSrc = ref("");
    this.quoteText = ref("");
    this.popoverStyle = ref({
      position: "absolute",
      zIndex: "20",
      left: "0px",
      top: "0px",
      display: "none",
    });
    this.vueApp = null;
  }

  init() {
    const { debounceTime } = this.options;

    this.createPopover();
    this.debouncedShowPopover = debounce(
      this.showPopover.bind(this),
      debounceTime
    );
    document.addEventListener(
      "selectionchange",
      this.handleSelectionChange.bind(this)
    );
  }

  createPopover() {
    const popoverContainer = document.createElement("div");

    document.body.appendChild(popoverContainer);
    this.vueApp = createApp(SharePopover, {
      imageSrc: this.imageSrc,
      quoteText: this.quoteText,
      popoverStyle: this.popoverStyle,
    });
    this.vueApp.mount(popoverContainer);
  }

  handleSelectionChange() {
    const selection = window.getSelection();

    if (selection.toString().trim().length === 0) {
      this.hidePopover();

      return;
    }

    if (this.isEnabled && this.checkSelection()) {
      this.debouncedShowPopover(selection);
    }
  }

  checkSelection() {
    const { contentSelector } = this.options;
    const selection = window.getSelection();
    const contentContainer = document.querySelector(contentSelector);

    return contentContainer.contains(selection.anchorNode.parentElement);
  }

  showPopover(selection) {
    const imageSrc = document.querySelector(this.options.imageSelector)?.src;
    const range = selection.getRangeAt(0);
    const rect = range.getBoundingClientRect();

    this.isEnabled = false;
    this.popoverStyle.value = {
      ...this.popoverStyle.value,
      left: `${rect.left + window.scrollX}px`,
      top: `${rect.bottom + window.scrollY}px`,
      display: "block",
    };
    this.imageSrc.value = imageSrc;
    this.quoteText.value = selection.toString().trim();
  }

  hidePopover() {
    this.isEnabled = true;
    this.popoverStyle.value = {
      ...this.popoverStyle.value,
      display: "none",
    };
  }
}

// Export the library
if (typeof module !== "undefined" && typeof module.exports !== "undefined") {
  module.exports = QuoteShare;
} else if (typeof exports !== "undefined") {
  exports.default = QuoteShare;
} else {
  window.QuoteShare = QuoteShare;
}

export default QuoteShare;
