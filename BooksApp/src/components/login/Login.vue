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
            v-model="loginPresentable.email"
          />
          <TextInput
            label="Password"
            type="password"
            name="password"
            :required="true"
            v-model="loginPresentable.password"
          />

          <hr />
          <input type="submit" class="btn btn-primary" value="Login" />
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

    function submitHandler(event: any) {
      console.log('submitHandler', event)
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
