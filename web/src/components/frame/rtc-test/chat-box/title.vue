<template>
  <div class="flex-3">
    <div class="text-xl py-1 mb-8 border-b-2 border-gray-200">
      <input type="text" v-model="modelValue" @blur="$emit('update:modelValue', this.modelValue)" />
    </div>
  </div>
</template>
<script>
export default {
  props: ['modelValue'],
  emits: ['update:modelValue']
}
</script>
