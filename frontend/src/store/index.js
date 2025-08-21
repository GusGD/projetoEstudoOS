import { createPinia } from 'pinia'

// Criar instância do Pinia
const pinia = createPinia()

export default pinia

// Exportar store para uso em outros módulos
export { useAuthStore } from './modules/auth'
export { useOSStore } from './modules/os'
export { useUserStore } from './modules/users'
export { useNotificationStore } from './modules/notifications'
