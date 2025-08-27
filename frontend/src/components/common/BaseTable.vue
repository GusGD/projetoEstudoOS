<template>
  <div class="base-table-wrapper">
    <!-- Header da tabela -->
    <div v-if="showHeader" class="table-header">
      <div class="table-header__left">
        <slot name="header-left">
          <h3 v-if="title" class="table-title">{{ title }}</h3>
        </slot>
      </div>
      
      <div class="table-header__right">
        <slot name="header-right">
          <div v-if="searchable" class="table-search">
            <BaseInput
              v-model="searchQuery"
              placeholder="Buscar..."
              icon="🔍"
              clearable
              @input="handleSearch"
            />
          </div>
        </slot>
      </div>
    </div>
    
    <!-- Tabela -->
    <div class="table-container">
      <table class="base-table">
        <thead>
          <tr>
            <!-- Checkbox de seleção múltipla -->
            <th v-if="selectable" class="table-checkbox">
              <input
                type="checkbox"
                :checked="isAllSelected"
                :indeterminate="isIndeterminate"
                @change="toggleSelectAll"
              />
            </th>
            
            <!-- Cabeçalhos das colunas -->
            <th
              v-for="column in visibleColumns"
              :key="column.key"
              :class="getHeaderClasses(column)"
              @click="handleSort(column)"
            >
              <div class="table-header-cell">
                <span>{{ column.label }}</span>
                <span v-if="column.sortable" class="sort-indicator">
                  {{ getSortIcon(column.key) }}
                </span>
              </div>
            </th>
            
            <!-- Coluna de ações -->
            <th v-if="showActions" class="table-actions">Ações</th>
          </tr>
        </thead>
        
        <tbody>
          <tr
            v-for="(item, index) in paginatedData"
            :key="getRowKey(item, index)"
            :class="getRowClasses(item, index)"
            @click="handleRowClick(item, index)"
          >
            <!-- Checkbox de seleção -->
            <td v-if="selectable" class="table-checkbox">
              <input
                type="checkbox"
                :checked="isSelected(item)"
                @change="toggleSelection(item)"
                @click.stop
              />
            </td>
            
            <!-- Células de dados -->
            <td
              v-for="column in visibleColumns"
              :key="column.key"
              :class="getCellClasses(column)"
            >
              <slot
                :name="`cell-${column.key}`"
                :item="item"
                :column="column"
                :value="getCellValue(item, column)"
              >
                {{ formatCellValue(getCellValue(item, column), column) }}
              </slot>
            </td>
            
            <!-- Célula de ações -->
            <td v-if="showActions" class="table-actions">
              <slot name="actions" :item="item" :index="index">
                <div class="action-buttons">
                  <BaseButton
                    size="small"
                    variant="info"
                    @click.stop="handleEdit(item)"
                  >
                    Editar
                  </BaseButton>
                  <BaseButton
                    size="small"
                    variant="danger"
                    @click.stop="handleDelete(item)"
                  >
                    Excluir
                  </BaseButton>
                </div>
              </slot>
            </td>
          </tr>
          
          <!-- Linha vazia -->
          <tr v-if="paginatedData.length === 0">
            <td :colspan="totalColumns" class="table-empty">
              <slot name="empty">
                <div class="empty-state">
                  <p>Nenhum item encontrado</p>
                </div>
              </slot>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    
    <!-- Paginação -->
    <div v-if="showPagination && totalPages > 1" class="table-pagination">
      <div class="pagination-info">
        <span>
          Mostrando {{ startIndex + 1 }}-{{ endIndex }} de {{ totalItems }} itens
        </span>
      </div>
      
      <div class="pagination-controls">
        <BaseButton
          size="small"
          variant="outline"
          :disabled="currentPage === 1"
          @click="goToPage(currentPage - 1)"
        >
          Anterior
        </BaseButton>
        
        <div class="page-numbers">
          <button
            v-for="page in visiblePages"
            :key="page"
            :class="['page-number', { active: page === currentPage }]"
            @click="goToPage(page)"
          >
            {{ page }}
          </button>
        </div>
        
        <BaseButton
          size="small"
          variant="outline"
          :disabled="currentPage === totalPages"
          @click="goToPage(currentPage + 1)"
        >
          Próxima
        </BaseButton>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, ref, watch } from 'vue'
import BaseInput from './BaseInput.vue'
import BaseButton from './BaseButton.vue'

