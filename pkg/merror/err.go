package merror

import (
	"errors"
	"fmt"
)

type Merror struct {
	Code int64
	Msg  string
}

 func (err Merror) Error() string {
	return fmt.Sprintf("[%v]: %s", err.Code, err.Msg)
}

func NewMerror (code int64, msg string) *Merror {
	return &Merror{
		Code: code,
		Msg:  msg,
	}
}

func CoverError(err error) Merror {
	if err == nil {
		return Merror{
			Code: 200,
			Msg:  "success",
		}
	}
	var merrorPtr *Merror
	if errors.As(err, &merrorPtr) && merrorPtr != nil {
		return *merrorPtr
	}
	var merror Merror
	if errors.As(err, &merror) {
		return merror
	}
	return Merror{
		Code: 50000,
		Msg:  err.Error(),
	}
}