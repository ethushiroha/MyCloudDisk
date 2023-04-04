package test

import (
	"fmt"
	"testing"
)

func TestNilChannel(t *testing.T) {
	ch := make(chan *int)

	go func() {
		ch <- nil
	}()

	d := <-ch
	fmt.Println(d)
}
