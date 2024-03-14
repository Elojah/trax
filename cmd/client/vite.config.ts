import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vuetify from 'vite-plugin-vuetify'
import vueJsx from '@vitejs/plugin-vue-jsx'
import fs from 'fs';

export default async() =>{
  let config = {}
  const loadConfigPlugin = {
    name: 'load-config',
    buildStart() {
      const configFile = process.env.CONFIG_FILE;

      if (!configFile) {
        throw new Error('CONFIG_FILE not set');
      }

      const configData = fs.readFileSync(configFile, 'utf-8');

      if (!configData) {
        throw new Error('CONFIG_FILE is empty');
      }

      config = JSON.parse(configData);
    },
    transform(code:any) {
      return code.replace('__CONFIG__', JSON.stringify(config));
    }
  }

  return defineConfig({
    plugins: [
      vue(),
      vueJsx(),
      vuetify(),
      loadConfigPlugin,
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
}
