<template>
  <div class="message mb-4 flex me text-right">
    <div class="flex-1 px-2">
      <div class="inline-block rounded-full p-2 px-6 bg-gray-300 text-gray-700">
        <span>{{ msg.content }}</span>
      </div>
      <div v-if="msg.time" class="pl-4">
        <small class="text-gray-500">{{ msg.time }}</small>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ['msg']
}
</script>

<style>
</style>
