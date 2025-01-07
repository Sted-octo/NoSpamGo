<template>
  <div class="app">
    <header>
      <img src="/nospam.svg" alt="Logo" class="logo" />
    </header>
    <router-view></router-view>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'App',
})
</script>

<style>
.app {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

.logo {
  height: 50px;
  margin-bottom: 20px;
}
</style>
