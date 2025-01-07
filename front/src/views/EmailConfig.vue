<template>
  <div class="email-config">
    <div v-if="error">
      {{ error }}
    </div>
    <h1>Configuration Email</h1>
    <form @submit.prevent="saveConfig">
      <div class="form-group">
        <label for="mail">Mail</label>
        <input type="text" id="mail" v-model="config.mail" readonly />
      </div>
      <div class="form-group">
        <label for="username">Nom d'utilisateur</label>
        <input type="text" id="username" v-model="config.username" required />
      </div>
      <div class="form-group">
        <label for="password">Mot de passe</label>
        <input type="password" id="password" v-model="config.password" required />
      </div>
      <div class="form-group">
        <label for="server">Serveur IMAP</label>
        <input type="text" id="server" v-model="config.server" required />
      </div>
      <div class="form-group">
        <label for="port">Port</label>
        <input type="number" id="port" v-model="config.port" required />
      </div>
      <button type="submit" :disabled="isSaving">Sauvegarder</button>
    </form>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, type PropType } from 'vue'
import { useRouter } from 'vue-router'
import type { MailConfig } from '@/domain/mailConfig'
import { MailAccessUpdator } from '@/dataprovider/mailAccessUpdator'

export default defineComponent({
  name: 'EmailConfig',
  props: {
    email: {
      type: String as PropType<string>,
      required: true,
    },
  },
  setup(props) {
    const router = useRouter()
    const isSaving = ref<boolean>(false)
    const config = reactive<MailConfig>({
      mail: props.email,
      username: '',
      password: '',
      server: '',
      port: 993,
    })
    const error = ref<string | null>(null)

    const saveConfig = async (): Promise<void> => {
      isSaving.value = true
      error.value = null

      try {
        const result = await MailAccessUpdator.updateMailConfig(config)

        if (result.saved) {
          setTimeout(() => {
            router.push('/dashboard')
          }, 2000)
        }
      } catch (err) {
        error.value = 'Erreur lors de la sauvegarde des la configuration'
        console.error(err)
      } finally {
        isSaving.value = false
      }
    }

    return {
      config,
      isSaving,
      saveConfig,
      error,
    }
  },
})
</script>
