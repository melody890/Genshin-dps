package modifier

import (
	"dps/src/core/attributes"
	character3 "dps/src/core/player/character"
	"dps/src/template/modifier"
)

func ElementResonance(cfg *[4]character3.CharWrapper) {
	var Elements [attributes.EndEleType]int
	var tag bool
	tag = false
	for i := range cfg {
		Elements[cfg[i].Base.Element] += 1
	}
	for i := range Elements {
		if Elements[i] == 2 {
			tag = true
			switch i {
			case 0:
			//雷
			case 1:
				FerventFlames(cfg)
			}
		}
	}
	if tag == false{
		//四杂色
	}
}

func FerventFlames(cfg *[4]character3.CharWrapper) {
	mod := modifier.Mods{
		Name:       "FerventFlames",
		Modifier:   [attributes.EndStatType]float64{},
		StartFrame: -1,
		Dur:        -1,
		CharIndex:  5,
	}
	mod.Modifier[attributes.ATKP] = 0.25
	for i := range cfg {
		(*cfg)[i].Mods = append((*cfg)[i].Mods, mod)
	}
}
