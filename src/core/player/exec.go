package player

import (
	"dps/src/core/action"
	character2 "dps/src/core/character"
	"dps/src/core/combat"
	"dps/src/core/event"
	"dps/src/core/player/character"
	"fmt"
)

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

func UseAbility(c character.Character, ActionName string, cfg character2.Charcfg) combat.AttackInfo {
	ai := combat.AttackInfo{}
	if c == nil {
		fmt.Println("Error: Character is nil")
		return ai
	}
	switch ActionName {
	case "ActionSkill":
		ai = c.Skill(cfg)
	}
	return ai
}
