import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { quasar, transformAssetUrls } from '@quasar/vite-plugin'
import path from 'path'
// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@': path.resolve(__dirname, "src"),
      'vue-i18n': 'vue-i18n/dist/vue-i18n.cjs.js'
    },
  },
  server: {
    proxy: {
      '/m': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/m/, ''),
      }
    }
  },
  plugins: [
    vue({ template: { transformAssetUrls } }),
    quasar({
      sassVariables: 'src/quasar-variables.sass'
    })],
})
