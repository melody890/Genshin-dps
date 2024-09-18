package frames

import (
	"dps/src/core/action"
)

func InitAbilSlice(animation int) []int {
	t := make([]int, action.EndActionType)
	for i := range t {
		t[i] = animation
	}
	return t
}

func InitNormalCancelSlice(hitmark, animation int) []int {
	t := make([]int, action.EndActionType)
	for i := range t {
		t[i] = animation
	}
	t[action.ActionAim] = hitmark
	t[action.ActionSkill] = hitmark
	t[action.ActionBurst] = hitmark
	t[action.ActionDash] = hitmark
	t[action.ActionJump] = hitmark
	t[action.ActionSwap] = hitmark
	return t
}
