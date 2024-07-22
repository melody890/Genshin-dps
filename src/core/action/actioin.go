package action

type Action int

const (
	InvalidAction Action = iota
	ActionSkill
	ActionBurst
	ActionAttack
	ActionCharge
	ActionHighPlunge
	ActionLowPlunge
	ActionAim
	ActionDash
	ActionJump
	// following action have to implementations
	ActionSwap
	ActionWalk
	ActionWait  // character should stand around and wait
	ActionDelay // delay before executing next action
	EndActionType
	// these are only used for frames purposes and that's why it's after end
	ActionSkillHoldFramesOnly
)

type Info struct {
	Frames              func(next Action) int
	AnimationLength     int
	CanQueueAfter       int
	State               AnimationState
	FramePausedOnHitlag func() bool               `json:"-"`
	OnRemoved           func(next AnimationState) `json:"-"`
	// following are exposed only so we can log it properly
	TimePassed           float64
	NormalizedTimePassed float64
	UseNormalizedTime    func(next Action) bool
	// hidden stuff
	queued []queuedAction
}

type queuedAction struct {
	f     func()
	delay float64
}

type AnimationState int

const (
	Idle AnimationState = iota
	NormalAttackState
	ChargeAttackState
	PlungeAttackState
	SkillState
	BurstState
	AimState
	DashState
	JumpState
	WalkState
	SwapState
)
