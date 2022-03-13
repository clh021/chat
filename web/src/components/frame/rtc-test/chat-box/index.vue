<template>
  <div class="chat-area flex-1 flex flex-col">
    <Title></Title>
    <MessageBox></MessageBox>
    <div class="flex-2 pt-4 pb-10">
      <InputBox></InputBox>
    </div>
  </div>
</template>
<script>
import rtc from '@/libs/rtc.js'
import Title from './title.vue'
import MessageBox from './message-box.vue'
import InputBox from './input-box.vue'

export default {
  components: { Title, MessageBox, InputBox },
  methods: {
    getRtcMedia() {
      const constraints = { video: true }
      navigator.getUserMedia(
        constraints,
        (stream) => {
          document.querySelector('video').src = window.URL.createObjectURL(stream)
        },
        (error) => {
          console.log('navigator.getUserMedia error: ', error)
        }
      )
    },
    openHandle(id) {
      console.log(`system: register success ${id}.`)
    },
    connHandle(conn) {
      conn.on('data', (data) => {
        console.log(JSON.parse(data))
      })
    }
  },
  mounted() {
    console.log('mounted.')
    this.rtc = new rtc('test')
    this.rtc.handles(this.openHandle, this.connHandle)
  }
}
</script>
