<template>
  <!-- https://daisyui.com/components/alert/ -->
  <!-- 注意有以下扩展方式可以使用 -->
  <!-- Alert with buttons -->
  <!-- Alert with title and description -->
  <div :class="`alert${color ? ' alert-' + color : ''} shadow-lg`">
    <div>
      <svg
        v-show="color == 'info'"
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        class="stroke-current flex-shrink-0 w-6 h-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
        ></path>
      </svg>

      <svg
        v-show="color == 'success'"
        xmlns="http://www.w3.org/2000/svg"
        class="stroke-current flex-shrink-0 h-6 w-6"
        fill="none"
        viewBox="0 0 24 24"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
        />
      </svg>

      <svg
        v-show="color == 'warning'"
        xmlns="http://www.w3.org/2000/svg"
        class="stroke-current flex-shrink-0 h-6 w-6"
        fill="none"
        viewBox="0 0 24 24"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
        />
      </svg>

      <svg
        v-show="color == 'error'"
        xmlns="http://www.w3.org/2000/svg"
        class="stroke-current flex-shrink-0 h-6 w-6"
        fill="none"
        viewBox="0 0 24 24"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
        />
      </svg>

      <span>{{ text }}</span>
    </div>
  </div>
</template>

<script>
export default {
  props: [
    'color', // color: warning, success, error, info ,''
    // alert-warning, alert-success, alert-error, alert-info // 有这个注释可以避免 css 压缩程序以为没有使用样式而去除"冗余"
    'text'
  ]
}
</script>

<style>
</style>
