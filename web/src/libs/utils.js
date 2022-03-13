//peer连接时，id不允许有中文，所以转换成 hashCode 数字
function hashCode (str) {
  var hash = 5381, i = str.length;
  while (i) {
    hash = (hash * 33) ^ str.charCodeAt(--i);
  }
  return hash >>> 0;
}

function testFunc (params) {
  return params;
}
export { hashCode };
