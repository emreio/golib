package golib_types

import "fmt"

type Universal struct {
	UUID         string
	Device       string
	Manufacturer int64
	Parent       interface{}
}

func (universal Universal) Describe() string {
	return fmt.Sprint(universal.UUID, universal.Manufacturer, universal.Device)
}
