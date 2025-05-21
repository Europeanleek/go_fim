package response

// Response，统一返回格式
import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response 是一个HTTP响应函数，它根据提供的错误和响应数据来生成HTTP响应。
// 如果提供的错误为nil，则返回一个成功的JSON响应，其中包含状态码0、成功消息和数据。
// 如果提供的错误不为nil，则返回一个包含错误码和错误消息的JSON响应，并返回HTTP状态码400。
func Response(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		r := &Body{
			Code: 0,
			Msg:  "成功",
			Data: resp,
		}
		httpx.WriteJson(w, http.StatusOK, r)
		return
	}
	errCode := uint32(10086)
	errMsg := err.Error()
	httpx.WriteJson(w, http.StatusBadRequest, &Body{
		Code: errCode,
		Msg:  errMsg,
		Data: nil,
	})
}
