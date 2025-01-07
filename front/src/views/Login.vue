<template>
  <div class="login-form">
    <div v-if="error">
      {{ error }}
    </div>
    <h1>Connexion</h1>
    <form @submit.prevent="verifyCode">
      <div class="form-group">
        <label for="email">Email</label>
        <input type="email" id="email" v-model="email" required />
      </div>
      <div class="form-group">
        <label for="code">Code Microsoft Authenticator</label>
        <input type="text" id="code" v-model="authCode" required />
      </div>
      <button type="submit" :disabled="isLoading">Se connecter</button>
    </form>
    <router-link to="/register">S'inscrire</router-link>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { useRouter } from 'vue-router'
import { Auth2FactorVerificator } from '@/dataprovider/auth2FactorVerificator'

export default defineComponent({
  // eslint-disable-next-line vue/multi-word-component-names
  name: 'Login',
  setup() {
    const router = useRouter()
    const email = ref<string>('')
    const authCode = ref<string>('')
    const isLoading = ref<boolean>(false)
    const verificationResult = ref<boolean | null>(null)
    const error = ref<string | null>(null)

    const verifyCode = async () => {
      isLoading.value = true
      error.value = null

      try {
        const result = await Auth2FactorVerificator.verify2FA({
          mail: email.value,
          token: authCode.value,
        })

        verificationResult.value = result.valid

        if (result.valid) {
          setTimeout(() => {
            verificationResult.value = null
            authCode.value = ''
            router.push(`/email-config/${encodeURIComponent(email.value)}`)
          }, 2000)
        }
      } catch (err) {
        error.value = 'Erreur lors de la vérification du code'
        console.error(err)
      } finally {
        isLoading.value = false
      }
    }

    return {
      email,
      authCode,
      isLoading,
      verifyCode,
      error,
    }
  },
})
</script>
