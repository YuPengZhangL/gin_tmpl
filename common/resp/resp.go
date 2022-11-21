package resp

type BaseResp struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// NewBaseResp 返回简单的Response
func NewBaseResp(code int64, msg string) *BaseResp {
	return &BaseResp{
		StatusCode: code,
		StatusMsg:  msg,
	}
}
