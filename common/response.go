package common

type (
	//HTTP请求正确返回结构，带data
	successResponse struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}

	//HTTP请求错误返回结构
	errorResponse struct {
		Success bool `json:"success"`
		Error   struct {
			Code    uint16 `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}
)

//组织正确时数据结构,带data
func Success(data interface{}) *successResponse {
	response := &successResponse{}
	response.Success = true
	response.Data = data
	return response
}

//组织错误时数据结构
func Error(code uint16, message string) *errorResponse {
	response := &errorResponse{}
	response.Success = false
	response.Error.Code = code
	response.Error.Message = message
	return response
}
