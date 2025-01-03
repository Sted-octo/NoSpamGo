<template>
  <div class="min-h-screen bg-gray-100">
    <TwoFactorAuth />
  </div>
</template>

<script lang="ts">
import TwoFactorAuth from './components/TwoFactorAuth.vue'

export default {
  name: 'App',
  components: {
    TwoFactorAuth
  }
}
</script>