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
  data() {
    return {
      zoneTitle: '银河宇宙空间'
    }
  },
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
    initRtc() {
      this.rtc = new rtc(this.zoneTitle)
      this.rtc.handles(this.openHandle, this.connHandle)
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
    // TODO: 日志中需要较为方便的看出是否是 p2p 连接
    // TODO: 方便转发或者别的方式记录收藏内容
    // this.initRtc()
  }
}
</script>
