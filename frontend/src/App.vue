<template>
  <div id="app">
    <div class="container">
      <header class="header">
        <h1>🚀 Sistema de Ordem de Serviços</h1>
        <p class="subtitle">Frontend Vue 3 + Backend Go</p>
      </header>
      
      <main class="main">
        <div class="status-card">
          <h2>Status dos Serviços</h2>
          
          <div class="service-status">
            <div class="status-item">
              <span class="status-icon frontend">✅</span>
              <div class="status-info">
                <h3>Frontend (Vue 3)</h3>
                <p>Rodando na porta 3000</p>
                <span class="status-badge success">ATIVO</span>
              </div>
            </div>
            
            <div class="status-item">
              <span class="status-icon backend" :class="{ 'checking': checkingBackend }">🔄</span>
              <div class="status-info">
                <h3>Backend (Go)</h3>
                <p v-if="checkingBackend">Verificando conexão...</p>
                <p v-else-if="backendStatus.online">Rodando na porta 8080</p>
                <p v-else>Serviço offline</p>
                <span class="status-badge" :class="backendStatus.online ? 'success' : 'error'">
                  {{ backendStatus.online ? 'ATIVO' : 'OFFLINE' }}
                </span>
              </div>
            </div>
          </div>
          
          <div class="connection-test" v-if="backendStatus.online">
            <h3>🧪 Teste de Conexão</h3>
            <div class="test-result">
              <p><strong>Endpoint:</strong> /api/v1/health</p>
              <p><strong>Resposta:</strong> {{ backendStatus.response }}</p>
              <p><strong>Tempo:</strong> {{ backendStatus.responseTime }}ms</p>
            </div>
          </div>
          
          <div class="actions">
            <button @click="testBackendConnection" class="btn btn-primary">
              🔄 Testar Conexão
            </button>
            <button @click="openBackendHealth" class="btn btn-secondary">
              🌐 Abrir Backend
            </button>
          </div>
        </div>
        
        <div class="info-card">
          <h3>📋 Informações do Sistema</h3>
          <ul>
            <li><strong>Frontend:</strong> Vue 3 + Vite + Pinia</li>
            <li><strong>Backend:</strong> Go + Gin + PostgreSQL</li>
            <li><strong>Padrões:</strong> Repository, Service Layer, Event Sourcing</li>
            <li><strong>Status:</strong> Desenvolvimento ativo</li>
          </ul>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const checkingBackend = ref(false)
const backendStatus = ref({
  online: false,
  response: '',
  responseTime: 0
})

const testBackendConnection = async () => {
  checkingBackend.value = true
  const startTime = Date.now()
  
  try {
    const response = await fetch('/api/v1/health')
    const data = await response.json()
    const endTime = Date.now()
    
    backendStatus.value = {
      online: true,
      response: JSON.stringify(data, null, 2),
      responseTime: endTime - startTime
    }
  } catch (error) {
    backendStatus.value = {
      online: false,
      response: `Erro: ${error.message}`,
      responseTime: 0
    }
  } finally {
    checkingBackend.value = false
  }
}

const openBackendHealth = () => {
  window.open('http://localhost:8080/api/v1/health', '_blank')
}

onMounted(() => {
  // Testar conexão automaticamente ao carregar
  testBackendConnection()
})
</script>

<style>
#app {
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin: 0;
  padding: 0;
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

* {
  box-sizing: border-box;
}

body {
  margin: 0;
  padding: 0;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

.header {
  text-align: center;
  margin-bottom: 3rem;
  color: white;
}

.header h1 {
  font-size: 3rem;
  margin: 0 0 1rem 0;
  font-weight: 700;
}

.subtitle {
  font-size: 1.2rem;
  opacity: 0.9;
  margin: 0;
}

.main {
  display: grid;
  gap: 2rem;
  grid-template-columns: 1fr;
}

.status-card, .info-card {
  background: white;
  border-radius: 16px;
  padding: 2rem;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
}

.status-card h2, .info-card h3 {
  margin: 0 0 1.5rem 0;
  color: #1f2937;
  font-size: 1.5rem;
}

.service-status {
  display: grid;
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.status-item {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1.5rem;
  background: #f8fafc;
  border-radius: 12px;
  border: 1px solid #e2e8f0;
}

.status-icon {
  font-size: 2rem;
  width: 60px;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background: #f0f9ff;
}

.status-icon.frontend {
  background: #dcfce7;
}

.status-icon.backend {
  background: #fef3c7;
}

.status-icon.backend.checking {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.status-info h3 {
  margin: 0 0 0.5rem 0;
  font-size: 1.1rem;
  color: #374151;
}

.status-info p {
  margin: 0 0 0.5rem 0;
  color: #6b7280;
  font-size: 0.9rem;
}

.status-badge {
  display: inline-block;
  padding: 0.25rem 0.75rem;
  border-radius: 9999px;
  font-size: 0.75rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.status-badge.success {
  background: #dcfce7;
  color: #166534;
}

.status-badge.error {
  background: #fee2e2;
  color: #991b1b;
}

.connection-test {
  margin-top: 2rem;
  padding: 1.5rem;
  background: #f0f9ff;
  border-radius: 12px;
  border: 1px solid #bae6fd;
}

.connection-test h3 {
  margin: 0 0 1rem 0;
  color: #0369a1;
}

.test-result p {
  margin: 0.25rem 0;
  font-family: 'Monaco', 'Menlo', monospace;
  font-size: 0.9rem;
  color: #0c4a6e;
}

.actions {
  display: flex;
  gap: 1rem;
  margin-top: 2rem;
  flex-wrap: wrap;
}

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  text-decoration: none;
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
}

.btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
}

.btn-primary {
  background: #3b82f6;
  color: white;
}

.btn-primary:hover {
  background: #2563eb;
}

.btn-secondary {
  background: #6b7280;
  color: white;
}

.btn-secondary:hover {
  background: #4b5563;
}

.info-card ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.info-card li {
  padding: 0.75rem 0;
  border-bottom: 1px solid #e5e7eb;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.info-card li:last-child {
  border-bottom: none;
}

.info-card strong {
  color: #374151;
}

@media (max-width: 768px) {
  .container {
    padding: 1rem;
  }
  
  .header h1 {
    font-size: 2rem;
  }
  
  .actions {
    flex-direction: column;
  }
  
  .btn {
    width: 100%;
    justify-content: center;
  }
}
</style>
