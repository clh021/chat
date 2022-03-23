<template>
  <div class="flex-3">
    <div class="text-xl py-1 mb-8 border-b-2 border-gray-200 flex justify-between">
      <!-- <input
        ref="friendName"
        type="text"
        v-model="modelValue"
        @blur="$emit('update:modelValue', this.modelValue)"
      /> -->
      <label>{{ modelValue }}</label>
      <label>{{ selfName }}</label>
    </div>
  </div>
</template>
<script>
export default {
  props: ['modelValue', 'selfName'],
  emits: ['update:modelValue'],
  methods: {
    friendFocus() {
      this.$refs.friendName.focus()
    }
  }
}
</script>
