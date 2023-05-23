<template>
  <LoginComponent />
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import LoginComponent from '@/components/login/Login.vue'

export default defineComponent({
  name: 'LoginView',
  components: {
    LoginComponent
  }
  // Your component's logic here
})
</script>

<style scoped>
/* Your component's styles here */
</style>
