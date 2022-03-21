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
                <ChatBox v-model="friendName" :selfName="selfName"></ChatBox>

                <!-- modal-box begin -->
                <!-- modal 想办法，一定可以封装成为一个组件的 -->
                <input
                  type="checkbox"
                  id="my-modal-5"
                  class="modal-toggle"
                  v-model="SettingToggle"
                />
                <div class="modal">
                  <!--  w-11/12 max-w-5xl -->
                  <div :class="`modal-box${SettingToggle ? ' modal-open' : ''}`">
                    <Setting v-model="selfName" @change="selfNameChangeHandle"></Setting>
                  </div>
                </div>
                <!-- modal-box end -->
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
      SettingToggle: false,
      Settings: [],
      friendName: '小陀螺',
      selfName: ''
    }
  },
  mounted() {
    this.checkSelfName()
  },
  methods: {
    selfNameChangeHandle() {
      console.log('selfNameChangeHandle')
      this.SettingToggle = true
    },
    checkSelfName() {
      if (this.selfName.length <= 0) {
        this.SettingToggle = true
      }
    },
    rtcSupport() {
      return rtc.isSupport()
    },
    openChatListHandle(v) {
      this.ChatListToggle = true
    }
  }
}
</script>

<style scoped></style>
