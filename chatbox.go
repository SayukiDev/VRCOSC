package VRCOSC

import (
	"fmt"
	"github.com/hypebeast/go-osc/osc"
)

func (v *VRCOsc) ChatBoxInput(s string, b, n bool) (err error) {
	r := []rune(s)
	if len(r) > 144 {
		v.logger.Println("Warn: [ Send chat box error: the chat box message too long ]")
		s = string(r[0:143])
	}
	m := osc.NewMessage("/chatbox/input", s, b, n)
	err = v.sender.Send(m)
	if err != nil {
		return fmt.Errorf("send packet error: %s", err)
	}
	return err
}
