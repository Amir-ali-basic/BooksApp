<template>
  <div class="mb-3">
    <label :for="name" class="form-label">{{ label }}</label>
    <select name="name" :id="name" :required="required">
      <option v-for="option in items" :value="option.value" :key="option.value">
        {{ option.text }}
      </option>
    </select>
  </div>
</template>

<script lang="ts">
import { defineComponent, type PropType } from 'vue'

export default defineComponent({
  name: 'SelectInput',
  props: {
    name: String,
    items: {
      type: Array as PropType<any[]>,
      default: () => []
    },
    required: Boolean,
    label: String
  },
  setup() {
    return {}
  }
})
</script>
