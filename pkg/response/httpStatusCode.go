package response

const (
	CodeSuccess      = 20001
	CodeParamInvalid = 20003
	ErrInvalidToken  = 30001
)

// message
var Msg = map[int]string{
	CodeSuccess:      "success",
	CodeParamInvalid: "param invalid",
	ErrInvalidToken:  "token is invalid",
}
