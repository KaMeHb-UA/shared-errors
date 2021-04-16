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
