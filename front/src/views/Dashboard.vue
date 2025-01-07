<template>
  <div class="dashboard">
    <div v-if="error">
      {{ error }}
    </div>
    <h1>Boite mail : {{ mail }}</h1>
    <div class="container" v-if="messages.length">
      <div>
        <h3>{{ filters.length }} filtres</h3>
        <form @submit.prevent="addFilter">
          <div class="form-group">
            <label for="newFilter">Nouveau filtre</label>
            <input type="text" id="newFilter" v-model="newFilter.name" required />
          </div>

          <button type="submit">Ajouter</button>
        </form>
        <div v-for="filter in filters" :key="filter.name" class="message-card">
          <p>{{ filter.name }}</p>
        </div>
      </div>
      <div>
        <h3>{{ messages.length }} mails non lus</h3>
        <div v-for="message in messages" :key="message.Id" class="message-card">
          <div v-for="expeditor in message.Mails" :key="expeditor.PersonalName" class="from-card">
            <p>
              {{ expeditor.PersonalName }} ( {{ expeditor.MailboxName }}@{{ expeditor.HostName }} )
            </p>
          </div>
          <p><strong>Sujet:</strong> {{ message.Subject }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, type PropType, onMounted, reactive } from 'vue'
import { type Message } from '@/domain/message'
import { type Filter } from '@/domain/filter'
import { UnseenMessageLoader } from '@/dataprovider/unseenMessagesLoader'

export default defineComponent({
  // eslint-disable-next-line vue/multi-word-component-names
  name: 'Dashboard',
  props: {
    email: {
      type: String as PropType<string>,
      required: true,
    },
  },
  setup(props) {
    const error = ref<string | null>(null)
    const mail = ref<string | null>(props.email)
    const messages = ref<Message[]>([])
    const filters = ref<Filter[]>([])
    const isLoading = ref<boolean>(false)
    const newFilter = reactive<Filter>({
      name: '',
      usage: 0,
    })

    const fetchMessages = async (): Promise<void> => {
      const datas = await UnseenMessageLoader.load(props.email)
      messages.value = datas
    }
    const addFilter = async () => {
      filters.value.push(<Filter>{ name: newFilter.name, usage: newFilter.usage })
      newFilter.name = ''
      newFilter.usage = 0
    }

    onMounted(fetchMessages)

    return {
      mail,
      error,
      messages,
      filters: filters,
      isLoading,
      newFilter,
      addFilter,
    }
  },
})
</script>
