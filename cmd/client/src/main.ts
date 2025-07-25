import { createApp } from 'vue'
import { createPinia } from 'pinia'
import './index.css'

import PrimeVue from 'primevue/config';
import { ConfirmationService, ToastService } from 'primevue';
import Aura from '@primevue/themes/aura';
import 'primeicons/primeicons.css';

import GoogleSignInPlugin from "vue3-google-signin"

import App from './App.vue'
import router from './router'
import { config } from './config'
import { definePreset } from '@primevue/themes';

const app = createApp(App)

app.use(createPinia())

app.use(GoogleSignInPlugin, {
  clientId: config.google_client_id,
});


const preset = definePreset(Aura, {
  semantic: {
    primary: {
      50: '{zinc.50}',
      100: '{zinc.100}',
      200: '{zinc.200}',
      300: '{zinc.300}',
      400: '{zinc.400}',
      500: '{zinc.500}',
      600: '{zinc.600}',
      700: '{zinc.700}',
      800: '{zinc.800}',
      900: '{zinc.900}',
      950: '{zinc.950}'
    },
    secondary: {
      50: '{slate.50}',
      100: '{slate.100}',
      200: '{slate.200}',
      300: '{slate.300}',
      400: '{slate.400}',
      500: '{slate.500}',
      600: '{slate.600}',
      700: '{slate.700}',
      800: '{slate.800}',
      900: '{slate.900}',
      950: '{slate.950}'
    }
  }
});



app.use(PrimeVue, {
  theme: {
    preset: preset,
  },
  options: {
    darkModeSelector: 'system',
  },
});
app.use(ConfirmationService);
app.use(ToastService);

app.use(router)

app.mount('#app')
