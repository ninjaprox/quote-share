import vue from "@vitejs/plugin-vue";
import { resolve } from "path";
import { defineConfig } from "vite";

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    rollupOptions: {
      input: {
        main: resolve(__dirname, "index.html"),
        library: resolve(__dirname, "src/QuoteShare.js"),
      },
    },
  },
  plugins: [vue()],
});
