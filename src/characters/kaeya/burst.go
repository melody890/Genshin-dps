package kaeya

import (
	"dps/src/core/action"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/frames"
	character3 "dps/src/core/player/character"
	enemy "dps/src/enemies"
)

var burstFrames []int

const (
	burstCDStart  = 48
	burstHitmark  = 53
	burstDuration = 480
	burstKey      = "kaeya-q"
)

func init() {
	burstFrames = frames.InitAbilSlice(78) // Q -> E
	burstFrames[action.ActionAttack] = 77  // Q -> N1
	burstFrames[action.ActionDash] = 62    // Q -> D
	burstFrames[action.ActionJump] = 61    // Q -> J
	burstFrames[action.ActionSwap] = 77    // Q -> Swap
}

func (c Kaeya) Burst(cfg *[4]character3.CharWrapper, charindex int, dotAction *[]combat.AttackInfo, frame int, enemy *enemy.Enemy) []combat.AttackInfo {
	//var ai combat.AttackInfo
	for i := range cfg {
		if i != charindex {
			(*cfg)[i].Active = false
		} else {
			(*cfg)[i].Active = true
		}
	}
	ai := combat.AttackInfo{
		ActorIndex:      charindex,
		Abil:            "kaeyaActionBurst",
		Element:         attributes.Cryo,
		ElementQuantity: 0,
		Mult:            0,
		MultType:        []string{"atk"},
	}
	cfg[charindex].Energy = 0

	// 每个冰棱120帧的时候转一圈，但是内置最短CD是0.5秒
	count := 3 //冰棱数
	if cfg[charindex].Base.Cons >= 6 {
		count = 4
	}
	offset := 120 / count      // 每个冰棱之间的帧数偏移量，每offset帧会触发一次冰棱伤害。
	for i := range 4 * count { // 大招持续8秒，每2秒转一圈，所以一个冰棱有4次伤害.
		dotAI := combat.AttackInfo{
			ActorIndex:      charindex,
			Abil:            "kaeyaDotBurst",
			Element:         attributes.Cryo,
			ElementQuantity: 1,
			Mult:            burst[cfg[charindex].Talents.Burst-1],
			MultType:        []string{"atk"},
			AttackFrame:     offset*i + frame + burstHitmark,
			Lock:            true,
		}
		if i != 0 {
			dotAI.AttackFrame -= 20
		}
		*dotAction = append(*dotAction, dotAI)
	}
	var ais []combat.AttackInfo
	ais = append(ais, ai)
	return ais
}

func (c Kaeya) BurstMod(startframe int, cfg *[4]character3.CharWrapper, attack []combat.AttackInfo, enemy *enemy.Enemy, charIndex int) {
}

func (c Kaeya) BurstElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Kaeya) DotBurstElementAttach(enemy *enemy.Enemy, timer []int) []bool {
	var triggers []bool
	return triggers
}

func (c Kaeya) DotBurstElement(enemy *enemy.Enemy, timer []int) []bool {
	var triggers []bool
	timeTag := 0
	count := 1
	// 初始化所有凯亚大招dot
	for _ = range 12 {
		(*enemy).ElementAttach.ActionName = append((*enemy).ElementAttach.ActionName, "kaeyaDotBurst")
		(*enemy).ElementAttach.Counter = append((*enemy).ElementAttach.Counter, 1)
		(*enemy).ElementAttach.Timer = append((*enemy).ElementAttach.Timer, 0)
	}
	for i := range enemy.ElementAttach.ActionName {
		if enemy.ElementAttach.ActionName[i] == "kaeyaDotBurst" {
			//按照timer更新计时器
			(*enemy).ElementAttach.Timer[i] = timer[timeTag]
			timeTag++
			//如果计时器更新，那么trigger是true
			if (*enemy).ElementAttach.Timer[i] == 0 {
				(*enemy).ElementAttach.Counter[i] = 1
				count = 1
				triggers = append(triggers, true)
			} else {
				//如果计时器没更新，那么计数器+1
				(*enemy).ElementAttach.Counter[i] = count + 1
				count++
				if (*enemy).ElementAttach.Counter[i]%3 == 1 {
					triggers = append(triggers, true)
				} else {
					triggers = append(triggers, false)
				}
			}
		}
	}
	return triggers
}

func (c Kaeya) DotBurstElementTrans(enemy *enemy.Enemy, timer []int) []bool {
	var triggers []bool
	return triggers
}
