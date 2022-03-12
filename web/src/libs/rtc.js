export default class rtc {
  /**
   * isRtcSupport
   * @returns bool rtcSupport
   */
  static isSupport () {
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

  /**
   * Rtc Init
   * @param {string} name
   * @param {function} openHandle function(conn){}
   * @param {function} connHandle function(id){}
   */
  rtcInit (name, openHandle, connHandle) {
    // peerserver的连接选项(debug:3表示打开调试，将在浏览器的console输出详细日志)
    let connOption = { host: '192.168.31.127', port: 9000, path: '/', debug: 3 };

    // 创建peer实例
    peer = new Peer(hashCode(name), connOption);

    // register成功的回调
    peer.on('open', openHandle);

    peer.on('connection', connHandle);
    // let connHandle = (conn) => {
    //   conn.on('data', (data) => {
    //     console.log(JSON.parse(data));
    //   });
    // };
  }

}

