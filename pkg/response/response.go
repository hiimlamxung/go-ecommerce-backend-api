package response

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`    // status code
	Message string      `json:"message"` // thông báo lỗi
	Data    interface{} `json:"data"`    // dữ liệu return
}

// ReturnResponse - function cơ bản nhất, yêu cầu đầy đủ tất cả tham số
// Params:
//   - c: gin context
//   - httpCode: HTTP status code (200, 404, 500...)
//   - message: thông báo trả về cho client
//   - internalCode: mã lỗi nội bộ của hệ thống
//   - data: dữ liệu response
func ReturnResponse(c *gin.Context, httpCode int, message string, internalCode int, data interface{}) {
	c.JSON(httpCode, ResponseData{
		Code:    internalCode,
		Message: message,
		Data:    data,
	})
}

// ResponseWithMessage - chỉ cần truyền message. internalCode sẽ tự động = httpCode
// Params:
//   - c: gin context
//   - httpCode: HTTP status code, sẽ được dùng làm internalCode
//   - message: thông báo custom
//   - data: dữ liệu response
//
// Example:
//
//	ResponseWithMessage(c, 200, "Thành công", userData)
//	ResponseWithMessage(c, 400, "Dữ liệu không hợp lệ", nil)
func ResponseWithMessage(c *gin.Context, httpCode int, message string, data interface{}) {
	ReturnResponse(c, httpCode, message, httpCode, data)
}

// ResponseWithInternalCode - chỉ cần truyền internalCode, message sẽ lấy từ Msg map
// Params:
//   - c: gin context
//   - httpCode: HTTP status code
//   - internalCode: mã lỗi nội bộ, message sẽ được tự động lấy từ Msg map
//   - data: dữ liệu response
//
// Example:
//
//	ResponseWithInternalCode(c, 200, CodeSuccess, userData)       // message = "success"
//	ResponseWithInternalCode(c, 400, CodeParamInvalid, nil)       // message = "param invalid"
func ResponseWithInternalCode(c *gin.Context, httpCode int, internalCode int, data interface{}) {
	message, exists := Msg[internalCode]
	if !exists {
		// Nếu không tìm thấy message cho internalCode, thử dùng httpCode
		if fallbackMsg, existsFallback := Msg[httpCode]; existsFallback {
			message = fallbackMsg
		} else {
			message = ""
		}
	}
	ReturnResponse(c, httpCode, message, internalCode, data)
}

// ResponseWithInternalCodeAndMessage - function đầy đủ với cả internalCode và message custom
// Params:
//   - c: gin context
//   - httpCode: HTTP status code
//   - internalCode: mã lỗi nội bộ
//   - message: thông báo custom (ghi đè message từ Msg map)
//   - data: dữ liệu response
//
// Example:
//
//	ResponseWithInternalCodeAndMessage(c, 200, CodeSuccess, "Tạo user thành công!", userData)
func ResponseWithInternalCodeAndMessage(c *gin.Context, httpCode int, internalCode int, message string, data interface{}) {
	ReturnResponse(c, httpCode, message, internalCode, data)
}
