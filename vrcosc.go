package VRCOSC

import (
	"context"
	"fmt"
	"github.com/hypebeast/go-osc/osc"
	"github.com/sourcegraph/conc/pool"
	"log"
	"os"
)

type VRCOsc struct {
	sender        *osc.Client
	recver        *osc.Server
	EnableForward bool
	senderF       *osc.Client
	recverF       *osc.Server
	dispatcher    *osc.StandardDispatcher
	logger        *log.Logger
}

type Options struct {
	Logger            *log.Logger
	Host              string
	SendPort          int
	RecvPort          int
	EnableForward     bool
	ForwardListenHost string
	ForwardSendPort   int
	ForwardRecvPort   int
}

func New(o *Options) *VRCOsc {
	if o.Logger == nil {
		o.Logger = log.New(os.Stdout, "VRCOSC", 0)
	}

	c := osc.NewClient(o.Host, o.SendPort)
	d := osc.NewStandardDispatcher()
	s := &osc.Server{
		Addr:       fmt.Sprintf("%s:%d", o.Host, o.RecvPort),
		Dispatcher: d,
	}

	v := &VRCOsc{
		sender:        c,
		recver:        s,
		dispatcher:    d,
		logger:        o.Logger,
		EnableForward: o.EnableForward,
	}
	if o.EnableForward {
		cf := osc.NewClient(o.Host, o.SendPort)
		fdSource := newForwardDispatcher(o.Logger, cf, d)
		s.Dispatcher = fdSource
		fdTarget := newForwardDispatcher(o.Logger, c, nil)
		v.recverF = &osc.Server{
			Addr:       fmt.Sprintf("%s:%d", o.Host, o.ForwardRecvPort),
			Dispatcher: fdTarget,
		}
	}
	return v
}

func (v *VRCOsc) Run() error {
	p := pool.New().WithContext(context.Background()).WithCancelOnError()
	p.Go(func(_ context.Context) error {
		err := v.recver.ListenAndServe()
		if err != nil {
			return err
		}
		return nil
	})
	if v.EnableForward {
		p.Go(func(_ context.Context) error {
			err := v.recverF.ListenAndServe()
			if err != nil {
				return err
			}
			return nil
		})
	}
	err := p.Wait()
	if err != nil {
		return err
	}
	return nil
}
