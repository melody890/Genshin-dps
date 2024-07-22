package player

import (
	"dps/src/core/action"
	"dps/src/core/event"
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

func UseAbility(
	f func(),
) {
	f()
}
