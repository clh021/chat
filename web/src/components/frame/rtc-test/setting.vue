<template>
  <div>
    <div class="form-control w-full">
      <h1 class="text-center text-2xl mb-2">起一个名字</h1>
      <label class="label">
        <span class="label-text">您的朋友们输入您的名字来找您</span>
        <!-- <span class="label-text-alt"></span> -->
      </label>
      <div class="flex">
        <input
          type="text"
          placeholder="这里输入"
          class="input input-bordered w-full"
          v-model="modelValue"
          @change="$emit('update:modelValue', modelValue)"
        />
      </div>
      <!-- <label class="label"> -->
      <!-- <span class="label-text-alt"></span> -->
      <!-- <span class="label-text-alt"></span> -->
      <!-- </label> -->
    </div>
  </div>
</template>
<script>
export default {
  props: ['modelValue'],
  emits: ['update:modelValue']
}
</script>
