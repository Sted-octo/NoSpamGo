<template>
  <div class="dashboard">
    <div v-if="error">
      {{ error }}
    </div>
    <h1>Tableau de bord de {{mail}}</h1>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref,  type PropType } from 'vue'

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

    return {
      mail,
      error,
    }
  },
})
</script>
