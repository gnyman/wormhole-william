// Package mobile provides an interface compatible with
// gomobile bindings. It is not indented to be used
// on its own and its interface is not stable.
package mobile

import (
	"context"
	"sync/atomic"

	"github.com/psanford/wormhole-william/wormhole"
)

type SendResult struct {
	code       string
	resultChan chan wormhole.SendResult

	gotResult int32
	ok        bool
	error     error
}

func (s *SendResult) Code() string {
	return s.code
}

func (s *SendResult) Ready() bool {
	return atomic.LoadInt32(&s.gotResult) != 0
}

func (s *SendResult) Result() (ok bool, err error) {
	return s.ok, s.error
}

func (s *SendResult) wait() {
	result := <-s.resultChan
	s.ok = result.OK
	s.error = result.Error
	atomic.StoreInt32(&s.gotResult, 1)
}

func SendText(msg string) (*SendResult, error) {
	var c wormhole.Client
	code, result, err := c.SendText(context.Background(), msg)
	if err != nil {
		return nil, err
	}

	sr := &SendResult{
		code:       code,
		resultChan: result,
	}

	return sr, nil
}

func main() {
}
