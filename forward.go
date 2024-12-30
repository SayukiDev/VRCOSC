package VRCOSC

import (
	"github.com/hypebeast/go-osc/osc"
	"log"
)

type ForwardDispatcher struct {
	logger *log.Logger
	c      *osc.Client
	d      osc.Dispatcher
}

func newForwardDispatcher(l *log.Logger, c *osc.Client, d osc.Dispatcher) *ForwardDispatcher {
	return &ForwardDispatcher{
		logger: l,
		c:      c,
		d:      d,
	}
}

func (f *ForwardDispatcher) Dispatch(p osc.Packet) {
	err := f.c.Send(p)
	if err != nil {
		f.logger.Println("Warn:", "[", "Forward message packet error:", err, "]")
	}
	if f.d != nil {
		f.d.Dispatch(p)
	}
}

func (v *VRCOsc) forwardHandler(msg *osc.Message) {
	err := v.senderF.Send(msg)
	if err != nil {
		v.logger.Println("Warn:", "[", "Forward message packet error:", err, "]")
	}
}
