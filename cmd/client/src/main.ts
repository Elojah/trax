import './assets/main.css'
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createVuetify } from 'vuetify'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createVuetify({
  theme: {
    defaultTheme:'dark',
  },
  icons: {
    defaultSet: 'mdi',
  },
}))

app.use(createPinia())

app.use(router)

app.mount('#app')

