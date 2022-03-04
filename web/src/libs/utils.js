export default {
  install: (app, options) => {
    app.config.globalProperties.rtcSupport = () => {
      navigator.getUserMedia =
        navigator.getUserMedia ||
        navigator.webkitGetUserMedia ||
        navigator.mozGetUserMedia ||
        navigator.msGetUserMedia
      if (navigator.getUserMedia) {
        return true
      }
      return false
    }
  }
}
