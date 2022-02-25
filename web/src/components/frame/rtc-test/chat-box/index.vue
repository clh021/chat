<template>
  <div class="chat-area flex-1 flex flex-col">
    <Title></Title>
    <MessageBox></MessageBox>
    <div class="flex-2 pt-4 pb-10">
      <InputBox></InputBox>
    </div>
  </div>
</template>
<script lang="ts">
import Title from './title.vue'
import MessageBox from './message-box.vue'
import InputBox from './input-box.vue'

export default {
  components: { Title, MessageBox, InputBox }
}
</script>
