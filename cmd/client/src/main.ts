import { createApp } from 'vue'
import { createPinia } from 'pinia'
import './index.css'

import PrimeVue from 'primevue/config';
import { ConfirmationService, ToastService } from 'primevue';
import Aura from '@primevue/themes/aura';
import 'primeicons/primeicons.css';

import GoogleSignInPlugin from "vue3-google-signin"

import App from '@/App.vue'
import router from '@/router'
import { config } from '@/config'
import { definePreset } from '@primevue/themes';

const app = createApp(App)

app.use(createPinia())

app.use(GoogleSignInPlugin, {
  clientId: config.google_client_id,
});


const preset = definePreset(Aura, {
  semantic: {
    primary: {
      50: '{amber.50}',
      100: '{amber.100}',
      200: '{amber.200}',
      300: '{amber.300}',
      400: '{amber.400}',
      500: '{amber.500}',
      600: '{amber.600}',
      700: '{amber.700}',
      800: '{amber.800}',
      900: '{amber.900}',
      950: '{amber.950}'
    },
    secondary: {
      50: '{sky.50}',
      100: '{sky.100}',
      200: '{sky.200}',
      300: '{sky.300}',
      400: '{sky.400}',
      500: '{sky.500}',
      600: '{sky.600}',
      700: '{sky.700}',
      800: '{sky.800}',
      900: '{sky.900}',
      950: '{sky.950}'
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
