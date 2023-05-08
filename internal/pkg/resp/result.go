package resp

type CommonResult struct {
	Code int         `json:"code"` // 自定义消息码
	Msg  string      `json:"msg"`  // 消息提示
	Data interface{} `json:"data"` // 附加数据: 比如查询操作查询到的结果集就可以附加到这个数据字段上
}

// NewCommonResult
//	@parameter	code 错误码
//	@parameter	msg 错误码描述
//	@parameter	data 附加数据
//	@return		*CommonResult
func NewCommonResult(code int, msg string, data interface{}) *CommonResult {
	return &CommonResult{Code: code, Msg: msg, Data: data}
}