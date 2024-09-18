package diluc

import (
	"dps/src/core/action"
	"dps/src/core/attacks"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/frames"
	character3 "dps/src/core/player/character"
	"dps/src/core/player/infusion"
	enemy "dps/src/enemies"
	"dps/src/template/modifier"
)

var burstFrames []int

const burstHitmark = 100

func init() {
	burstFrames = frames.InitAbilSlice(140) // Q -> D
	burstFrames[action.ActionAttack] = 139  // Q -> N1
	burstFrames[action.ActionSkill] = 139   // Q -> E
	burstFrames[action.ActionJump] = 139    // Q -> J
	burstFrames[action.ActionSwap] = 138    // Q -> Swap
}

func (c Diluc) Burst(cfg *[4]character3.CharWrapper, charindex int, dotAction *[]combat.AttackInfo, frame int, enemy *enemy.Enemy) []combat.AttackInfo {
	for i := range cfg {
		if i != charindex {
			(*cfg)[i].Active = false
		} else {
			(*cfg)[i].Active = true
		}
	}
	ai := combat.AttackInfo{
		ActorIndex:      charindex,
		Abil:            "dilucActionBurst",
		Element:         attributes.Pyro,
		ElementQuantity: 2,
		Mult:            burstInitial[cfg[charindex].Talents.Burst-1],
		MultType:        []string{"atk"},
	}
	cfg[charindex].Energy = 0
	//加一下迪卢克黎明的dot，这里不考虑爆裂伤害，打boss只有两段dot
	//大招一共持续1.7秒，其中有8段dot，间隔12帧
	offset := 12
	for i := range 2 {
		dotAI := combat.AttackInfo{
			ActorIndex:      charindex,
			Abil:            "dilucDotBurst",
			Element:         attributes.Pyro,
			ElementQuantity: 2,
			Mult:            burstDOT[cfg[charindex].Talents.Burst-1],
			MultType:        []string{"atk"},
			AttackFrame:     offset*(i+2) + frame + burstHitmark,
		}
		*dotAction = append(*dotAction, dotAI)
	}
	var ais []combat.AttackInfo
	ais = append(ais, ai)
	return ais
}

func (c Diluc) BurstMod(startframe int, cfg *[4]character3.CharWrapper, attack []combat.AttackInfo, enemy *enemy.Enemy, charIndex int) {
	//charIndex := attack[len(attack)-1].ActorIndex
	// 附魔
	class := []string{"claymore"}
	element := attributes.Pyro
	tags := []attacks.AttackTag{attacks.AttackTagNormal, attacks.AttackTagExtra, attacks.AttackTagPlunge}
	inf := infusion.AddWeaponInfuse("dilucBurstInfusion", class, element, tags, true, startframe, 480, charIndex)
	//突破四：黎明提供的火元素附魔持续时间延长4秒
	if (*cfg)[charIndex].Base.Ascension >= 4 {
		inf.Dur += 240
	}
	(*cfg)[charIndex].Infusion = append((*cfg)[charIndex].Infusion, inf)
	//突破四：黎明提供的火元素附魔持续时间延长4秒；此外，在效果持续期间，迪卢克获得20%火元素伤害加成。
	if (*cfg)[charIndex].Base.Ascension >= 4 {
		mod := modifier.Mods{
			Name:       "DilucBurstAscension4",
			Modifier:   [attributes.EndStatType]float64{},
			StartFrame: startframe,
			Dur:        720,
			CharIndex:  charIndex,
		}
		//fmt.Println("startframe:", mod.StartFrame)
		mod.Modifier[attributes.PyroP] = 0.2
		(*cfg)[charIndex].Mods = append((*cfg)[charIndex].Mods, mod)
	}

	//命座一：对生命值高于50%的敌人，造成的伤害提高15%
	c1(cfg, enemy, startframe, charIndex)
}

func (c Diluc) DotBurstElement(enemy *enemy.Enemy, timer []int) []bool {
	var triggers []bool
	return triggers
}

func (c Diluc) BurstElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Diluc) DotBurstElementAttach(enemy *enemy.Enemy, timer []int) []bool {
	var triggers []bool
	return triggers
}

func (c Diluc) DotBurstElementTrans(enemy *enemy.Enemy, timer []int) []bool {
	var triggers []bool
	return triggers

}
