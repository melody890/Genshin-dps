package attacks

type AttackTag int // attacktag is used instead of actions etc..

const (
	AttackTagNone AttackTag = iota
	AttackTagNormal
	AttackTagExtra
	AttackTagPlunge
	AttackBurst
	AttackCharge
	AttackTagElementalArt
	AttackTagElementalArtHold
	AttackTagElementalBurst
	AttackTagWeaponSkill
	AttackTagMonaBubbleBreak
	AttackTagNoneStat
	ReactionAttackDelim
	AttackTagOverloadDamage
	AttackTagSuperconductDamage
	AttackTagECDamage
	AttackTagShatter
	AttackTagSwirlPyro
	AttackTagSwirlHydro
	AttackTagSwirlCryo
	AttackTagSwirlElectro
	AttackTagBurningDamage
	AttackTagBloom
	AttackTagBurgeon
	AttackTagHyperbloom
	AttackTagLength
)

type StrikeType int

const (
	StrikeTypeDefault StrikeType = iota
	StrikeTypePierce
	StrikeTypeBlunt
	StrikeTypeSlash
	StrikeTypeSpear
)
