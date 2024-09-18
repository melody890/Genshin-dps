package player

import (
	"dps/src/core/action"
	"dps/src/core/combat"
	"dps/src/core/event"
	"dps/src/core/player/character"
	enemy "dps/src/enemies"
	"log"
	"strings"
)

type Actions struct {
	Characters []string
	Actions    [][]string
	Delay      []int //用于调整动作间隔
}

var actionToEvent = map[action.Action]event.Event{
	action.ActionDash:       event.OnDash,
	action.ActionSkill:      event.OnSkill,
	action.ActionBurst:      event.OnBurst,
	action.ActionAttack:     event.OnAttack,
	action.ActionCharge:     event.OnChargeAttack,
	action.ActionLowPlunge:  event.OnPlunge,
	action.ActionHighPlunge: event.OnPlunge,
	action.ActionAim:        event.OnAimShoot,
}

func UseAbility(c character.Character, ActionName string, cfg *[4]character.CharWrapper, charindex int, actions Actions, framei int, enemy *enemy.Enemy, dotAction *[]combat.AttackInfo) []combat.AttackInfo {
	//ae := combat.AttackEvent{}
	var ai []combat.AttackInfo
	if c == nil {
		log.Println("Error: Character is nil")
		return ai
	}
	if ActionName == "ActionAttack" {
		cfg[charindex].NormalCounter += 1
		for j := range 4 {
			if j != charindex {
				cfg[j].NormalCounter = 0
			}
		}
	} else {
		for j := range 4 {
			cfg[j].NormalCounter = 0
		}
	}
	if ActionName == "ActionSkill" {
		cfg[charindex].SkillCounter += 1
		for j := range 4 {
			if j != charindex {
				cfg[j].SkillCounter = 0
			}
		}
	}
	switch ActionName {
	case "ActionSkill":
		ai = c.Skill(cfg, charindex, framei)
	case "ActionSkillPress":
		ai = c.SkillPress(cfg, charindex, enemy, framei)
	case "ActionSkillHold":
		ai = c.SkillHold(cfg, charindex, enemy, framei)
	case "ActionAttack":
		ai = c.Attack(cfg, charindex, framei, enemy)
	case "ActionBurst":
		ai = c.Burst(cfg, charindex, dotAction, framei, enemy)
	case "ActionPlungeHold":
		ai = c.HighHoldPlunge(cfg, charindex, dotAction, framei, enemy)
	case "ActionPlungePress":
		ai = c.HighPressPlunge(cfg, charindex, dotAction, framei, enemy)
	}
	return ai
}

func ApplyCharMod(c character.Character, attacks []combat.AttackInfo, framei int, actions Actions, cfg *[4]character.CharWrapper, enemy *enemy.Enemy, actionNum int) {
	charIndex := attacks[actionNum].ActorIndex
	switch actions.Actions[actionNum][1] {
	case "ActionAttack":
		c.AttackMod(framei, cfg, attacks, enemy, charIndex)
	case "ActionSkill":
		c.SkillMod(cfg, charIndex, framei, enemy)
	case "ActionSkillPress":
		c.SkillPressMod()
	case "skill hold":
		c.SkillHoldMod()
	case "ActionBurst":
		c.BurstMod(attacks[len(attacks)-1].Startframe, cfg, attacks, enemy, charIndex)
	case "ActionPlungePress":
		c.PlungeHighMod(cfg, charIndex, framei, enemy, attacks)
	}
	//log.Println(attacks[len(attacks)-1].Abil, " mod ", mods)
}

