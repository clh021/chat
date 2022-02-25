<script setup lang="ts">
import {ref} from 'vue'

const count = ref(0)
</script>

<template>
  <div class="w-full h-screen">
    <div class="flex h-full">
      <router-link to="/">Go to root</router-link>
      .
      <router-link to="/home">Go to Home</router-link>
      .
      <router-link to="/about">Go to About</router-link>
      <br />
      <button type="button" @click="count++">count is: {{ count }}</button>
    </div>
  </div>
</template>

<style scoped>
a {
  color: #42b983;
}

label {
  margin: 0 0.5em;
  font-weight: bold;
}

code {
  background-color: #eee;
  padding: 2px 4px;
  border-radius: 4px;
  color: #304455;
}
</style>
