package VRCOSC

var v *VRCOsc

func init() {
	v = New(&Options{
		SendPort: 9000,
		RecvPort: 9001,
	})
}
