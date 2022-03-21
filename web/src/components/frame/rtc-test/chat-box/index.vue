<template>
  <div class="chat-area flex-1 flex flex-col">
    <Title
      v-model="modelValue"
      v-if="selfName"
      :selfName="selfName"
      @update:modelValue="initRtc()"
    ></Title>
    <Alert
      color="warning"
      text="提示: 先设置好一个名字，方便朋友找到您"
      v-else
      class="mb-4"
    ></Alert>
    <MessageBox></MessageBox>
    <div class="flex-2 pt-4 pb-10">
      <InputBox></InputBox>
    </div>
  </div>
</template>
<script>
import rtc from '@/libs/rtc.js'
import Alert from '../components/alert.vue'
import Title from './title.vue'
import MessageBox from './message-box.vue'
import InputBox from './input-box.vue'
import { watch } from '@vue/runtime-core'

export default {
  components: { Alert, Title, MessageBox, InputBox },
  props: ['modelValue', 'selfName'],
  emits: ['update:modelValue'],
  data() {
    return {
      // friendName: '宇宙空间'
    }
  },
  created() {
    this.$watch('modelValue', (newVal, oldVal) => {
      console.log('modelValue:val:', newVal, oldVal)
    })
    this.$watch('selfName', (newVal, oldVal) => {
      console.log('selfName:val:', newVal, oldVal)
      this.initRtc(newVal, oldVal)
    })
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
    initRtc(newSelfName, oldSelfName) {
      console.log('selfName:val:', newSelfName, oldSelfName)
      if (newSelfName) {
        this.rtc = new rtc(newSelfName)
        this.rtc.handles(this.openHandle, this.connHandle)
      }
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
