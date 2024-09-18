package modifier

import (
	"dps/src/core/attacks"
	"dps/src/core/attributes"
)

type Mods struct {
	Name       string                          //加成的名字
	Modifier   [attributes.EndStatType]float64 // 具体加成的内容
	StartFrame int                             // 具体从哪一帧开始
	Dur        int                             // 持续的帧数, -1表示永久
	CharIndex  int                             // 队内哪个角色的加成，1-4代表角色号，5代表前台，6代表全体
	Condition  attacks.AttackTag
}

type Base struct {
	ModName    string
	Dur        int
	StartFrame int
	Condition  attacks.AttackTag
	CD         int
}

type EnemyMod struct {
	Name       string
	Res        [attributes.EndEleType]float64 // 减抗
	Def        float64                        // 减防
	StartFrame int
	Dur        int
}
