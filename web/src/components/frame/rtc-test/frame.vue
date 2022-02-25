<template>
  <div class="w-full h-screen">
    <div class="flex h-full">
      <!-- <ChatMenu></ChatMenu> -->
      <div class="flex-1 bg-gray-100 w-full h-full">
        <div class="main-body container m-auto w-11/12 h-full flex flex-col">
          <ChatTopbar></ChatTopbar>
          <div class="main flex-1 flex flex-col">
            <div class="hidden lg:block heading flex-2">
              <h1 class="text-3xl text-gray-700 mb-4">Chat</h1>
            </div>

            <div class="flex-1 flex h-full">
              <ChatList></ChatList>
              <ChatBox></ChatBox>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import ChatTopbar from './chat-topbar.vue'
import ChatList from './chat-list.vue'
import ChatBox from './chat-box/index.vue'
// import ChatMenu from './chat-menu.vue'

export default {
  components: {
    ChatTopbar,
    ChatList,
    ChatBox
    // ChatMenu
  }
}
</script>

<style scoped></style>
