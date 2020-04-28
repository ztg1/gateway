package msg

type messageVo struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

//错误返回
func NewMessageVo(code int64, error error) *messageVo {
	return &messageVo{
		Code:    code,
		Message: error.Error(),
	}
}
