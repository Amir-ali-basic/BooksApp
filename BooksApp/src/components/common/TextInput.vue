<template>
  <div class="mb-3">
    <label :for="name" class="form-label">{{ label }}</label>
    <input
      :type="type"
      :name="name"
      :placeholder="placeholder"
      :required="required"
      :min="min"
      :max="max"
      :value="value"
      :autocomplete="name + '-new'"
      class="form-control"
      @input="handleInput"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

export default defineComponent({
  name: 'TextInput',
  props: {
    type: String,
    name: String,
    placeholder: String,
    required: Boolean,
    min: String,
    max: String,
    value: String,
    label: String
  },
  setup(props, { emit }) {
    const inputValue = ref(props.value)

    const handleInput = (event: Event) => {
      const target = event.target as HTMLInputElement
      inputValue.value = target.value
      emit('update:value', inputValue.value)
    }

    return {
      inputValue,
      handleInput
    }
  }
})
</script>
