package VRCOSC

import (
	"github.com/hypebeast/go-osc/osc"
	"path"
	"path/filepath"
)

func (v *VRCOsc) AddParamsHandler(params string, handler osc.HandlerFunc) error {
	err := v.AddPathHandler(path.Join("/avatar/parameters", params), handler)
	if err != nil {
		return err
	}
	return nil
}

func (v *VRCOsc) AddPathHandler(path string, handler osc.HandlerFunc) error {
	err := v.dispatcher.AddMsgHandler(path, handler)
	if err != nil {
		return err
	}
	return nil
}

type HandlerGroup struct {
	path       string
	middleware []MiddlewareFunc
	handlerManager
}

type handlerManager interface {
	AddParamsHandler(params string, handler osc.HandlerFunc) error
	AddPathHandler(path string, handler osc.HandlerFunc) error
}

type MiddlewareFunc func(m *osc.Message) (break1 bool)

func (v *VRCOsc) HandlerGroup(path string, middleware ...MiddlewareFunc) *HandlerGroup {
	return &HandlerGroup{
		handlerManager: v,
		path:           path,
		middleware:     middleware,
	}
}

func (g *HandlerGroup) buildGroupHandler(h osc.HandlerFunc) osc.HandlerFunc {
	return func(msg *osc.Message) {
		for _, m := range g.middleware {
			b := m(msg)
			if b {
				return
			}
		}
		h(msg)
	}
}

func (g *HandlerGroup) AddParamsHandler(params string, handler osc.HandlerFunc) error {
	return g.handlerManager.AddParamsHandler(params, g.buildGroupHandler(handler))
}

func (g *HandlerGroup) AddPathHandler(path string, handler osc.HandlerFunc) error {
	return g.handlerManager.AddPathHandler(filepath.Join(path), g.buildGroupHandler(handler))
}
