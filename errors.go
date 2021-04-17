package sharederrs

import "github.com/go-stack/stack"

//APIErrorType - short version of internal.APIError
type APIErrorType = APIError

// APIErrorData describes an additional info about an error
type APIErrorData struct {
	Name     string `json:"name"`
	CodeName string `json:"codeName,omitempty"`
	Stack    string `json:"stack,omitempty"`
}

// APIError describes the error that occurred when calling the API methods with the error name, code, execution stack and message
type APIError struct {
	Message string        `json:"message"`
	Code    int           `json:"code"`
	Data    *APIErrorData `json:"data"`
}

//SetStack - add error stack with trace
func (err *APIError) SetStack(stack string) *APIError {
	if err.Data == nil {
		err.Data = &APIErrorData{
			Stack: stack,
		}
	} else {
		err.Data.Stack = stack
	}
	return err
}

//SetTrace - add error stack with trace
func (err *APIError) SetTrace() *APIError {
	return err.SetStack(stack.Trace().TrimRuntime().String())
}

//SetMessage - assign a message to error
func (err *APIError) SetMessage(message string) *APIError {
	err.Message = message
	return err
}

//M - short version of SetMessage
func (err *APIError) M(message string) *APIError {
	return err.SetMessage(message)
}

//SetName - sets the error name
func (err *APIError) SetName(name string) *APIError {
	if err.Data == nil {
		err.Data = &APIErrorData{
			Name: name,
		}
	} else {
		err.Data.Name = name
	}
	return err
}

//N - short version of SetName
func (err *APIError) N(name string) *APIError {
	return err.SetName(name)
}

//GetName - get error name
func (err *APIError) GetName() string {
	if err.Data == nil {
		return ""
	}
	return err.Data.Name
}

//GetStack - get error stack
func (err *APIError) GetStack() string {
	if err.Data == nil {
		return ""
	}
	return err.Data.Stack
}

//GetCodeName - get error code name ^ↀᴥↀ^
func (err *APIError) GetCodeName() string {
	if err.Data == nil {
		return ""
	}
	return err.Data.CodeName
}

// APIErrorPredefined creates new APIError struct.
// 4'th parameter (stack) is an optional one
func APIErrorPredefined(code int, name string, msg string, codeName ...string) *APIError {
	res := &APIError{
		Message: msg,
		Code:    code,
	}
	res.SetName(name)
	//stack ...string
	/*if len(stack) > 0 {
		res.Data.Stack = stack[0]
	}*/
	if len(codeName) > 0 {
		res.Data.CodeName = codeName[0]
	}
	return res
}

//JSON errors

//MethodNameInvalid ..
var MethodNameInvalid = APIErrorPredefined(-3200, "Method name invalid", "Method name invalid")

//ParseErr ..
var ParseErr = APIErrorPredefined(-32700, "Parse error", "Invalid JSON was received by the server")

//InvalidRequestErr ..
var InvalidRequestErr = APIErrorPredefined(-32600, "Invalid Request", "The JSON sent is not a valid Request object")

//MethodNotFoundErr ..
var MethodNotFoundErr = APIErrorPredefined(-32601, "Method not found", "The method does not exist / is not available")

//InvalidArgsErr ..
var InvalidArgsErr = APIErrorPredefined(-32602, "Invalid params", "Invalid method parameter(s)")

//data errors
var dataInvalidErr = APIErrorPredefined(-35100, "Data handle error", "DATA_HANDLE_ERROR")
var dataEmptyErr = APIErrorPredefined(-35200, "Data empty", "", "DATA_EMPTY")
var dataUnknownErr = APIErrorPredefined(-35300, "Data unknown", "", "DATA_UNKNOWN")
var dataNotFoundErr = APIErrorPredefined(-35010, "Data not found", "", "DATA_NOT_FOUND")
var dataRequestErr = APIErrorPredefined(-35020, "Data request error", "", "DATA_REQ_ERR")
var dataExistsErr = APIErrorPredefined(-35030, "Data exists", "", "DATA_EXISTS")
var dataNoAccessErr = APIErrorPredefined(-35040, "Data no acesss", "", "DATA_NO_ACCESS")
var dataHandleErr = APIErrorPredefined(-35001, "Data handle error", "", "DATA_HANDLE_ERR")

//user errors
var userIsAuthErr = APIErrorPredefined(-36100, "User is authorized", "User is authorized", "USER_IS_AUTH")
var userNotAuthErr = APIErrorPredefined(-36200, "User not authorized", "User not authorized", "USER_NOT_AUTH")
var userNotFoundErr = APIErrorPredefined(-36010, "User not found", "", "USER_NOT_FOUND")
var userDeletedErr = APIErrorPredefined(-36001, "User deleted", "", "USER_DELETED")
var userStatusInvalidErr = APIErrorPredefined(-36020, "User status invalid", "", "USER_STATUS_INVALID")

//bot errors
var botNotFoundErr = APIErrorPredefined(-34100, "Bot not found", "", "BOT_NOT_FOUND")
var botIsActiveErr = APIErrorPredefined(-34200, "Bot is active", "", "BOT_IS_ACTIVE")
var botInactiveErr = APIErrorPredefined(-34300, "Bot is inactive", "", "BOT_INACTIVE")
var botBalanceNotEnoughErr = APIErrorPredefined(-34010, "Not enough balance", "", "BALANCE_NOT_ENOUGH")
var botOrderInvalidErr = APIErrorPredefined(-34020, "Invalid order", "", "BOT_ORDER_INVALID")

//service errors
var serviceReqFailedErr = APIErrorPredefined(-37010, "Service request failed", "", "SERVICE_REQ_FAILED")
var serviceDisconnectedErr = APIErrorPredefined(-37100, "Service disconnected", "", "SERVICE_DISCONNECTED")
var serviceNoAccess = APIErrorPredefined(-37010, "No access", "", "SERVICE_NO_ACCESS")

//notification errors
var telegramUserAuthErr = APIErrorPredefined(-38000, "Telegram user authentication error", "", "TELEGRAM_USER_AUTH_ERROR")
var notificationsConfigErr = APIErrorPredefined(-38010, "Error configuring notifications", "", "NOTIFICATIONS_CONFIG_ERROR")
var saveNotificationEmailErr = APIErrorPredefined(-38020, "Error saving email for notifications", "", "NOTIF_SAVE_EMAIL_ERROR")
