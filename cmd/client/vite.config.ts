import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueJsx(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
      '@internal': fileURLToPath(new URL('./../../internal', import.meta.url)),
      '@pkg': fileURLToPath(new URL('./../../pkg', import.meta.url)),
      '@api': fileURLToPath(new URL('./../../cmd/api/grpc', import.meta.url)),
      // imports from external files
      '@protobuf-ts': fileURLToPath(new URL('./node_modules/@protobuf-ts', import.meta.url)),
    }
  }
})
