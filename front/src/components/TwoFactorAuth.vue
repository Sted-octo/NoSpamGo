<template>
  <div class="container mx-auto p-4">
    <div
      v-if="error"
      class="max-w-md mx-auto mb-4 bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded"
    >
      {{ error }}
    </div>

    <div class="mb-4">
      <label class="block text-gray-700 mb-2">Email / Nom d'utilisateur</label>
      <input
        v-model="username"
        type="email"
        class="w-full px-3 py-2 border rounded-lg"
        placeholder="utilisateur@exemple.com"
        :disabled="isLoading"
      />
    </div>

    <div v-if="!setupComplete" class="max-w-md mx-auto bg-white rounded-xl shadow-md p-6">
      <h2 class="text-2xl font-bold mb-4">Configuration 2FA</h2>

      <button
        @click="setup2FA"
        class="w-full bg-blue-500 text-white py-2 rounded-lg hover:bg-blue-600 disabled:opacity-50"
        :disabled="!username || isLoading"
      >
        <span v-if="isLoading">Chargement...</span>
        <span v-else>Configurer 2FA</span>
      </button>

      <div class="mt-6">
        <h3 class="text-lg font-semibold mb-3">Scannez ce QR Code</h3>
        <div class="flex justify-center mb-4">
          <img :src="qrCodeImageUrl" alt="QR Code pour 2FA" class="border p-2" />
        </div>
        <p class="text-sm text-gray-600 mb-4">
          Scannez ce QR code avec Microsoft Authenticator pour commencer.
        </p>
        <button
          @click="confirmSetup"
          class="w-full bg-green-500 text-white py-2 rounded-lg hover:bg-green-600"
          :disabled="isLoading"
        >
          J'ai scanné le QR Code
        </button>
      </div>
    </div>

    <div v-else class="max-w-md mx-auto bg-white rounded-xl shadow-md p-6">
      <h2 class="text-2xl font-bold mb-4">Vérification 2FA</h2>

      <div class="mb-4">
        <label class="block text-gray-700 mb-2">Code d'authentification</label>
        <input
          v-model="verificationCode"
          type="text"
          class="w-full px-3 py-2 border rounded-lg"
          placeholder="Entrez le code à 6 chiffres"
          maxlength="6"
          :disabled="isLoading"
        />
      </div>

      <button
        @click="verifyCode"
        class="w-full bg-blue-500 text-white py-2 rounded-lg hover:bg-blue-600 disabled:opacity-50"
        :disabled="!verificationCode || isLoading"
      >
        <span v-if="isLoading">Vérification...</span>
        <span v-else>Vérifier le code</span>
      </button>

      <div v-if="verificationResult !== null" class="mt-4">
        <p :class="verificationResult ? 'text-green-600' : 'text-red-600'">
          {{ verificationResult ? 'Code valide !' : 'Code invalide, veuillez réessayer.' }}
        </p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { ref } from 'vue'
import { Auth2FactorSetuper } from '@/dataprovider/auth2FactorSetuper'
import { Auth2FactorVerificator } from '@/dataprovider/auth2FactorVerificator'
import QRCode from 'qrcode'

export default defineComponent({
  name: 'TwoFactorAuth',
  setup() {
    const username = ref('')
    const qrCodeUrl = ref('')
    const qrCodeImageUrl = ref('')
    const setupComplete = ref(false)
    const verificationCode = ref('')
    const verificationResult = ref<boolean | null>(null)
    const isLoading = ref(false)
    const error = ref<string | null>(null)

    const setup2FA = async () => {
      isLoading.value = true
      error.value = null

      try {
        const data = await Auth2FactorSetuper.setup2FA(username.value)
        qrCodeUrl.value = data.qr_code
        qrCodeImageUrl.value = await QRCode.toDataURL(data.qr_code)
      } catch (err) {
        error.value = 'Erreur lors de la configuration 2FA'
        console.error(err)
      } finally {
        isLoading.value = false
      }
    }

    const confirmSetup = () => {
      setupComplete.value = true
    }

    const verifyCode = async () => {
      isLoading.value = true
      error.value = null

      try {
        const result = await Auth2FactorVerificator.verify2FA({
          mail: username.value,
          token: verificationCode.value,
        })

        verificationResult.value = result.valid

        if (result.valid) {
          setTimeout(() => {
            verificationResult.value = null
            verificationCode.value = ''
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
      username,
      qrCodeUrl,
      qrCodeImageUrl,
      setupComplete,
      verificationCode,
      verificationResult,
      isLoading,
      error,
      setup2FA,
      confirmSetup,
      verifyCode,
    }
  },
})
</script>
