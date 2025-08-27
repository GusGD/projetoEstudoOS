<template>
  <Teleport to="body">
    <Transition name="modal">
      <div v-if="isOpen" class="modal-overlay" @click="handleBackdropClick">
        <div 
          class="modal-container"
          :class="modalClasses"
          @click.stop
        >
          <!-- Header -->
          <div v-if="showHeader" class="modal-header">
            <h3 v-if="title" class="modal-title">{{ title }}</h3>
            <button
              v-if="closable"
              type="button"
              class="modal-close"
              @click="handleClose"
              aria-label="Fechar modal"
            >
              ×
            </button>
          </div>
          
          <!-- Content -->
          <div class="modal-content">
            <slot />
          </div>
          
          <!-- Footer -->
          <div v-if="showFooter" class="modal-footer">
            <slot name="footer">
              <BaseButton
                v-if="closable"
                variant="secondary"
                @click="handleClose"
              >
                {{ cancelText }}
              </BaseButton>
              <BaseButton
                v-if="confirmText"
                variant="primary"
                @click="handleConfirm"
              >
                {{ confirmText }}
              </BaseButton>
            </slot>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { computed, watch, onMounted, onUnmounted, nextTick } from 'vue'
import BaseButton from './BaseButton.vue'

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  title: {
    type: String,
    default: ''
  },
  size: {
    type: String,
    default: 'medium',
    validator: (value) => ['small', 'medium', 'large', 'full'].includes(value)
  },
  closable: {
    type: Boolean,
    default: true
  },
  showHeader: {
    type: Boolean,
    default: true
  },
  showFooter: {
    type: Boolean,
    default: true
  },
  confirmText: {
    type: String,
    default: 'Confirmar'
  },
  cancelText: {
    type: String,
    default: 'Cancelar'
  },
  closeOnBackdrop: {
    type: Boolean,
    default: true
  },
  closeOnEscape: {
    type: Boolean,
    default: true
  }
})

const emit = defineEmits(['update:modelValue', 'close', 'confirm'])

const isOpen = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const modalClasses = computed(() => [
  'modal-container',
  `modal-container--${props.size}`
])

const handleClose = () => {
  isOpen.value = false
  emit('close')
}

const handleConfirm = () => {
  emit('confirm')
}

const handleBackdropClick = () => {
  if (props.closeOnBackdrop) {
    handleClose()
  }
}

const handleEscape = (event) => {
  if (event.key === 'Escape' && props.closeOnEscape) {
    handleClose()
  }
}

// Gerenciar foco e scroll
const manageFocus = () => {
  if (isOpen.value) {
    // Salvar elemento ativo
    const activeElement = document.activeElement
    document.body.style.overflow = 'hidden'
    
    // Focar no modal
    nextTick(() => {
      const modal = document.querySelector('.modal-container')
      if (modal) {
        modal.focus()
      }
    })
    
    return activeElement
  } else {
    document.body.style.overflow = ''
    return null
  }
}

// Observar mudanças no modal
watch(isOpen, (newValue) => {
  if (newValue) {
    manageFocus()
  } else {
    document.body.style.overflow = ''
  }
})

// Event listeners
onMounted(() => {
  if (props.closeOnEscape) {
    document.addEventListener('keydown', handleEscape)
  }
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleEscape)
  document.body.style.overflow = ''
})
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  padding: 1rem;
}

.modal-container {
  background: white;
  border-radius: 0.75rem;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
  max-height: 90vh;
  overflow: hidden;
  display: flex;
  flex-direction: column;
  outline: none;
}

/* Tamanhos */
.modal-container--small {
  width: 100%;
  max-width: 24rem;
}

.modal-container--medium {
  width: 100%;
  max-width: 32rem;
}

.modal-container--large {
  width: 100%;
  max-width: 48rem;
}

.modal-container--full {
  width: 100%;
  max-width: 90vw;
  height: 90vh;
}

/* Header */
.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1.5rem 1.5rem 0;
  border-bottom: 1px solid #e5e7eb;
  padding-bottom: 1rem;
}

.modal-title {
  margin: 0;
  font-size: 1.25rem;
  font-weight: 600;
  color: #111827;
}

.modal-close {
  background: none;
  border: none;
  font-size: 1.5rem;
  color: #6b7280;
  cursor: pointer;
  padding: 0.5rem;
  border-radius: 0.375rem;
  transition: all 0.2s ease-in-out;
  line-height: 1;
}

.modal-close:hover {
  background-color: #f3f4f6;
  color: #374151;
}

/* Content */
.modal-content {
  flex: 1;
  padding: 1.5rem;
  overflow-y: auto;
}

/* Footer */
.modal-footer {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 0.75rem;
  padding: 0 1.5rem 1.5rem;
  border-top: 1px solid #e5e7eb;
  padding-top: 1rem;
}

/* Animações */
.modal-enter-active,
.modal-leave-active {
  transition: all 0.3s ease-in-out;
}

.modal-enter-from {
  opacity: 0;
  transform: scale(0.95) translateY(-10px);
}

.modal-leave-to {
  opacity: 0;
  transform: scale(0.95) translateY(10px);
}

/* Responsividade */
@media (max-width: 640px) {
  .modal-overlay {
    padding: 0.5rem;
  }
  
  .modal-container {
    border-radius: 0.5rem;
    max-height: 95vh;
  }
  
  .modal-header,
  .modal-content,
  .modal-footer {
    padding: 1rem;
  }
  
  .modal-footer {
    flex-direction: column;
    align-items: stretch;
  }
  
  .modal-footer .base-button {
    width: 100%;
  }
}

/* Scrollbar personalizado */
.modal-content::-webkit-scrollbar {
  width: 6px;
}

.modal-content::-webkit-scrollbar-track {
  background: #f1f5f9;
  border-radius: 3px;
}

.modal-content::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 3px;
}

.modal-content::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}
</style>
