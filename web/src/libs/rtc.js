import Peer from 'peerjs';
import { hashCode } from './utils';
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

  constructor(name) {
    // docker run -d -p 47201:9000 peerjs/peerjs-server
    // peerserver的连接选项(debug:3表示打开调试，将在浏览器的console输出详细日志)
    let port = '.', two = '2', seven = '7', nine = '9';
    this.connOption = { host: `1${two}0${port}${two}4${port}${nine}${nine}${port}1${seven}${seven}`, port: `4${seven}${two}01`, path: '/myapp', debug: 3 };
    // 创建peer实例 // hashCode(name)
    this.peer = new Peer(hashCode(name), this.connOption);
    return this;
  }

  /**
   * 事件绑定
   * @param {Object} handles peer.on(eventName,{eventName: <function>})
   */
  handles (handles) {
    let events = ['open', 'connection', 'call', 'close', 'disconnected', 'error'];
    events.forEach(eventName => {
      if (handles.hasOwnProperty(eventName)) {
        this.peer.on(eventName, handles[eventName])
      }
    });
    // let handles = {
    //   // register成功的回调
    //   open: (id) => {
    //     console.log(JSON.parse(id))
    //   },
    //   connection: (dataConnection) => { },
    //   call: (mediaConnection) => { },
    //   close: () => { },
    //   disconnected: () => { },
    //   error: (err) => {
    //     console.log(err.type)
    //   }
    // };
  }
}

