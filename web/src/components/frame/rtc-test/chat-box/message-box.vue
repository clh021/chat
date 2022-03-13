<template>
  <div class="messages flex-1 overflow-auto">
    <div
      :class="`message mb-4 flex${msg.isSelf ? ' me text-right' : ''}`"
      v-for="(msg, index) in messages"
      :key="`msg_${index}`"
    >
      <div class="flex-2" v-if="!msg.isSelf">
        <div class="w-12 h-12 relative">
          <img class="w-12 h-12 rounded-full mx-auto" :src="msg.img" alt="chat-user" />
          <span
            class="absolute w-4 h-4 bg-gray-400 rounded-full right-0 bottom-0 border-2 border-white"
          ></span>
        </div>
      </div>
      <div class="flex-1 px-2">
        <div
          :class="`inline-block rounded-full p-2 px-6 ${
            msg.isSelf ? 'bg-blue-600 text-white' : 'bg-gray-300 text-gray-700'
          }`"
        >
          <span>{{ msg.content }}</span>
        </div>
        <div class="pl-4"><small class="text-gray-500">15 April</small></div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      messages: [
        {
          isSelf: false,
          img: '/twc/resources/profile-image.png',
          content: '可以邀请我的家人们都来试试吗？'
        },
        { isSelf: true, content: '好呀！' },
        {
          isSelf: false,
          img: '/twc/resources/profile-image.png',
          content: '啊哈！太谢谢你啦！'
        },
        { isSelf: true, content: '不客气哈！:D' }
      ]
    }
  }
}
</script>
