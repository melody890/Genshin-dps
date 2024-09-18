package modifier

import (
	"dps/src/core/attributes"
	"dps/src/template/modifier"
)

//所以我开一个新的参数叫mod，是指不在面板上体现，战斗里才体现的加成。跟states一样拥有那么多属性。然后每条mod额外多一个持续时间。每一次新动作检查是否触发4个武器，4个圣遗物，24个命座，8个天赋
// 记录每个圣遗物武器带来的各种加成，以及持续时间。

type Mod interface {
	AddMod(Modifier [attributes.EndStatType]float64, Dur int, CharIndex int) modifier.Mods
}
