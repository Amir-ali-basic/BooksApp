<template>
  <div class="form-chek">
    <input :checked="checked" :required="required" type="text" class="form-chek-input" />
    <label :for="name" class="form-check-label"></label>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'CheckBox',
  props: {
    checked: {
      type: Boolean,
      default: false
    },
    required: {
      type: Boolean,
      default: false
    },
    name: {
      type: String,
      default: ''
    }
  },
  setup() {
    return {}
  }
})
</script>
