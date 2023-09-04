<template>
  <nav class="navbar navbar-expand-lg bg-body-tertiary" data-bs-theme="dark">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">Book app</a>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <router-link to="/" class="nav-link active" aria-current="page" href="#"
              >Home</router-link
            >
          </li>
          <li class="nav-item">
            <router-link
              :to="authStore.userToken === null ? '/login' : '/dashboard'"
              class="nav-link"
            >
              {{ authStore.userToken === null ? 'Log in' : 'Log out' }}
            </router-link>
          </li>
          <li class="nav-item">
            <a class="nav-link disabled">Disabled</a>
          </li>
        </ul>
        <form class="d-flex" role="search">
          <input class="form-control me-2" type="search" placeholder="Search" aria-label="Search" />
          <button class="btn btn-outline-success" type="submit">Search</button>
        </form>
      </div>
    </div>
  </nav>
</template>

<script lang="ts">
import { useAuthStore } from '@/stores/AuthStore'
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'HeaderComponent',
  setup() {
    const authStore = useAuthStore()
    return {
      authStore
    }
  }
})
</script>
