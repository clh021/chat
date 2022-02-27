<template>
  <div class="write bg-white shadow flex rounded-lg">
    <!-- <div class="flex-3 flex content-center items-center text-center p-4 pr-0">
      <span class="block text-center text-gray-400 hover:text-gray-800">
        <svg
          fill="none"
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          stroke="currentColor"
          viewBox="0 0 24 24"
          class="h-6 w-6"
        >
          <path
            d="M14.828 14.828a4 4 0 01-5.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
          ></path>
        </svg>
      </span>
    </div> -->
    <div class="flex-1">
      <textarea
        name="message"
        class="w-full block outline-none py-4 px-4 bg-transparent"
        rows="1"
        placeholder="输入文字消息..."
        autofocus
      ></textarea>
    </div>
    <div class="flex-2 w-16 p-2 flex content-center items-center">
      <!-- <div class="flex-1 text-center">
        <span class="text-gray-400 hover:text-gray-800">
          <span class="inline-block align-text-bottom">
            <svg
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              stroke="currentColor"
              viewBox="0 0 24 24"
              class="w-6 h-6"
            >
              <path
                d="M15.172 7l-6.586 6.586a2 2 0 102.828 2.828l6.414-6.586a4 4 0 00-5.656-5.656l-6.415 6.585a6 6 0 108.486 8.486L20.5 13"
              ></path>
            </svg>
          </span>
        </span>
      </div> -->
      <div class="flex-1">
        <label class="swap swap-rotate">
          <!-- this hidden checkbox controls the state -->
          <input type="checkbox" />

          <!-- 发送 icon -->
          <svg
            class="swap-on fill-current w-10 h-10"
            version="1.0"
            xmlns="http://www.w3.org/2000/svg"
            width="50"
            height="50"
            viewBox="0 0 50 50"
            preserveAspectRatio="xMidYMid meet"
          >
            <g transform="translate(0,50) scale(0.100000,-0.100000)" fill="#000000" stroke="none">
              <path
                d="M155 456 c-60 -28 -87 -56 -114 -116 -36 -79 -19 -183 42 -249 33
-36 115 -71 167 -71 52 0 134 35 167 71 34 37 63 110 63 159 0 52 -35 134 -71
167 -37 34 -110 63 -159 63 -27 0 -65 -10 -95 -24z m180 -15 c128 -58 164
-223 72 -328 -101 -115 -283 -88 -348 52 -79 171 104 354 276 276z"
              />
              <path
                d="M190 335 c-29 -30 -51 -58 -47 -62 4 -4 28 13 52 37 l45 44 0 -122
c0 -75 4 -122 10 -122 6 0 10 47 10 122 l0 122 45 -44 c24 -24 48 -41 52 -37
6 7 -94 117 -107 117 -3 0 -30 -25 -60 -55z"
              />
            </g>
          </svg>

          <!-- 加号 icon -->
          <svg
            class="swap-off fill-current w-10 h-10"
            xmlns="http://www.w3.org/2000/svg"
            x="0px"
            y="0px"
            width="50"
            height="50"
            viewBox="0 0 50 50"
            style="fill: #000000"
          >
            <path
              d="M 25 2 C 12.309295 2 2 12.309295 2 25 C 2 37.690705 12.309295 48 25 48 C 37.690705 48 48 37.690705 48 25 C 48 12.309295 37.690705 2 25 2 z M 25 4 C 36.609824 4 46 13.390176 46 25 C 46 36.609824 36.609824 46 25 46 C 13.390176 46 4 36.609824 4 25 C 4 13.390176 13.390176 4 25 4 z M 24 13 L 24 24 L 13 24 L 13 26 L 24 26 L 24 37 L 26 37 L 26 26 L 37 26 L 37 24 L 26 24 L 26 13 L 24 13 z"
            ></path>
          </svg>
        </label>
        <!-- <button class="bg-blue-400 w-10 h-10 rounded-full inline-block">
          <span class="inline-block align-text-bottom">
            <svg
              fill="none"
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              viewBox="0 0 24 24"
              class="w-4 h-4 text-white"
            >
              <path d="M5 13l4 4L19 7"></path>
            </svg>
          </span>
        </button> -->
      </div>
    </div>
  </div>
</template>
<script lang="ts">
export default {}
</script>
