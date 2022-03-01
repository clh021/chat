<template>
  <div>
    <InputBoxWrite v-model="showBtnPanel" @onSendMsg="sendMsgHandle" />
    <InputBoxBtn v-show="showBtnPanel" />
  </div>
</template>
<script>
import InputBoxWrite from './input-box-write.vue'
import InputBoxBtn from './input-box-btn.vue'
export default {
  components: {
    InputBoxWrite,
    InputBoxBtn
  },
  data() {
    return { showBtnPanel: false }
  },
  methods: {
    sendMsgHandle(msgTxt) {
      console.log('sendMsg', msgTxt)
    },
    moreBtnHandle() {
      console.log('moreBtn')
    }
  }
}
</script>
