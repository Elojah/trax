import { createApp } from 'vue'
import { createPinia } from 'pinia'
import './index.css'

import PrimeVue from 'primevue/config';
import Aura from '@primevue/themes/aura';
import 'primeicons/primeicons.css';

import GoogleSignInPlugin from "vue3-google-signin"

import App from './App.vue'
import router from './router'
import { config } from './config'

const app = createApp(App)

app.use(createPinia())

app.use(GoogleSignInPlugin, {
  clientId: config.google_client_id,
});

app.use(PrimeVue, {
  ripple: true,
  theme: {
    preset: Aura
  }
});

app.use(router)

app.mount('#app')
