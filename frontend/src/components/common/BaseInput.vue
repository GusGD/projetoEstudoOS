<template>
  <div class="base-input-wrapper">
    <label v-if="label" :for="inputId" class="base-input__label">
      {{ label }}
      <span v-if="required" class="base-input__required">*</span>
    </label>
    
    <div class="base-input__container">
      <input
        :id="inputId"
        :type="type"
        :value="modelValue"
        :placeholder="placeholder"
        :disabled="disabled"
        :readonly="readonly"
        :maxlength="maxlength"
        :minlength="minlength"
        :pattern="pattern"
        :autocomplete="autocomplete"
        :autofocus="autofocus"
        class="base-input"
        :class="inputClasses"
        @input="handleInput"
        @blur="handleBlur"
        @focus="handleFocus"
        @keyup="handleKeyup"
      />
      
      <div v-if="icon" class="base-input__icon">
        <slot name="icon">
          <span class="base-input__icon-default">{{ icon }}</span>
        </slot>
      </div>
      
      <button
        v-if="clearable && modelValue"
        type="button"
        class="base-input__clear"
        @click="handleClear"
        aria-label="Limpar campo"
      >
        ×
      </button>
    </div>
    
    <div v-if="error" class="base-input__error">
      {{ error }}
    </div>
    
    <div v-if="hint" class="base-input__hint">
      {{ hint }}
    </div>
    
    <div v-if="showCharacterCount" class="base-input__character-count">
      {{ currentLength }}/{{ maxlength }}
    </div>
  </div>
</template>

<script setup>
import { computed, ref, nextTick } from 'vue'

const props = defineProps({
  modelValue: {
    type: [String, Number],
    default: ''
  },
  type: {
    type: String,
    default: 'text'
  },
  label: {
    type: String,
    default: ''
  },
  placeholder: {
    type: String,
    default: ''
  },
  required: {
    type: Boolean,
    default: false
  },
  disabled: {
    type: Boolean,
    default: false
  },
  readonly: {
    type: Boolean,
    default: false
  },
  error: {
    type: String,
    default: ''
  },
  hint: {
    type: String,
    default: ''
  },
  icon: {
    type: String,
    default: ''
  },
  clearable: {
    type: Boolean,
    default: false
  },
  maxlength: {
    type: Number,
    default: null
  },
  minlength: {
    type: Number,
    default: null
  },
  pattern: {
    type: String,
    default: ''
  },
  autocomplete: {
    type: String,
    default: 'off'
  },
  autofocus: {
    type: Boolean,
    default: false
  },
  showCharacterCount: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'blur', 'focus', 'keyup', 'clear'])

const inputId = ref(`input-${Math.random().toString(36).substr(2, 9)}`)
const isFocused = ref(false)

const inputClasses = computed(() => [
  'base-input',
  {
    'base-input--error': props.error,
    'base-input--focused': isFocused,
    'base-input--disabled': props.disabled,
    'base-input--readonly': props.readonly,
    'base-input--with-icon': props.icon,
    'base-input--clearable': props.clearable
  }
])

const currentLength = computed(() => {
  return String(props.modelValue || '').length
})

const handleInput = (event) => {
  emit('update:modelValue', event.target.value)
}

const handleBlur = (event) => {
  isFocused.value = false
  emit('blur', event)
}

const handleFocus = (event) => {
  isFocused.value = true
  emit('focus', event)
}

const handleKeyup = (event) => {
  emit('keyup', event)
}

const handleClear = () => {
  emit('update:modelValue', '')
  emit('clear')
  nextTick(() => {
    document.getElementById(inputId.value)?.focus()
  })
}
</script>

<style scoped>
.base-input-wrapper {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.base-input__label {
  font-size: 0.875rem;
  font-weight: 500;
  color: #374151;
  cursor: pointer;
}

.base-input__required {
  color: #ef4444;
  margin-left: 0.125rem;
}

.base-input__container {
  position: relative;
  display: flex;
  align-items: center;
}

.base-input {
  width: 100%;
  padding: 0.75rem 1rem;
  border: 2px solid #d1d5db;
  border-radius: 0.5rem;
  font-size: 1rem;
  line-height: 1.5;
  color: #111827;
  background-color: white;
  transition: all 0.2s ease-in-out;
  outline: none;
}

.base-input:focus {
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.base-input--error {
  border-color: #ef4444;
}

.base-input--error:focus {
  border-color: #ef4444;
  box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.1);
}

.base-input--disabled {
  background-color: #f9fafb;
  color: #9ca3af;
  cursor: not-allowed;
}

.base-input--readonly {
  background-color: #f9fafb;
}

.base-input--with-icon {
  padding-right: 2.5rem;
}

.base-input--clearable {
  padding-right: 2.5rem;
}

.base-input__icon {
  position: absolute;
  right: 1rem;
  color: #6b7280;
  pointer-events: none;
}

.base-input__icon-default {
  font-size: 1.25rem;
}

.base-input__clear {
  position: absolute;
  right: 0.75rem;
  background: none;
  border: none;
  font-size: 1.25rem;
  color: #9ca3af;
  cursor: pointer;
  padding: 0.25rem;
  border-radius: 50%;
  transition: all 0.2s ease-in-out;
}

.base-input__clear:hover {
  background-color: #f3f4f6;
  color: #6b7280;
}

.base-input__error {
  font-size: 0.875rem;
  color: #ef4444;
}

.base-input__hint {
  font-size: 0.875rem;
  color: #6b7280;
}

.base-input__character-count {
  font-size: 0.75rem;
  color: #9ca3af;
  text-align: right;
}

/* Estados de hover */
.base-input:hover:not(:disabled):not(:readonly) {
  border-color: #9ca3af;
}

/* Placeholder */
.base-input::placeholder {
  color: #9ca3af;
}

/* Autofill */
.base-input:-webkit-autofill {
  -webkit-box-shadow: 0 0 0 30px white inset;
  -webkit-text-fill-color: #111827;
}
</style>
