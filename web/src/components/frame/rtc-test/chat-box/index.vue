<template>
  <div class="chat-area flex-1 flex flex-col">
    <div v-if="selfName">
      <Alert
        v-if="!modelValue"
        color="warning"
        text="提示: 填写您朋友的名字，试着联系TA吧！"
        class="mb-4"
        @click="noticeFriendNameHandle"
      ></Alert>
      <Title ref="title" v-model="modelValue" :selfName="selfName"></Title>
    </div>
    <div v-else>
      <Alert color="warning" text="提示: 先设置好一个名字，方便朋友找到您" class="mb-4"></Alert>
    </div>
    <MessageBox v-model="Messages"></MessageBox>
    <div class="flex-2 pt-4 pb-10">
      <InputBox @sendMsg="sendMsgHandle"></InputBox>
    </div>
  </div>
</template>
<script>
import rtc from '@/libs/rtc.js'
import Alert from '../components/alert.vue'
import Title from './title.vue'
import MessageBox from './message-box.vue'
import InputBox from './input-box.vue'

export default {
  components: { Alert, Title, MessageBox, InputBox },
  props: ['modelValue', 'selfName'],
  emits: ['update:modelValue', 'stateChange', 'openChatList'],
  data() {
    return {
      Messages: [
        // {
        //   type: 'RECEIVE',
        //   img: '/twc/resources/profile-pingli.png',
        //   content: '可以邀请我的家人们都来试试吗？',
        //   time: '15 April'
        // },
        // { type: 'SEND', content: '好呀！', time: '15 April' },
        // {
        //   type: 'RECEIVE',
        //   img: '/twc/resources/profile-pingli.png',
        //   content: '啊哈！太谢谢你啦！',
        //   time: '15 April'
        // },
        // { type: 'SEND', content: '不客气哈！:D', time: '15 April' },
        // { type: 'SYS', content: '对方已离线', time: '15 April' }
      ]
    }
  },
  created() {
    this.$watch('modelValue', (newVal, oldVal) => {
      console.log('modelValue:val:', newVal, oldVal)
      let dataConnection = this.rtc.connect(newVal)
      console.log(this.rtc.connections())
      dataConnection.send('发送一个消息试试')
    })
    this.$watch('selfName', (newVal, oldVal) => {
      console.log('selfName:val:', newVal, oldVal)
      this.rtcInit(newVal, oldVal)
    })
  },
  methods: {
    noticeFriendNameHandle() {
      // this.$refs.title.friendFocus() // 聚焦子组件输入框，用于填写朋友名字
      this.$emit('openChatList') // 打开侧边朋友列表，由朋友列表界面负责朋友的维护和选择
    },
    sendMsgHandle(msg) {
      this.triggerMsgHandle(msg, 'SEND')
    },
    triggerMsgHandle(msg, type) {
      this.Messages.push({
        type: type,
        content: msg
        // time: '15 April'
      })
    },
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
    rtcInit(newSelfName, oldSelfName) {
      console.log('selfName:val:', newSelfName, oldSelfName)
      if (newSelfName) {
        this.rtc = new rtc(newSelfName)
        this.rtc.handles(this.rtcHandles())
      }
    },
    rtcHandles() {
      return {
        // register成功的回调
        open: (id) => {
          console.log('rtc:open:')
          console.log(id)
          if (id == 'unavailable-id') {
            console.log('名字已经被暂时使用了')
            this.triggerMsgHandle('名字已经被暂时使用了', 'SYS')
          }
          this.$emit('stateChange', true)
        },
        // 被连接的事件
        connection: (dataConnection) => {
          console.log('rtc:connection')
          // 收到消息
          dataConnection.on('data', function (data) {
            this.triggerMsgHandle('收到了消息', 'SYS')
            console.log('Received', data) // is serialized by BinaryPack by default and sent to the remote peer.
            this.triggerMsgHandle(data, 'Receive')
          })
          // 发送消息
          console.log('data.send: 哈哈，有新伙伴连接我了。')
          dataConnection.send('哈哈，有新伙伴连接我了。') // send any type of data, including objects, strings, and blobs.
          this.triggerMsgHandle('有新伙伴连接你了', 'SYS')
        },
        // 被 call 的事件
        call: (mediaConnection) => {
          console.log('rtc:call')
          navigator.mediaDevices.getUserMedia(
            { video: true, audio: true },
            (stream) => {
              mediaConnection.answer(stream) // Answer the mediaConnection with an A/V stream.
              mediaConnection.on('stream', (remoteStream) => {
                // Show stream in some <video> element.
              })
            },
            (err) => {
              console.error('Failed to get local stream', err)
            }
          )
        },
        // 被 close 的事件
        close: () => {
          console.log('rtc:close')
        },
        // disconnected 的事件
        disconnected: () => {
          console.log('rtc:disconnected')
          this.$emit('stateChange', false)
          this.triggerMsgHandle('当前已离线', 'SYS')
        },
        // error 的事件
        error: (err) => {
          console.log(err.type)
        }
      }
    },
    rtcCall() {
      let friendName = modelValue
      navigator.mediaDevices.getUserMedia(
        { video: true, audio: true },
        (stream) => {
          const call = this.rtc.call(friendName, stream)
          call.on('stream', (remoteStream) => {
            // Show stream in some <video> element.
          })
        },
        (err) => {
          console.error('Failed to get local stream', err)
        }
      )
    }
  },
  mounted() {
    // TODO: 发送消息，接收消息
    // TODO: 发送视频请求，接受适配请求
    // TODO: 建立连接的过程可以搞一个系统日志打开
    // TODO: 所有建立连接的过程都放到系统日志里面去
    // TODO: 日志中需要较为方便的看出是否是 p2p 连接
    // TODO: 方便转发或者别的方式记录收藏内容
    // this.rtcInit()
  }
}
</script>
