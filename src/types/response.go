package types

type Response[T any] struct {
	Status bool
	Code int
	Message string
	Content *T
	ContentArray []T
}

func NewResponse[T any](
	status bool, 
	code int, 
	msg string, 
	content T,
) *Response[T] {
	return &Response[T]{
		Status: status,
		Code: code,
		Message: msg,
		Content: &content,
		ContentArray: nil,
	}
}

func NewManyResponse[T any](
	status bool,
	code int,
	msg string,
	arr []T,
) *Response[T] {
	return &Response[T]{
		Status: status,
		Code: code,
		Message: msg,
		Content: nil,
		ContentArray: arr,
	}
}