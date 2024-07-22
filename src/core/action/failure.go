package action

type Failure int

const (
	NoFailure Failure = iota
	SwapCD
	SkillCD
	BurstCD
	InsufficientEnergy
	InsufficientStamina
	CharacterDeceased // TODO: need chars to die first
	DashCD
)
