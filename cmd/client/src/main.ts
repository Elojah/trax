import { createApp } from 'vue'
import { createPinia } from 'pinia'

import ElementPlus from 'element-plus'

import GoogleSignInPlugin from "vue3-google-signin"

import App from './App.vue'
import router from './router'
import { config } from './config'

const app = createApp(App)

app.use(createPinia())

app.use(GoogleSignInPlugin, {
  clientId: config.google_client_id,
});

app.use(ElementPlus)
// app.use(PrimeVue, {
//   theme: {
//     preset: Aura,
//     options: {
//       darkModeSelector: ".p-dark",
//     },
//     ripple: true,
//   },
// });

app.use(router)

app.mount('#app')
