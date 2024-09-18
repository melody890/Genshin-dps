package info

import "dps/src/core/attributes"

type Sets struct {
	SetNum        [2]int //每种圣遗物有几个
	SetName       [2]string
	SetBonusState [attributes.EndStatType]float64
}

type Set interface {
	SetIndex(int)
	Init() error
}
