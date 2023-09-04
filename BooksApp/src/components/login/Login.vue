<template>
  <div class="container">
    <div class="row">
      <div class="col-lg-12">
        <h1 class="mt-5">Login</h1>
        <h1>{{ authStore.userToken }}</h1>
        <hr />
        <FormTag name="myForm" event="submitEvent" @submitEvent="submitHandler">
          <TextInput
            label="Email"
            type="email"
            name="email"
            :required="true"
            :value="loginPresentable.email"
            @input="loginPresentable.email = $event.target.value"
          />
          <TextInput
            label="Password"
            type="password"
            name="password"
            :required="true"
            :value="loginPresentable.password"
            @input="loginPresentable.password = $event.target.value"
          />

          <hr />
          <button type="submit" class="btn btn-primary">Login</button>
        </FormTag>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import TextInput from '../common/TextInput.vue'
import FormTag from '../common/FormTag.vue'
import { useAuthStore } from '@/stores/AuthStore'

export default defineComponent({
  name: 'LoginComponent',
  components: {
    TextInput,
    FormTag
  },

  setup() {
    const authStore = useAuthStore()
    const loginPresentable = authStore.loginData

    function submitHandler() {
      console.log('submitHandler', loginPresentable)
      authStore.loginSubmitHandler(loginPresentable)
    }

    return {
      submitHandler,
      loginPresentable,
      authStore
    }
  }
})
</script>

<style scoped>
/* Your component's styles here */
</style>
