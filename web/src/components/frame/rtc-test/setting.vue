<template>
  <div>
    <!-- <label for="my-modal-4" class="btn modal-button">open modal</label> -->
    <input
      type="checkbox"
      id="my-modal-4"
      class="modal-toggle"
      v-model="modelValue"
      @blur="$emit('update:modelValue', this.modelValue)"
    />
    <label for="my-modal-4" class="modal cursor-pointer">
      <label class="modal-box relative" for="">
        <div class="form-control w-full max-w-xs">
          <h1>给自己取一个名字</h1>
          <label class="label">
            <!-- <span class="label-text">给自己取一个名字，方便朋友们找到你</span> -->
            <!-- <span class="label-text-alt">Alt label</span> -->
          </label>
          <input
            type="text"
            placeholder="Type here"
            class="input input-bordered w-full max-w-xs"
            v-model="selfName"
            @blur="$emit('newName', this.selfName)"
          />
          <label class="label">
            <span class="label-text-alt">朋友们将需要输入这个名字来找到你</span>
            <!-- <span class="label-text-alt">Alt label</span> -->
          </label>
        </div>
      </label>
    </label>
  </div>
</template>
<script>
export default {
  props: ['modelValue'],
  emits: ['update:modelValue'],
  data() {
    return {
      selfName: ''
    }
  }
}
</script>