const props = defineProps({
  data: {
    type: Array,
    default: () => []
  },
  columns: {
    type: Array,
    default: () => []
  },
  title: {
    type: String,
    default: ''
  },
  showHeader: {
    type: Boolean,
    default: true
  },
  showActions: {
    type: Boolean,
    default: true
  },
  selectable: {
    type: Boolean,
    default: false
  },
  searchable: {
    type: Boolean,
    default: false
  },
  showPagination: {
    type: Boolean,
    default: true
  },
  pageSize: {
    type: Number,
    default: 10
  },
  sortable: {
    type: Boolean,
    default: true
  },
  rowKey: {
    type: String,
    default: 'id'
  }
})

const emit = defineEmits([
  'row-click',
  'edit',
  'delete',
  'selection-change',
  'sort-change',
  'page-change'
])

// Estado interno
const searchQuery = ref('')
const currentPage = ref(1)
const sortColumn = ref('')
const sortDirection = ref('asc')
const selectedItems = ref(new Set())

// Computed properties
const filteredData = computed(() => {
  if (!searchQuery.value) return props.data
  
  return props.data.filter(item => {
    return props.columns.some(column => {
      const value = getCellValue(item, column)
      return String(value).toLowerCase().includes(searchQuery.value.toLowerCase())
    })
  })
})

const sortedData = computed(() => {
  if (!sortColumn.value || !props.sortable) return filteredData.value
  
  return [...filteredData.value].sort((a, b) => {
    const aValue = getCellValue(a, { key: sortColumn.value })
    const bValue = getCellValue(b, { key: sortColumn.value })
    
    if (aValue < bValue) return sortDirection.value === 'asc' ? -1 : 1
    if (aValue > bValue) return sortDirection.value === 'asc' ? 1 : -1
    return 0
  })
})

const totalItems = computed(() => sortedData.value.length)
const totalPages = computed(() => Math.ceil(totalItems.value / props.pageSize))
const startIndex = computed(() => (currentPage.value - 1) * props.pageSize)
const endIndex = computed(() => Math.min(startIndex.value + props.pageSize, totalItems.value))

const paginatedData = computed(() => {
  return sortedData.value.slice(startIndex.value, endIndex.value)
})

const visibleColumns = computed(() => {
  return props.columns.filter(column => !column.hidden)
})

const totalColumns = computed(() => {
  let count = visibleColumns.value.length
  if (props.selectable) count++
  if (props.showActions) count++
  return count
})

const isAllSelected = computed(() => {
  return paginatedData.value.length > 0 && 
         paginatedData.value.every(item => isSelected(item))
})

const isIndeterminate = computed(() => {
  const selectedCount = paginatedData.value.filter(item => isSelected(item)).length
  return selectedCount > 0 && selectedCount < paginatedData.value.length
})

const visiblePages = computed(() => {
  const pages = []
  const maxVisible = 5
  
  if (totalPages.value <= maxVisible) {
    for (let i = 1; i <= totalPages.value; i++) {
      pages.push(i)
    }
  } else {
    const start = Math.max(1, currentPage.value - Math.floor(maxVisible / 2))
    const end = Math.min(totalPages.value, start + maxVisible - 1)
    
    for (let i = start; i <= end; i++) {
      pages.push(i)
    }
  }
  
  return pages
})

// Methods
const getRowKey = (item, index) => {
  return item[props.rowKey] || index
}

const getCellValue = (item, column) => {
  if (column.value) {
    return typeof column.value === 'function' ? column.value(item) : item[column.value]
  }
  return item[column.key]
}

const formatCellValue = (value, column) => {
  if (column.formatter) {
    return column.formatter(value, column)
  }
  return value
}

const getHeaderClasses = (column) => {
  return {
    'sortable': column.sortable !== false && props.sortable,
    'sorted': sortColumn.value === column.key
  }
}

const getCellClasses = (column) => {
  return {
    [`cell-${column.key}`]: true,
    [`align-${column.align || 'left'}`]: true
  }
}

const getRowClasses = (item, index) => {
  return {
    'table-row': true,
    'table-row--selected': isSelected(item),
    'table-row--clickable': true
  }
}

const getSortIcon = (columnKey) => {
  if (sortColumn.value !== columnKey) return '↕'
  return sortDirection.value === 'asc' ? '↑' : '↓'
}

const isSelected = (item) => {
  return selectedItems.value.has(getRowKey(item))
}

// Event handlers
const handleSearch = () => {
  currentPage.value = 1
}

const handleSort = (column) => {
  if (!column.sortable && !props.sortable) return
  
  if (sortColumn.value === column.key) {
    sortDirection.value = sortDirection.value === 'asc' ? 'desc' : 'asc'
  } else {
    sortColumn.value = column.key
    sortDirection.value = 'asc'
  }
  
  emit('sort-change', { column: sortColumn.value, direction: sortDirection.value })
}

