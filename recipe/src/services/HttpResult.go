package services

type HttpResult struct {
	Err  error  // エラーデータ
	Msg  string // メッセージ
	Code int    // HTTP ステータスコード
}

func (result *HttpResult) Error() string {
	return result.Err.Error()
}

func NewResult(err error,msg string,code int) HttpResult {
	return HttpResult{
		Err:  err,
		Msg:  msg,
		Code: code,
	}
}