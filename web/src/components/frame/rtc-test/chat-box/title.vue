<template>
  <div class="flex-3">
    <div class="text-xl py-1 mb-8 border-b-2 border-gray-200 flex justify-between">
      <input type="text" v-model="modelValue" @blur="$emit('update:modelValue', this.modelValue)" />
      <label>{{ selfName }}</label>
    </div>
  </div>
</template>
<script>
export default {
  props: ['modelValue', 'selfName'],
  emits: ['update:modelValue']
}
</script>
