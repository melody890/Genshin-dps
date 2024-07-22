package characters

import (
	"dps/src/characters/bennett"
	"dps/src/characters/kazuha"
	character2 "dps/src/core/character"
	"fmt"
)

type Character interface {
	ActionCharge()
	ActionDash()
	ActionJump()
	ActionWalk()
	ActionAim()
	ActionSkill()
	ActionBurst()
	ActionAttack()
	ActionHighPlunge()
	ActionLowPlunge()
	ActionSwap()
	ActionWait()
	ActionDelay()
	// Add other actions as needed
}

// 返回角色名，角色初始三围和突破加什么属性
func InitCharacter(name string) (character2.Base_attribute, string, error) {
	switch name {
	case "bennett":
		char, breakthrough := bennett.Init_bennett()
		return char, breakthrough, nil
	case "Kazuha":
		char, breakthrough := kazuha.Init_kazuha()
		return char, breakthrough, nil
	default:
		var zeroChar character2.Base_attribute
		return zeroChar, "", fmt.Errorf("unknown character: %s", name)
	}
}
