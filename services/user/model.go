package user

import (
	"time"
)

type User struct {
	Id      string    // 编号
	Name    string    // 真实姓名
	Mobile  string    // 手机号码
	Role    string    // 角色
	Rcode   string    // 邀请码
	State   bool      // 在线状态
	LoginAt time.Time // 登陆时间
	ReMark  string    // 备注
}

type Rule struct {
	Id     int
	Role   string // 角色名
	Access string // 权限列表
}
