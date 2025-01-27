package entities

import ()



type AppErr struct{
    Code int
    Message string
}
func (e *AppErr) Error() string{
    return e.Message
}

func NewAppErr (code int, message string) *AppErr{
    return &AppErr{Code: code, Message: message}
}
func NewNotFoundError (message string) *AppErr{
    return NewAppErr(404, message)
}
func NewInternalServerError (message string) *AppErr {
    return NewAppErr(500, message)
}
func NewBadRequestError (message string) *AppErr{
    return NewAppErr(400, message)
}
func NewUnauthorizedError (message string) *AppErr{
    return NewAppErr(401, message)
}