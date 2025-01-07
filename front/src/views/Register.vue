<template>
  <div class="register-form">
    <div v-if="error">
      {{ error }}
    </div>
    <h1>Inscription</h1>
    <form @submit.prevent="setup2FA">
      <div class="form-group">
        <label for="email">Email</label>
        <input type="email" id="email" v-model="email" required />
      </div>
      <button type="submit" :disabled="isLoading">Obtenir le QR Code</button>
    </form>

    <div v-if="qrCodeImageUrl" class="qr-section">
      <img :src="qrCodeImageUrl" alt="QR Code" />
      <div class="validation-section">
        <input type="text" v-model="authCode" placeholder="Code de validation" />
        <button @click="verifyCode" :disabled="isValidating">Valider</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { useRouter } from 'vue-router'
import { Auth2FactorSetuper } from '@/dataprovider/auth2FactorSetuper'
import { Auth2FactorVerificator } from '@/dataprovider/auth2FactorVerificator'
import QRCode from 'qrcode'

export default defineComponent({
  // eslint-disable-next-line vue/multi-word-component-names
  name: 'Register',
  setup() {
    const router = useRouter()
    const email = ref<string>('')
    const authCode = ref<string>('')
    const isLoading = ref<boolean>(false)
    const isValidating = ref<boolean>(false)
    const error = ref<string | null>(null)
    const qrCodeUrl = ref('')
    const qrCodeImageUrl = ref('')
    const verificationResult = ref<boolean | null>(null)

    const setup2FA = async () => {
      isLoading.value = true
      error.value = null

      try {
        const data = await Auth2FactorSetuper.setup2FA(email.value)
        qrCodeUrl.value = data.qr_code
        qrCodeImageUrl.value = await QRCode.toDataURL(data.qr_code)
      } catch (err) {
        error.value = 'Erreur lors de la configuration 2FA'
        console.error(err)
      } finally {
        isLoading.value = false
      }
    }

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
      isValidating,
      setup2FA,
      verifyCode,
      error,
      qrCodeUrl,
      qrCodeImageUrl,
    }
  },
})
</script>
