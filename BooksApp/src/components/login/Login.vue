<template>
  <div class="container">
    <div class="row">
      <div class="col-lg-12">
        <h1 class="mt-5">Login</h1>
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
import { useLoginStore } from '@/stores/LoginStore'

export default defineComponent({
  name: 'LoginComponent',
  components: {
    TextInput,
    FormTag
  },

  setup() {
    const loginStore = useLoginStore()
    const loginPresentable = loginStore.loginData

    function submitHandler() {
      console.log('submitHandler', loginPresentable)
      loginStore.loginSubmitHandler(loginPresentable)
    }

    return {
      submitHandler,
      loginPresentable
    }
  }
})
</script>

<style scoped>
/* Your component's styles here */
</style>
