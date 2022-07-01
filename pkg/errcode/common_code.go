package errcode

var (
	Success                  = NewError(200, "成功")
	ServerError              = NewError(500, "服务内部错误")
	InvalidParams            = NewError(400, "入参错误")
	NotFound                 = NewError(404, "找不到")
	UnauthorizedAuthNotExist = NewError(422, "鉴权失败，找不到对应的AccessToken")
	UnauthorizedTokenError   = NewError(423, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout = NewError(424, "鉴权失败，Token 超时")
	SignError                = NewError(425, "验签失败，签名错误")
	SignTimeOut              = NewError(426, "验签失败，签名过期")
	TooManyRequests          = NewError(429, "请求过多")
)
