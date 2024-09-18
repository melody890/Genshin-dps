package find

func FindIndex[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

//func FindAttackFrame(ai combat.AttackInfo, cfg [4]character.CharWrapper) int {
//	var c character2.Character
//	c = FindCharC(cfg[ai.ActorIndex].Base.CharName)
//	c.FindAttackFrame(ai)
//	return 0
//}
