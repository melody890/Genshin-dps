package character

import (
	character2 "dps/src/core/character"
	"dps/src/core/combat"
	"dps/src/template/character"
)

type Character interface {
	//Init() error // init function built into every char to setup any variables etc.
	InitChar(CharName string) *character.Character
	Skill(cfg character2.Charcfg) combat.AttackInfo
	//Attack(p map[string]int) (action.Info, error)
	//Aimed(p map[string]int) (action.Info, error)
	//ChargeAttack(p map[string]int) (action.Info, error)
	//HighPlungeAttack(p map[string]int) (action.Info, error)
	//LowPlungeAttack(p map[string]int) (action.Info, error)
	//Skill(p map[string]int) (action.Info, error)
	//Burst(p map[string]int) (action.Info, error)
	//Dash(p map[string]int) (action.Info, error)
	//Walk(p map[string]int) (action.Info, error)
	//Jump(p map[string]int) (action.Info, error)
	//
	//ActionStam(a action.Action, p map[string]int) float64
	//
	//ActionReady(a action.Action, p map[string]int) (bool, action.Failure)
	//SetCD(a action.Action, dur int)
	//Cooldown(a action.Action) int
	//ResetActionCooldown(a action.Action)
	//ReduceActionCooldown(a action.Action, v int)
	//Charges(a action.Action) int

	//AddEnergy(src string, amt float64)
	//
	//ApplyHitlag(factor, dur float64)
	//AnimationStartDelay(model.AnimationDelayKey) int
	//
	//Condition([]string) (any, error)

	//ResetNormalCounter()
	//NextNormalCounter() int
}
