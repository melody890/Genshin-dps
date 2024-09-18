package combat

import (
	"dps/src/core/attributes"
	"dps/src/reactable"
	"dps/src/template/modifier"
)

type AttackEvent struct {
	Info AttackInfo
	// Timing        AttackTiming
	Snapshot Snapshot
	Reacted  bool // true if a reaction already took place - for purpose of attach/refill
}

type AttackInfo struct {
	ActorIndex int    // 释放攻击角色序号
	Abil       string // name of ability triggering the damage
	//PoiseDMG  float64            // 目前仅在冰冻状态下的钝器攻击中需要，以在破碎之前消耗冰冻效果

	Element         attributes.Element // 是什么元素
	ElementQuantity float64            // 元素量
	Trigger         bool               //该次攻击是否有元素附着

	Mult     float64  // 这里是代表倍率具体数值
	MultType []string // 这里是代表用什么计算倍率

	//Durability float64
	//StrikeType       attacks.StrikeType
	UseHP            bool
	UseDef           bool    // we use this instead of flatdmg to make sure stat snapshotting works properly
	FlatDmg          float64 // flat dmg;
	IgnoreDefPercent float64 // by default this value is 0; if = 1 then the attack will ignore defense; raiden c2 should be set to 0.6 (i.e. ignore 60%)
	Amped            bool    // 增幅反应
	Trans            bool    // 剧变反应
	AmpMult          float64 // 增幅反应倍率
	ReactBonus       [reactable.EndReaction]float64

	Tags []string // 代表一些奇怪的加成的名字，比如风套的扩散反应伤害加成

	Reacted bool // true if a reaction already took place - for purpose of attach/refill

	Lock       bool            //对于Dot的锁面板
	LockMod    []modifier.Mods // 锁住Dot的时候的buff
	LockUpdate bool

	Startframe  int // 动作开始的帧数
	AttackFrame int
	Endframe    int
}

type Snapshot struct {
	CharLvl    int
	ActorEle   attributes.Element
	ExtraIndex int                             // this is currently purely for Kaeya icicle ICD
	Cancelled  bool                            // set to true if this snap should be ignored
	Stats      [attributes.EndStatType]float64 // total character stats including from artifact, bonuses, etc...
	BaseAtk    float64                         // base attack used in calc
	BaseDef    float64
	BaseHP     float64

	SourceFrame int // frame snapshot was generated at
}
