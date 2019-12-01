package stream

type Exception interface {
	// error type
	Type() string
	// return message(errType + detailMessage)
	Error() string
}

type ExceptionImpl struct {
	errType       string
	detailMessage string
}

func NewException(errType string, detailMessage string) Exception {
	return &ExceptionImpl{
		errType:       errType,
		detailMessage: detailMessage,
	}
}

func (this *ExceptionImpl) Type() string {
	return this.errType
}

func (this *ExceptionImpl) Error() string {
	if this.detailMessage == "" {
		return this.errType
	} else {
		return this.errType + ": " + this.detailMessage
	}
}

func ZeroSliceException(detailMessage string) Exception {
	return NewException("ZeroSliceException", detailMessage)
}

// IllegalStateException
func IllegalStateException(detailMessage string) Exception {
	return NewException("IllegalStateException", detailMessage)
}

// NullPointerException
func NullPointerException(detailMessage string) Exception {
	return NewException("NullPointerException", detailMessage)
}

// UnsupportedOperationException
func UnsupportedOperationException(detailMessage string) Exception {
	return NewException("UnsupportedOperationException", detailMessage)
}

// ClassCastException
func ClassCastException(detailMessage string) Exception {
	return NewException("ClassCastException", detailMessage)
}
