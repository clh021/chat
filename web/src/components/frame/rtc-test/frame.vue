<template>
  <div class="w-full h-screen">
    <!-- rtcSupport from @/libs/utils.js -->
    <div class="flex h-full" v-if="rtcSupport()">
      <!-- <ChatMenu></ChatMenu> -->
      <div class="flex-1 w-full h-full">
        <div class="main-body container m-auto w-11/12 h-full flex flex-col">
          <ChatTopbar></ChatTopbar>
          <div class="main flex-1 flex flex-col">
            <div class="hidden lg:block heading flex-2">
              <h1 class="text-3xl text-gray-700 mb-4">Chat</h1>
            </div>

            <div class="flex-1 flex h-full">
              <!-- <ChatList></ChatList> -->
              <ChatBox></ChatBox>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="flex h-full justify-center items-center" v-else>
      <div class="text-3xl">当前设备无法支持</div>
    </div>
  </div>
</template>

<script>
import rtc from '@/libs/rtc.js'
import ChatTopbar from './chat-topbar.vue'
// import ChatList from './chat-list.vue'
import ChatBox from './chat-box/index.vue'
// import ChatMenu from './chat-menu.vue'

export default {
  components: {
    ChatTopbar,
    ChatBox
    // ChatList,
    // ChatMenu
  },
  methods: {
    rtcSupport() {
      return rtc.isSupport()
    }
  }
}
</script>

<style scoped></style>
