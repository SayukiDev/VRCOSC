package VRCOSC

import (
	"github.com/hypebeast/go-osc/osc"
	"testing"
)

func TestVRCOsc_AddParamsHandler(t *testing.T) {
	c := make(chan struct{})
	t.Log(v.AddParamsHandler("test", func(msg *osc.Message) {
		t.Log(msg.Arguments)
		c <- struct{}{}
	}))
	go v.Run()
	select {
	case <-c:
	}
}
