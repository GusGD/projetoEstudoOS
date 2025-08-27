<template>
  <div class="component-library">
    <div class="library-header">
      <h1>Biblioteca de Componentes Base</h1>
      <p>Sistema de design reutilizável para o Sistema de Ordem de Serviços</p>
    </div>

    <!-- Seção de Botões -->
    <section class="component-section">
      <h2>Botões</h2>
      <div class="component-showcase">
        <div class="showcase-item">
          <h3>Variantes</h3>
          <div class="button-grid">
            <BaseButton variant="primary">Primário</BaseButton>
            <BaseButton variant="secondary">Secundário</BaseButton>
            <BaseButton variant="success">Sucesso</BaseButton>
            <BaseButton variant="danger">Perigo</BaseButton>
            <BaseButton variant="warning">Aviso</BaseButton>
            <BaseButton variant="info">Info</BaseButton>
            <BaseButton variant="outline">Outline</BaseButton>
          </div>
        </div>

        <div class="showcase-item">
          <h3>Tamanhos</h3>
          <div class="button-grid">
            <BaseButton size="small">Pequeno</BaseButton>
            <BaseButton size="medium">Médio</BaseButton>
            <BaseButton size="large">Grande</BaseButton>
          </div>
        </div>

        <div class="showcase-item">
          <h3>Estados</h3>
          <div class="button-grid">
            <BaseButton :loading="true">Carregando</BaseButton>
            <BaseButton disabled>Desabilitado</BaseButton>
            <BaseButton fullWidth>Largura Total</BaseButton>
          </div>
        </div>
      </div>
    </section>

    <!-- Seção de Inputs -->
    <section class="component-section">
      <h2>Inputs</h2>
      <div class="component-showcase">
        <div class="showcase-item">
          <h3>Tipos Básicos</h3>
          <div class="input-grid">
            <BaseInput
              v-model="inputValues.text"
              label="Texto"
              placeholder="Digite seu texto"
            />
            <BaseInput
              v-model="inputValues.email"
              type="email"
              label="Email"
              placeholder="seu@email.com"
              required
            />
            <BaseInput
              v-model="inputValues.password"
              type="password"
              label="Senha"
              placeholder="Digite sua senha"
              required
            />
          </div>
        </div>

        <div class="showcase-item">
          <h3>Estados e Validação</h3>
          <div class="input-grid">
            <BaseInput
              v-model="inputValues.withIcon"
              label="Com Ícone"
              icon="🔍"
              placeholder="Buscar..."
            />
            <BaseInput
              v-model="inputValues.clearable"
              label="Limpar"
              placeholder="Digite algo..."
              clearable
            />
            <BaseInput
              v-model="inputValues.withError"
              label="Com Erro"
              placeholder="Campo com erro"
              error="Este campo é obrigatório"
            />
            <BaseInput
              v-model="inputValues.withHint"
              label="Com Dica"
              placeholder="Campo com dica"
              hint="Esta é uma dica útil para o usuário"
            />
          </div>
        </div>

        <div class="showcase-item">
          <h3>Contadores e Limites</h3>
          <div class="input-grid">
            <BaseInput
              v-model="inputValues.withMaxLength"
              label="Com Limite"
              placeholder="Máximo 50 caracteres"
              :maxlength="50"
              showCharacterCount
            />
          </div>
        </div>
      </div>
    </section>

    <!-- Seção de Modais -->
    <section class="component-section">
      <h2>Modais</h2>
      <div class="component-showcase">
        <div class="showcase-item">
          <h3>Diferentes Tamanhos</h3>
          <div class="button-grid">
            <BaseButton @click="openModal('small')">Modal Pequeno</BaseButton>
            <BaseButton @click="openModal('medium')">Modal Médio</BaseButton>
            <BaseButton @click="openModal('large')">Modal Grande</BaseButton>
            <BaseButton @click="openModal('full')">Modal Completo</BaseButton>
          </div>
        </div>

        <div class="showcase-item">
          <h3>Modal Sem Footer</h3>
          <BaseButton @click="openModalNoFooter">Sem Footer</BaseButton>
        </div>
      </div>
    </section>

    <!-- Seção de Tabelas -->
    <section class="component-section">
      <h2>Tabelas</h2>
      <div class="component-showcase">
        <div class="showcase-item">
          <h3>Tabela Básica</h3>
          <BaseTable
            :data="tableData"
            :columns="tableColumns"
            title="Exemplo de Tabela"
            searchable
            selectable
            :page-size="5"
            @row-click="handleRowClick"
            @edit="handleEdit"
            @delete="handleDelete"
            @selection-change="handleSelectionChange"
          />
        </div>
      </div>
    </section>

    <!-- Modais -->
    <BaseModal
      v-model="modalState.isOpen"
      :title="modalState.title"
      :size="modalState.size"
      :show-footer="modalState.showFooter"
      @confirm="handleModalConfirm"
    >
      <div class="modal-content-demo">
        <p>Este é um modal de demonstração com tamanho <strong>{{ modalState.size }}</strong>.</p>
        <p>Você pode personalizar o conteúdo, tamanho e comportamento conforme necessário.</p>
      </div>
    </BaseModal>

    <BaseModal
      v-model="modalState.isOpenNoFooter"
      title="Modal Sem Footer"
      :show-footer="false"
    >
      <div class="modal-content-demo">
        <p>Este modal não possui footer e pode ser fechado clicando no X ou no backdrop.</p>
        <BaseButton @click="modalState.isOpenNoFooter = false">Fechar</BaseButton>
      </div>
    </BaseModal>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import BaseButton from '@/components/common/BaseButton.vue'
