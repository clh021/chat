<template>
  <div class="messages flex-1 overflow-auto">
    <div
      v-for="(msg, index) in modelValue"
      :class="`message mb-4 flex${msg.type == 'SEND' ? ' me text-right' : ''}`"
      :key="`msg_${index}`"
    >
      <div class="flex-2" v-if="msg.type == 'RECEIVE'">
        <div v-if="msg.img" class="w-12 h-12 relative">
          <img class="w-12 h-12 rounded-full mx-auto" :src="msg.img" alt="chat-user" />
          <span
            class="absolute w-4 h-4 bg-gray-400 rounded-full right-0 bottom-0 border-2 border-white"
          ></span>
        </div>
      </div>
      <div class="flex-1 px-2">
        <div
          :class="`inline-block rounded-full p-2 px-6 ${
            msg.type == 'RECEIVE' ? 'bg-blue-600 text-white' : 'bg-gray-300 text-gray-700'
          }`"
        >
          <span>{{ msg.content }}</span>
        </div>
        <div v-if="msg.time" class="pl-4">
          <small class="text-gray-500">{{ msg.time }}</small>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ['modelValue'],
  data() {
    return {
      // messages: [
      //   {
      //     isSelf: false,
      //     // img: '/twc/resources/profile-image.png',
      //     content: '可以邀请我的家人们都来试试吗？',
      //     time: '15 April'
      //   },
      //   { isSelf: true, content: '好呀！', time: '15 April' },
      //   {
      //     isSelf: false,
      //     // img: '/twc/resources/profile-image.png',
      //     content: '啊哈！太谢谢你啦！',
      //     time: '15 April'
      //   },
      //   { isSelf: true, content: '不客气哈！:D', time: '15 April' }
      // ]
    }
  }
}
</script>
