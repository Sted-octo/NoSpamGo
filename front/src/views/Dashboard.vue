<template>
  <div class="dashboard">
    <div v-if="error">
      {{ error }}
    </div>

    <div class="fixed-section container">
      <div>
        <h1>Boite mail : {{ mail }}</h1>
        <h3>{{ messages.length }} mails non lus (dont {{ spamCount }} spams)</h3>
        <h3>{{ filters.length }} filtres</h3>
        <form @submit.prevent="addFilter">
          <div class="form-group">
            <label for="newFilter">Nouveau filtre</label>
            <input type="text" id="newFilter" v-model="newFilter.Name" required />
          </div>
          <button type="submit">Ajouter</button>
        </form>
      </div>
      <div>
        <button type="button" @click="applyFilters">Appliquer les filtres</button>
        <button type="button" @click="fetchMessages">Recharger</button>
      </div>
    </div>

    <div class="container scrollable-section" v-if="messages.length">
      <div>
        <h3>{{ filters.length }} filtres</h3>

        <div v-for="filter in filterCounts" :key="filter.Name" class="message-card">
          <p>{{ filter.Name }} ({{ filter.spamCount }} messages détectés)</p>
        </div>
      </div>
      <div>
        <h3>{{ messages.length }} mails non lus (dont {{ spamCount }} spams)</h3>
        <div
          v-for="message in messages"
          :key="message.Id"
          :class="['message-card', { 'is-spam': isMessageSpam(message) }]"
        >
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
import { defineComponent, ref, type PropType, onMounted, reactive, computed } from 'vue'
import { type Message } from '@/domain/message'
import { type Filter } from '@/domain/filter'
import { UnseenMessageLoader } from '@/dataprovider/unseenMessagesLoader'
import { FiltesSaver } from '@/dataprovider/filtersSaver'
import { FiltersGetter } from '@/dataprovider/filtersGetter'

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
    const mail = ref<string>(props.email)
    const messages = ref<Message[]>([])
    const filters = ref<Filter[]>([])
    const filtersToSave = ref<Filter[]>([])
    const hasNewFilters = ref<boolean>(false)
    const newFilter = reactive<Filter>({
      Name: '',
      NumberOfSpamDetected: 0,
    })

    const fetchMessages = async (): Promise<void> => {
      const datas = await UnseenMessageLoader.load(props.email)

      messages.value = datas
    }

    const fetchFilters = async (): Promise<void> => {
      const datas = await FiltersGetter.get(props.email)

      filters.value = datas
    }

    const addFilter = async () => {
      if (newFilter.Name === '') return
      hasNewFilters.value = true
      filters.value.push(<Filter>{
        Name: newFilter.Name.toLowerCase(),
        NumberOfSpamDetected: newFilter.NumberOfSpamDetected,
      })
      filtersToSave.value.push(<Filter>{
        Name: newFilter.Name.toLowerCase(),
        NumberOfSpamDetected: newFilter.NumberOfSpamDetected,
      })
      newFilter.Name = ''
      newFilter.NumberOfSpamDetected = 0
    }
    const applyFilters = async () => {
      if (hasNewFilters.value) {
        const state = await FiltesSaver.save(filtersToSave.value, mail.value)
        if (state.saved) {
          filtersToSave.value = []
          hasNewFilters.value = false
        }
      }
    }

    const isMessageSpam = (message: Message) => {
      if (filters.value.length === 0) return
      return filters.value.some((filter) => isMessageSpamByFilter(message, filter))
    }

    const compareText = (text: string, filter: string): boolean => {
      if (text === undefined || filter === undefined) return false

      const cleanText = text.toLowerCase().trim()
      const cleanFilter = filter.toLowerCase().trim()
      return cleanText.includes(cleanFilter) || cleanText === cleanFilter
    }

    const isMessageSpamByFilter = (message: Message, filter: Filter) => {
      return (
        message.Mails.some(
          (origin) =>
            compareText(origin.PersonalName, filter.Name) ||
            compareText(origin.HostName, filter.Name) ||
            compareText(origin.MailboxName, filter.Name),
        ) || compareText(message.Subject, filter.Name)
      )
    }

    const spamCount = computed(() => {
      return messages.value.filter((message) => isMessageSpam(message)).length
    })

    const filterCounts = computed(() => {
      return filters.value.map((filter) => ({
        ...filter,
        spamCount: messages.value.filter((message) => isMessageSpamByFilter(message, filter))
          .length,
      }))
    })

    onMounted(() => {
      fetchMessages()
      fetchFilters()
    })

    return {
      mail,
      error,
      messages,
      filters: filters,
      newFilter,
      addFilter,
      isMessageSpam,
      spamCount,
      filterCounts,
      applyFilters,
      fetchMessages,
    }
  },
})
</script>

<style scoped>
.dashboard {
  position: relative;
  padding-top: 200px; /* Ajustez selon la hauteur de votre section fixe */
}

.fixed-section {
  position: sticky;
  top: 0;
  background-color: white;
  padding: 1rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  z-index: 100;
}

.scrollable-section {
  /* Le contenu défilant */
  position: relative;
}

form {
  display: flex;
  align-items: center;
  gap: 1rem; /* espace entre les éléments */
  vertical-align: middle;
}

.form-group {
  align-items: center;
  gap: 1rem;
  margin-bottom: 1rem;
}
/* Ajoutez d'autres styles selon vos besoins */
</style>
