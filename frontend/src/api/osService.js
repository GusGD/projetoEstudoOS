import axios from 'axios'
import { config } from '@/config/env'

const api = axios.create({
  baseURL: config.api.fullUrl,
  timeout: config.api.timeout,
  headers: {
    'Content-Type': 'application/json'
  }
})

api.interceptors.request.use(
  (requestConfig) => {
    const token = localStorage.getItem(config.auth.jwtStorageKey)
    if (token) {
      requestConfig.headers.Authorization = `Bearer ${token}`
    }
    return requestConfig
  },
  (error) => {
    return Promise.reject(error)
  }
)

api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem(config.auth.jwtStorageKey)
      window.location.href = config.auth.loginUrl
    }
    return Promise.reject(error)
  }
)

export const osService = {
  async getOrdens(prefeituraId, params = {}) {
    const queryParams = new URLSearchParams({
      prefeitura_id: prefeituraId,
      ...params
    })
    
    return api.get(`/os?${queryParams}`)
  },

  async getOSById(id) {
    return api.get(`/os/${id}`)
  },

  async createOS(osData) {
    return api.post('/os', osData)
  },

  async updateOS(id, osData) {
    return api.put(`/os/${id}`, osData)
  },

  async deleteOS(id) {
    return api.delete(`/os/${id}`)
  },

  async alterarStatus(id, novoStatus, motivo, usuarioId) {
    return api.patch(`/os/${id}/status`, {
      status: novoStatus,
      motivo: motivo,
      usuario_id: usuarioId
    })
  },

  async atribuirResponsavel(id, responsavelId, usuarioId) {
    return api.patch(`/os/${id}/responsavel`, {
      responsavel_id: responsavelId,
      usuario_id: usuarioId
    })
  },

  async getOSByStatus(status, prefeituraId) {
    return api.get(`/os/status/${status}`, {
      params: { prefeitura_id: prefeituraId }
    })
  },

  async getOSByResponsavel(responsavelId) {
    return api.get(`/os/responsavel/${responsavelId}`)
  },

  async getHistorico(osId) {
    return api.get(`/os/${osId}/historico`)
  },

  async adicionarTag(id, tag) {
    return api.post(`/os/${id}/tags`, { tag })
  },

  async removerTag(id, tag) {
    return api.delete(`/os/${id}/tags/${encodeURIComponent(tag)}`)
  },

  async getOSByPrioridade(prioridade, prefeituraId) {
    return api.get(`/os/prioridade/${prioridade}`, {
      params: { prefeitura_id: prefeituraId }
    })
  },

  async transferirOS(id, novaPrefeituraId, motivo, usuarioId) {
    return api.post(`/os/${id}/transferir`, {
      nova_prefeitura_id: novaPrefeituraId,
      motivo: motivo,
      usuario_id: usuarioId
    })
  },

  async getEstatisticas(prefeituraId) {
    return api.get(`/os/estatisticas`, {
      params: { prefeitura_id: prefeituraId }
    })
  },

  async exportarRelatorio(prefeituraId, formato = 'pdf', filtros = {}) {
    return api.post(`/os/relatorio/exportar`, {
      prefeitura_id: prefeituraId,
      formato: formato,
      filtros: filtros
    }, {
      responseType: 'blob'
    })
  }
}

export { api }
