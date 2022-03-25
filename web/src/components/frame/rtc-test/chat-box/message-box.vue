<template>
  <div class="messages flex-1 overflow-auto">
    <div v-for="(msg, index) in modelValue" :key="`msg_${index}`">
      <Send v-if="msg.type == 'SEND'" v-bind:msg="msg"></Send>
      <Receive v-else-if="msg.type == 'RECEIVE'" v-bind:msg="msg"></Receive>
      <Sys v-else-if="msg.type == 'SYS'" v-bind:msg="msg"></Sys>
    </div>
  </div>
</template>

<script>
import Send from './message/send.vue'
import Receive from './message/receive.vue'
import Sys from './message/sys.vue'
export default {
  components: { Send, Receive, Sys },
  props: ['modelValue'],
  data() {
    return {
      messages: [
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
  }
}
</script>
