import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { osService } from '@/api/osService'

export const useOSStore = defineStore('os', () => {
  // Estado reativo
  const ordens = ref([])
  const loading = ref(false)
  const error = ref(null)
  const currentOS = ref(null)
  const filters = ref({
    status: null,
    prioridade: null,
    responsavel: null,
    prefeitura: null
  })

  // Getters computados
  const ordensAbertas = computed(() => 
    ordens.value.filter(os => os.status === 'aberta')
  )
  
  const ordensEmAndamento = computed(() => 
    ordens.value.filter(os => os.status === 'em_andamento')
  )
  
  const ordensConcluidas = computed(() => 
    ordens.value.filter(os => os.status === 'concluida')
  )
  
  const ordensPorPrioridade = computed(() => {
    return ordens.value.sort((a, b) => b.prioridade - a.prioridade)
  })

  // Actions
  const fetchOrdens = async (prefeituraId, params = {}) => {
    loading.value = true
    error.value = null
    
    try {
      const response = await osService.getOrdens(prefeituraId, params)
      ordens.value = response.data
    } catch (err) {
      error.value = err.message || 'Erro ao carregar ordens de serviço'
      console.error('Erro ao buscar ordens:', err)
    } finally {
      loading.value = false
    }
  }

  const fetchOSById = async (id) => {
    loading.value = true
    error.value = null
    
    try {
      const response = await osService.getOSById(id)
      currentOS.value = response.data
      return response.data
    } catch (err) {
      error.value = err.message || 'Erro ao carregar ordem de serviço'
      console.error('Erro ao buscar OS:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  const createOS = async (osData) => {
    loading.value = true
    error.value = null
    
    try {
      const response = await osService.createOS(osData)
      ordens.value.unshift(response.data)
      return response.data
    } catch (err) {
      error.value = err.message || 'Erro ao criar ordem de serviço'
      console.error('Erro ao criar OS:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  const updateOS = async (id, osData) => {
    loading.value = true
    error.value = null
    
    try {
      const response = await osService.updateOS(id, osData)
      
      // Atualizar na lista local
      const index = ordens.value.findIndex(os => os.id === id)
      if (index !== -1) {
        ordens.value[index] = response.data
      }
      
      // Atualizar OS atual se for a mesma
      if (currentOS.value && currentOS.value.id === id) {
        currentOS.value = response.data
      }
      
      return response.data
    } catch (err) {
      error.value = err.message || 'Erro ao atualizar ordem de serviço'
      console.error('Erro ao atualizar OS:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  const alterarStatus = async (id, novoStatus, motivo, usuarioId) => {
    try {
      const response = await osService.alterarStatus(id, novoStatus, motivo, usuarioId)
      
      // Atualizar na lista local
      const index = ordens.value.findIndex(os => os.id === id)
      if (index !== -1) {
        ordens.value[index] = response.data
      }
      
      // Atualizar OS atual se for a mesma
      if (currentOS.value && currentOS.value.id === id) {
        currentOS.value = response.data
      }
      
      return response.data
    } catch (err) {
      error.value = err.message || 'Erro ao alterar status'
      console.error('Erro ao alterar status:', err)
      throw err
    }
  }

  const deleteOS = async (id) => {
    try {
      await osService.deleteOS(id)
      
      // Remover da lista local
      ordens.value = ordens.value.filter(os => os.id !== id)
      
      // Limpar OS atual se for a mesma
      if (currentOS.value && currentOS.value.id === id) {
        currentOS.value = null
      }
    } catch (err) {
      error.value = err.message || 'Erro ao deletar ordem de serviço'
      console.error('Erro ao deletar OS:', err)
      throw err
    }
  }

  const setFilters = (newFilters) => {
    filters.value = { ...filters.value, ...newFilters }
  }

  const clearFilters = () => {
    filters.value = {
      status: null,
      prioridade: null,
      responsavel: null,
      prefeitura: null
    }
  }

  const clearError = () => {
    error.value = null
  }

  const clearCurrentOS = () => {
    currentOS.value = null
  }

  return {
    // Estado
    ordens,
    loading,
    error,
    currentOS,
    filters,
    
    // Getters
    ordensAbertas,
    ordensEmAndamento,
    ordensConcluidas,
    ordensPorPrioridade,
    
    // Actions
    fetchOrdens,
    fetchOSById,
    createOS,
    updateOS,
    alterarStatus,
    deleteOS,
    setFilters,
    clearFilters,
    clearError,
    clearCurrentOS
  }
})