func UseElementAttach(c character.Character, attacks *[]combat.AttackInfo, enemy *enemy.Enemy, index int) {
	actionName := (*attacks)[index].Abil
	actionIndex := strings.Index(actionName, "Action")
	var actionResult string
	var elementTrigger bool
	elementTrigger = true
	if actionIndex != -1 {
		// 截取从 "Action" 开始的子字符串
		actionResult = actionName[actionIndex:]
		if actionResult[len(actionResult)-1] == '1' || actionResult[len(actionResult)-1] == '2' || actionResult[len(actionResult)-1] == '3' || actionResult[len(actionResult)-1] == '4' || actionResult[len(actionResult)-1] == '5' {
			actionResult = actionResult[:len(actionResult)-1]
		}
	}
	switch actionResult {
	case "ActionAttack":
		elementTrigger = c.AttackElementAttach(enemy)
	case "ActionSkill":
		elementTrigger = c.SkillElementAttach(enemy)
	case "ActionSkillPress":
		elementTrigger = c.SkillPressElementAttach(enemy)
	case "ActionSkillHold":
		elementTrigger = c.SkillHoldElementAttach(enemy)
	case "ActionBurst":
		elementTrigger = c.BurstElementAttach(enemy)
	case "ActionActionPlungeHold":
		elementTrigger = c.PlungeHoldElementAttach(enemy)
	case "ActionActionPlungePress":
		elementTrigger = c.PlungePressElementAttach(enemy)
	}
	(*attacks)[index].Trigger = elementTrigger
}

func UseDotElementAttach(c character.Character, attacks *[]combat.AttackInfo, enemy *enemy.Enemy, CharName string, dotLen int) {
	//为了只更新新来的dot，我们从后往前查找新进来的动作，它们的name是一样的
	var index int
	nowAttacks := (*attacks)[len(*attacks)-dotLen:]
	groupedAttacks := groupByAbil(nowAttacks)
	if len(nowAttacks) == 0 {
		return
	}
	for _, nowLittleAttack := range groupedAttacks {
		//更新index，将index更新为该dot的首位index
		index = findIndex(*attacks, nowLittleAttack[0])
		actionName := nowLittleAttack[0].Abil

		var dotResult string
		var actorName string
		dotIndex := strings.Index(actionName, "Dot")
		if dotIndex != -1 {
			dotResult = actionName[dotIndex:]
			actorName = actionName[:dotIndex]
		}
		var timer []int
		//只更新新来的这些个dot
		if actorName == CharName {
			time0 := nowLittleAttack[0].AttackFrame
			timer = append(timer, 0)
			for i := range nowLittleAttack {
				if i != 0 {
					intervalTime := nowLittleAttack[i].AttackFrame - time0
					if intervalTime >= 150 {
						time0 = nowLittleAttack[i].AttackFrame
						intervalTime = 0
					}
					timer = append(timer, intervalTime)
				}
			}
			var triggers []bool
			switch dotResult {
			// index是dotAction的index
			case "DotBurst":
				triggers = c.DotBurstElement(enemy, timer)
			case "DotBurstAttach":
				triggers = c.DotBurstElementAttach(enemy, timer)
			}
			trigI := 0
			if len(triggers) != 0 {
				for i := index; i < len(nowLittleAttack)+index; i++ {
					(*attacks)[i].Trigger = triggers[trigI]
					trigI++
				}
			}
		}
	}
}

func groupByAbil(attacks []combat.AttackInfo) map[string][]combat.AttackInfo {
	grouped := make(map[string][]combat.AttackInfo)

	for _, attack := range attacks {
		abil := attack.Abil
		grouped[abil] = append(grouped[abil], attack)
	}

	return grouped
}

func findIndex(original []combat.AttackInfo, target combat.AttackInfo) int {
	for i, attack := range original {
		if attack.Abil == target.Abil {
			return i
		}
	}
	return -1 // 如果未找到返回 -1
}

func ApllyLockDot(cfg *[4]character.CharWrapper, dotAction *[]combat.AttackInfo, charindex int, dotLen int, frame int) {
	nowDotAttacks := (*dotAction)[len(*dotAction)-dotLen:]
	lenDots := len(*dotAction)
	for i := range nowDotAttacks {
		if nowDotAttacks[i].Lock == true && (*dotAction)[lenDots-dotLen+i].LockUpdate == false {
			for j := range cfg[charindex].Mods {
				if cfg[charindex].Mods[j].StartFrame <= frame {
					(*dotAction)[lenDots-dotLen+i].LockMod = append((*dotAction)[lenDots-dotLen+i].LockMod, cfg[charindex].Mods[j])
					(*dotAction)[lenDots-dotLen+i].LockUpdate = true
				}
			}
		}
	}
}
