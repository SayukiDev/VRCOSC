package VRCOSC

import "github.com/hypebeast/go-osc/osc"

func (v *VRCOsc) SendRaw(packet osc.Packet) error {
	err := v.sender.Send(packet)
	if err != nil {
		return err
	}
	return nil
}
