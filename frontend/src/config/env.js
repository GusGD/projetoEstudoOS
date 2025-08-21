export const config = {
  env: import.meta.env.VITE_APP_ENV || 'development',
  isDevelopment: import.meta.env.VITE_APP_ENV === 'development',
  isProduction: import.meta.env.VITE_APP_ENV === 'production',
  
  api: {
    baseUrl: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080',
    version: import.meta.env.VITE_API_VERSION || 'v1',
    timeout: parseInt(import.meta.env.VITE_API_TIMEOUT || '10000'),
    fullUrl: `${import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'}/api/${import.meta.env.VITE_API_VERSION || 'v1'}`
  },
  
  app: {
    title: import.meta.env.VITE_APP_TITLE || 'Sistema de Ordem de Serviços',
    url: import.meta.env.VITE_APP_URL || 'http://localhost:3000'
  },
  
  auth: {
    jwtStorageKey: import.meta.env.VITE_JWT_STORAGE_KEY || 'auth_token',
    refreshTokenKey: import.meta.env.VITE_REFRESH_TOKEN_KEY || 'refresh_token'
  }
}

export const getConfig = (path) => {
  return path.split('.').reduce((obj, key) => obj?.[key], config)
}

export const getApiUrl = (endpoint = '') => {
  return `${config.api.fullUrl}${endpoint}`
}

export default config
