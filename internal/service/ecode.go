package service

import (
	"github.com/go-kratos/kratos/pkg/ecode"
)

func init() {
	codes := make(map[int]string)
	codes[0] = "获取成功"
	codes[-200] = "获取成功"
	codes[-202] = "获取失败"
	codes[-304] = "木有改动"
	codes[-307] = "撞车跳转"
	codes[-400] = "请求错误"
	codes[-401] = "未认证"
	codes[-403] = "访问权限不足"
	codes[-404] = "啥都木有"
	codes[-405] = "不支持该方法"
	codes[-409] = "冲突"
	codes[-498] = "客户端取消请求"
	codes[-500] = "服务器错误"
	codes[-503] = "过载保护,服务暂不可用"
	codes[-504] = "服务调用超时"
	codes[-509] = "超出限制"
	ecode.Register(codes)
}
