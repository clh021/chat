//peer连接时，id不允许有中文，所以转换成 字母数字
function nameEncode (str) {
  return encodeURIComponent(str).replaceAll("%", "_").substr(1);
}
function nameDecode (str) {
  return decodeURIComponent(str);
}

function testFunc (params) {
  return params;
}
export { nameEncode, nameDecode };
