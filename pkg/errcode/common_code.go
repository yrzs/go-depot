package errcode

var (
	Success                   = NewError(200, "成功")
	ServerError               = NewError(500, "服务内部错误")
	InvalidParams             = NewError(400, "入参错误")
	NotFound                  = NewError(404, "找不到")
	UnauthorizedAuthNotExist  = NewError(422, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(423, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewError(424, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewError(425, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewError(429, "请求过多")
)
