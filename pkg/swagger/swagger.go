package swagger

//
// HttpOk 为 Swagger 文档定义的结构体, 当接口调用成功之后, 返回该类型的结构体
//
type HttpOk struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"msg" example:"接口调用成功"`
	//Data interface{} `json:"data" example:"data"`
}

//
// Http400 为 Swagger 文档定义的结构体, 当接口调用发送非法请求的时候, 返回该类型的结构体
//
type Http400 struct {
	Code int         `json:"code" example:"400"`
	Msg  string      `json:"msg" example:"Bad Request"`
}

//
// Http404 为 Swagger 文档定义的结构体, 当接口调用发生页面不存在的错误时, 返回该类型的结构体
//
type Http404 struct {
	Code int         `json:"code" example:"404"`
	Msg  string      `json:"msg" example:"Page Not Found"`
}

//
// Http500 为 Swagger 文档定义的结构体, 当调用接口发生服务器内部错误时, 返回该类型的结构体
//
type Http500 struct {
	Code int         `json:"code" example:"500"`
	Msg  string      `json:"msg" example:"服务器内部错误"`
}
