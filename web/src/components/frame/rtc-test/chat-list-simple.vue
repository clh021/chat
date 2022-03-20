<template>
  <div class="sidebar flex flex-2 flex-col pr-6">
    <div class="search flex-2 pb-6 px-2">
      <input
        type="text"
        class="outline-none py-2 block w-full bg-transparent border-b-2 border-gray-200"
        placeholder="Search"
      />
    </div>
    <div class="flex-1 h-full overflow-auto px-2">
      <div
        class="
          entry
          cursor-pointer
          transform
          hover:scale-105
          duration-300
          transition-transform
          bg-white
          mb-4
          rounded
          p-4
          flex
          shadow-md
        "
        v-for="(friend, index) in friends"
        :key="`friend_${index}`"
      >
        <div class="flex-2">
          <div class="w-12 h-12 relative">
            <img
              class="w-12 h-12 rounded-full mx-auto"
              src="/twc/resources/profile-image.png"
              alt="chat-user"
            />
            <span
              :class="`
                absolute
                w-4
                h-4
                ${friend.online ? 'bg-green-400' : 'bg-gray-400'}
                rounded-full
                right-0
                bottom-0
                border-2 border-white
              `"
            ></span>
          </div>
        </div>
        <div class="flex-1 px-2">
          <div class="truncate w-32">
            <span class="text-gray-800">{{ friend.name }}</span>
          </div>
          <div>
            <small class="text-gray-600">{{ friend.lastMsg }}</small>
          </div>
        </div>
        <div class="flex-2 text-right">
          <div>
            <small class="text-gray-500">{{ friend.lastMsgTime }}</small>
          </div>
          <div>
            <small
              class="
                text-xs
                bg-red-500
                text-white
                rounded-full
                h-6
                w-6
                leading-6
                text-center
                inline-block
              "
            >
              {{ friend.unReadCount }}
            </small>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      friends: [
        {
          name: '良宏',
          img: '/twc/resources/profile-image.png',
          lastMsg: '哇！厉害！',
          lastMsgTime: '15 April',
          unReadCount: 3,
          online: true
        },
        {
          name: '平丽',
          img: '/twc/resources/profile-image.png',
          lastMsg: '呀！真不错！',
          lastMsgTime: '15 April',
          unReadCount: 5,
          online: true
        },
        {
          name: '亚琴',
          img: '/twc/resources/profile-image.png',
          lastMsg: '哈哈！真好！',
          lastMsgTime: '15 April',
          unReadCount: 2,
          online: false
        }
      ]
    }
  }
}
</script>
