package main

import (
	"fmt"
)

type Handler interface {
	Filter(err error, r interface{}) error
}

type Logger interface {
	Ef(format string, a ...interface{})
}

// Handle panic by hdr, which filter the error.
// Finally log err with logger.
func HandlePanic(hdr Handler, logger Logger) error {
	return handlePanic(recover(), hdr, logger)
}

type hdrFunc func(err error, r interface{}) error

func (v hdrFunc) Filter(err error, r interface{}) error {
	return v(err, r)
}

type loggerFunc func(format string, a ...interface{})

func (v loggerFunc) Ef(format string, a ...interface{}) {
	v(format, a...)
}

// Handle panic by hdr, which filter the error.
// Finally log err with logger.
func HandlePanicFunc(hdr func(err error, r interface{}) error,
	logger func(format string, a ...interface{}),
) error {
	var f Handler
	if hdr != nil {
		f = hdrFunc(hdr)
	}

	var l Logger
	if logger != nil {
		l = loggerFunc(logger)
	}

	return handlePanic(recover(), f, l)
}

func handlePanic(r interface{}, hdr Handler, logger Logger) error {
	if r != nil {
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("r is %v", r)
		}

		if hdr != nil {
			err = hdr.Filter(err, r)
		}

		if err != nil && logger != nil {
			logger.Ef("panic err %+v", err)
		}

		return err
	}

	return nil
}

func main() {
	func() {
		defer HandlePanicFunc(nil, func(format string, a ...interface{}) {
			fmt.Println(fmt.Sprintf(format, a...))
		})

		panic("ok")
	}()

	logger := func(format string, a ...interface{}) {
		fmt.Println(fmt.Sprintf(format, a...))
	}
	func() {
		defer HandlePanicFunc(nil, logger)

		panic("ok")
	}()
}
