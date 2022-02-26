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
import { reactive, onMounted } from 'vue'
import Title from './title.vue'
import MessageBox from './message-box.vue'
import InputBox from './input-box.vue'

export default {
  components: { Title, MessageBox, InputBox },
  setup() {
    const state = reactive({ name: 'Vue' })
    onMounted(() => {
      console.log('[onMounted]')

      function getRtcMedia() {
        const constraints = { video: true }
        function onSuccess(stream) {
          const video = document.querySelector('video')
          video.src = window.URL.createObjectURL(stream)
        }
        function onError(error) {
          console.log('navigator.getUserMedia error: ', error)
        }
        navigator.getUserMedia(constraints, onSuccess, onError)
      }
      function hasRtcApi() {
        navigator.getUserMedia =
          navigator.getUserMedia ||
          navigator.webkitGetUserMedia ||
          navigator.mozGetUserMedia ||
          navigator.msGetUserMedia
        if (navigator.getUserMedia) {
          console.log('支持')
          return true
        }
        console.log('不支持')
        return false
      }
      if (hasRtcApi()) {
        // getRtcMedia()
        // 如果支持则显示操作界面
      } else {
        // 如果不支持，直接全屏提示不支持
      }
    })
    return { state }
  }
}
</script>