import BaseInput from '@/components/common/BaseInput.vue'
import BaseModal from '@/components/common/BaseModal.vue'
import BaseTable from '@/components/common/BaseTable.vue'

// Estado dos inputs
const inputValues = reactive({
  text: '',
  email: '',
  password: '',
  withIcon: '',
  clearable: '',
  withError: '',
  withHint: '',
  withMaxLength: ''
})

// Estado dos modais
const modalState = reactive({
  isOpen: false,
  isOpenNoFooter: false,
  title: '',
  size: 'medium',
  showFooter: true
})

// Dados da tabela
const tableData = ref([
  { id: 1, title: 'OS #001', status: 'Aberto', priority: 'Alta', responsible: 'João Silva', created: '2024-01-15' },
  { id: 2, title: 'OS #002', status: 'Em Andamento', priority: 'Média', responsible: 'Maria Santos', created: '2024-01-14' },
  { id: 3, title: 'OS #003', status: 'Fechado', priority: 'Baixa', responsible: 'Pedro Costa', created: '2024-01-13' },
  { id: 4, title: 'OS #004', status: 'Aberto', priority: 'Alta', responsible: 'Ana Oliveira', created: '2024-01-12' },
  { id: 5, title: 'OS #005', status: 'Em Andamento', priority: 'Média', responsible: 'Carlos Lima', created: '2024-01-11' }
])

const tableColumns = ref([
  { key: 'id', label: 'ID', sortable: true, align: 'center' },
  { key: 'title', label: 'Título', sortable: true },
  { 
    key: 'status', 
    label: 'Status', 
    sortable: true,
    formatter: (value) => {
      const statusMap = {
        'Aberto': '🟢',
        'Em Andamento': '🟡',
        'Fechado': '🔴'
      }
      return `${statusMap[value] || '❓'} ${value}`
    }
  },
  { 
    key: 'priority', 
    label: 'Prioridade', 
    sortable: true,
    formatter: (value) => {
      const priorityMap = {
        'Alta': '🔴',
        'Média': '🟡',
        'Baixa': '🟢'
      }
      return `${priorityMap[value] || '❓'} ${value}`
    }
  },
  { key: 'responsible', label: 'Responsável', sortable: true },
  { 
    key: 'created', 
    label: 'Criado em', 
    sortable: true,
    formatter: (value) => new Date(value).toLocaleDateString('pt-BR')
  }
])

// Métodos dos modais
const openModal = (size) => {
  modalState.size = size
  modalState.title = `Modal ${size.charAt(0).toUpperCase() + size.slice(1)}`
  modalState.showFooter = true
  modalState.isOpen = true
}

const openModalNoFooter = () => {
  modalState.isOpenNoFooter = true
}

const handleModalConfirm = () => {
  modalState.isOpen = false
}

// Métodos da tabela
const handleRowClick = ({ item, index }) => {
  console.log('Linha clicada:', item, index)
}

const handleEdit = (item) => {
  console.log('Editar item:', item)
}

const handleDelete = (item) => {
  console.log('Excluir item:', item)
}

const handleSelectionChange = (selectedItems) => {
  console.log('Itens selecionados:', selectedItems)
}
</script>

<style scoped>
.component-library {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

.library-header {
  text-align: center;
  margin-bottom: 3rem;
}

.library-header h1 {
  font-size: 2.5rem;
  color: #111827;
  margin-bottom: 0.5rem;
}

.library-header p {
  font-size: 1.125rem;
  color: #6b7280;
}

.component-section {
  margin-bottom: 4rem;
}

.component-section h2 {
  font-size: 1.875rem;
  color: #111827;
  margin-bottom: 1.5rem;
  padding-bottom: 0.5rem;
  border-bottom: 2px solid #e5e7eb;
}

.component-showcase {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.showcase-item {
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 0.75rem;
  padding: 1.5rem;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
}

.showcase-item h3 {
  font-size: 1.25rem;
  color: #374151;
  margin-bottom: 1rem;
}

.button-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
  align-items: center;
}

.input-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 1rem;
}

.modal-content-demo {
  text-align: center;
}

.modal-content-demo p {
  margin-bottom: 1rem;
  color: #374151;
  line-height: 1.6;
}

/* Responsividade */
@media (max-width: 768px) {
  .component-library {
    padding: 1rem;
  }
  
  .library-header h1 {
    font-size: 2rem;
  }
  
  .button-grid {
    flex-direction: column;
    align-items: stretch;
  }
  
  .input-grid {
    grid-template-columns: 1fr;
  }
}
</style>
