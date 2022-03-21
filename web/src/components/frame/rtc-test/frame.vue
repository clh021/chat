<template>
  <div>
    <div class="drawer w-full h-screen" v-if="rtcSupport()">
      <input id="drawer-chat-list" type="checkbox" class="drawer-toggle" v-model="ChatListToggle" />
      <div class="drawer-content">
        <!-- Page content here -->
        <!-- rtcSupport from @/libs/utils.js -->
        <!-- <ChatMenu></ChatMenu> -->
        <div class="flex-1 w-full h-full">
          <div class="main-body container m-auto w-11/12 h-full flex flex-col">
            <ChatTopbar
              @openChatList="ChatListToggle = true"
              @openSetting="SettingToggle = true"
            ></ChatTopbar>
            <div class="main flex-1 flex flex-col">
              <!--div class="hidden lg:block heading flex-2">
                <h1 class="text-3xl text-gray-700 mb-4">Chat</h1>
              </div-->
              <div class="flex-1 flex h-full">
                <ChatBox></ChatBox>
                <Setting v-model="SettingToggle" @newName="applyNewNameHandle"></Setting>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="drawer-side">
        <label for="drawer-chat-list" class="drawer-overlay"></label>
        <ul class="menu p-4 overflow-y-auto w-96 bg-base-100 text-base-content">
          <!-- Sidebar content here -->
          <ChatList></ChatList>
        </ul>
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
import ChatList from './chat-list-simple.vue'
import ChatBox from './chat-box/index.vue'
// import ChatMenu from './chat-menu.vue'
import Setting from './setting.vue'

export default {
  components: {
    ChatTopbar,
    ChatBox,
    ChatList,
    // ChatMenu,
    Setting
  },
  data() {
    return {
      ChatListToggle: false,
      SettingToggle: false
    }
  },
  methods: {
    rtcSupport() {
      return rtc.isSupport()
    },
    openChatListHandle(v) {
      console.log('openChatListHandle', v)
      this.ChatListToggle = true
    },
    applyNewNameHandle(selfName) {
      console.log('selfName:', selfName)
    }
  }
}
</script>

<style scoped></style>
