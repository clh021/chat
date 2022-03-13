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
    // TODO: 名字可以修改，一旦修改，就重新建立连接
    // TODO: 建立连接的部分独立为一个方法来调用
    // TODO: 发送消息，接收消息
    // TODO: 发送视频请求，接受适配请求
    // TODO: 建立连接的过程可以搞一个系统日志打开
    // TODO: 所有建立连接的过程都放到系统日志里面去
    console.log('mounted.')
    this.rtc = new rtc('test')
    this.rtc.handles(this.openHandle, this.connHandle)
  }
}
</script>
