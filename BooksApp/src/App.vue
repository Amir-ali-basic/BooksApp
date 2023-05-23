<template>
  <HeaderComponent />
  <div>
    <router-view />
  </div>
  <FooterComponent />
</template>

<script setup lang="ts">
import { defineComponent } from 'vue'
import HeaderComponent from './components/common/Header.vue'
import FooterComponent from './components/common/Footer.vue'

defineComponent({
  name: 'App',
  components: {
    HeaderComponent,
    FooterComponent
  }
})
</script>

<style scoped></style>
