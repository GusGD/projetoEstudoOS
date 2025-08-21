import { createApp } from 'vue'
import { createPinia } from 'pinia'
import router from './router'
import App from './App.vue'

// Importar estilos globais
import './assets/main.css'

const app = createApp(App)

// Configurar Pinia para gerenciamento de estado
const pinia = createPinia()
app.use(pinia)

// Configurar Vue Router
app.use(router)

// Montar aplicação
app.mount('#app')
