<template>
  <div class="message mb-4 flex sys text-center">
    <div class="flex-1 px-2">
      <div class="text-sm text-gray-400">
        <span>{{ msg.content }}</span>
      </div>
      <!--div v-if="msg.time" class="pl-4">
        <small class="text-gray-500">{{ msg.time }}</small>
      </div-->
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
