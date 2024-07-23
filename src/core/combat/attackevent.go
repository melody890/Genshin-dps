package combat

import (
	"dps/src/core/attacks"
	"dps/src/core/attributes"
)

type AttackInfo struct {
	ActorName        string             // 释放攻击角色名
	Abil             string             // name of ability triggering the damage
	PoiseDMG         float64            // 目前仅在冰冻状态下的钝器攻击中需要，以在破碎之前消耗冰冻效果
	Element          attributes.Element // element of ability
	NoImpulse        bool
	HitWeakPoint     bool
	Mult             float64 // ability multiplier. could set to 0 from initial Mona dmg
	Durability       float64
	StrikeType       attacks.StrikeType
	UseDef           bool    // we use this instead of flatdmg to make sure stat snapshotting works properly
	FlatDmg          float64 // flat dmg;
	IgnoreDefPercent float64 // by default this value is 0; if = 1 then the attack will ignore defense; raiden c2 should be set to 0.6 (i.e. ignore 60%)
	IgnoreInfusion   bool
	// catalyze info
	SourceIsSim bool
	DoNotLog    bool
	// hitlag stuff
	HitlagHaltFrames     float64 // this is the number of frames to pause by
	HitlagFactor         float64 // this is factor to slow clock by
	CanBeDefenseHalted   bool    // for whacking ruin gaurds
	IsDeployable         bool    // if this is true, then hitlag does not affect owner
	HitlagOnHeadshotOnly bool    // if this is true, will only apply if HitWeakpoint is also true
}
