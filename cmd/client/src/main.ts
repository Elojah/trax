import './assets/main.css'
import '@mdi/font/css/materialdesignicons.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import GoogleSignInPlugin from "vue3-google-signin"

import App from './App.vue'
import router from './router'
import { config } from './config'

const app = createApp(App)

app.use(createPinia())

app.use(GoogleSignInPlugin, {
  clientId: config.google_client_id,
});

app.use(createVuetify({
  theme: {
    defaultTheme: 'dark',
  },
  icons: {
    defaultSet: 'mdi',
  },
}))

app.use(router)

app.mount('#app')

