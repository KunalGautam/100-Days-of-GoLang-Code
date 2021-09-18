// Example copied from https://www.digitalocean.com/community/tutorials/creating-custom-errors-in-go
package main

import (
	"errors"
	"fmt"
)

type WrappedError struct {
	Context string
	Err     error
}

func (w *WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", w.Context, w.Err)
}

func Wrap(err error, info string) *WrappedError {
	return &WrappedError{
		Context: info,
		Err:     err,
	}
}

func main() {
	err := errors.New("boom!")
	err = Wrap(err, "main")

	fmt.Println(err)
}