const handleRowClick = (item, index) => {
  emit('row-click', { item, index })
}

const handleEdit = (item) => {
  emit('edit', item)
}

const handleDelete = (item) => {
  emit('delete', item)
}

const toggleSelection = (item) => {
  const key = getRowKey(item)
  if (selectedItems.value.has(key)) {
    selectedItems.value.delete(key)
  } else {
    selectedItems.value.add(key)
  }
  
  emit('selection-change', Array.from(selectedItems.value))
}

const toggleSelectAll = () => {
  if (isAllSelected.value) {
    selectedItems.value.clear()
  } else {
    paginatedData.value.forEach(item => {
      selectedItems.value.add(getRowKey(item))
    })
  }
  
  emit('selection-change', Array.from(selectedItems.value))
}

const goToPage = (page) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
    emit('page-change', page)
  }
}

// Watchers
watch(() => props.data, () => {
  currentPage.value = 1
  selectedItems.value.clear()
})

watch(() => props.pageSize, () => {
  currentPage.value = 1
})
</script>

<style scoped>
.base-table-wrapper {
  background: white;
  border-radius: 0.75rem;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06);
  overflow: hidden;
}

/* Header da tabela */
.table-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1.5rem;
  border-bottom: 1px solid #e5e7eb;
  background-color: #f9fafb;
}

.table-header__left {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.table-title {
  margin: 0;
  font-size: 1.25rem;
  font-weight: 600;
  color: #111827;
}

.table-header__right {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.table-search {
  width: 16rem;
}

/* Container da tabela */
.table-container {
  overflow-x: auto;
}

.base-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 0.875rem;
}

/* Cabeçalhos */
.base-table th {
  background-color: #f9fafb;
  padding: 0.75rem 1rem;
  text-align: left;
  font-weight: 600;
  color: #374151;
  border-bottom: 1px solid #e5e7eb;
  white-space: nowrap;
}

.base-table th.sortable {
  cursor: pointer;
  user-select: none;
}

.base-table th.sortable:hover {
  background-color: #f3f4f6;
}

.table-header-cell {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.sort-indicator {
  font-size: 0.75rem;
  color: #6b7280;
}

/* Células */
.base-table td {
  padding: 0.75rem 1rem;
  border-bottom: 1px solid #f3f4f6;
  color: #111827;
}

.base-table tr:hover {
  background-color: #f9fafb;
}

.table-row--selected {
  background-color: #eff6ff;
}

.table-row--selected:hover {
  background-color: #dbeafe;
}

/* Checkbox */
.table-checkbox {
  width: 3rem;
  text-align: center;
}

.table-checkbox input[type="checkbox"] {
  width: 1rem;
  height: 1rem;
}

/* Ações */
.table-actions {
  width: 8rem;
  text-align: center;
}

.action-buttons {
  display: flex;
  gap: 0.5rem;
  justify-content: center;
}

/* Estado vazio */
.table-empty {
  text-align: center;
  padding: 3rem;
  color: #6b7280;
}

.empty-state p {
  margin: 0;
  font-size: 1rem;
}

/* Paginação */
.table-pagination {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1rem 1.5rem;
  border-top: 1px solid #e5e7eb;
  background-color: #f9fafb;
}

.pagination-info {
  color: #6b7280;
  font-size: 0.875rem;
}

.pagination-controls {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.page-numbers {
  display: flex;
  gap: 0.25rem;
}

.page-number {
  padding: 0.5rem 0.75rem;
  border: 1px solid #d1d5db;
  background: white;
  color: #374151;
  cursor: pointer;
  border-radius: 0.375rem;
  font-size: 0.875rem;
  transition: all 0.2s ease-in-out;
}

.page-number:hover {
  background-color: #f3f4f6;
  border-color: #9ca3af;
}

.page-number.active {
  background-color: #3b82f6;
  border-color: #3b82f6;
  color: white;
}

/* Alinhamento de células */
.cell-left {
  text-align: left;
}

.cell-center {
  text-align: center;
}

.cell-right {
  text-align: right;
}

/* Responsividade */
@media (max-width: 768px) {
  .table-header {
    flex-direction: column;
    gap: 1rem;
    align-items: stretch;
  }
  
  .table-search {
    width: 100%;
  }
  
  .pagination-controls {
    flex-direction: column;
    gap: 0.5rem;
  }
  
  .page-numbers {
    order: -1;
  }
}
</style>
