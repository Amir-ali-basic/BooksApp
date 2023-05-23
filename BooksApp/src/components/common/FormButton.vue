<template>
  <button class="form-button" :disabled="disabled" @click="handleClick">
    {{ text }}
  </button>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'FormButton',
  props: {
    disabled: {
      type: Boolean,
      default: false
    },
    text: {
      type: String,
      default: ''
    }
  },
  setup(props, { emit }) {
    const handleClick = () => {
      emit('clickButton')
    }

    return {
      handleClick
    }
  }
})
</script>

<style scoped>
.form-button {
  display: inline-block;
  padding: 10px 20px;
  font-size: 16px;
  font-weight: bold;
  text-align: center;
  color: #ffffff;
  background-color: #007bff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.form-button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}
</style>
