package character

import (
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/info"
	"dps/src/core/player/infusion"
	enemy "dps/src/enemies"
	"dps/src/template/modifier"
)

type Character interface {
	FindElement() attributes.Element

	InitEnergy() float64
	NormalHitNum() int
	UpdateStates(cfg *[4]CharWrapper, i int)

	// FindFrame 这里注意，找的时候找的是下一个动作，如果下一个动作是attack，那么这里还没有cfg.normalcount+1，判断的时候要退一位判断
	FindFrame(aipre combat.AttackInfo, ActorIndex int, Abil string, cfg [4]CharWrapper) int
	// FindAttackFrame 找到刚好出伤的那个点
	FindAttackFrame(ai combat.AttackInfo) int

	Skill(cfg *[4]CharWrapper, charindex int, frame int) []combat.AttackInfo
	SkillPress(cfg *[4]CharWrapper, charindex int, enemy *enemy.Enemy, frame int) []combat.AttackInfo
	SkillHold(cfg *[4]CharWrapper, charindex int, enemy *enemy.Enemy, frame int) []combat.AttackInfo
	Attack(cfg *[4]CharWrapper, charindex int, frame int, enemy *enemy.Enemy) []combat.AttackInfo
	Burst(cfg *[4]CharWrapper, charindex int, dotAction *[]combat.AttackInfo, frame int, enemy *enemy.Enemy) []combat.AttackInfo
	HighHoldPlunge(cfg *[4]CharWrapper, charindex int, dotAction *[]combat.AttackInfo, frame int, enemy *enemy.Enemy) []combat.AttackInfo
	HighPressPlunge(cfg *[4]CharWrapper, charindex int, dotAction *[]combat.AttackInfo, frame int, enemy *enemy.Enemy) []combat.AttackInfo

	AttackMod(startframe int, cfg *[4]CharWrapper, attack []combat.AttackInfo, enemy *enemy.Enemy, charIndex int)
	SkillMod(cfg *[4]CharWrapper, charIndex int, startframe int, enemy *enemy.Enemy)
	SkillPressMod()
	SkillHoldMod()
	BurstMod(startframe int, cfg *[4]CharWrapper, attacks []combat.AttackInfo, enemy *enemy.Enemy, charIndex int)
	PlungeHighMod(cfg *[4]CharWrapper, charIndex int, startframe int, enemy *enemy.Enemy, attacks []combat.AttackInfo)

	AttackElementAttach(enemy *enemy.Enemy) bool
	SkillElementAttach(enemy *enemy.Enemy) bool
	SkillPressElementAttach(enemy *enemy.Enemy) bool
	SkillHoldElementAttach(enemy *enemy.Enemy) bool
	BurstElementAttach(enemy *enemy.Enemy) bool
	PlungeHoldElementAttach(enemy *enemy.Enemy) bool
	PlungePressElementAttach(enemy *enemy.Enemy) bool
	DotBurstElement(enemy *enemy.Enemy, timer []int) []bool
	DotBurstElementAttach(enemy *enemy.Enemy, timer []int) []bool
	DotBurstElementTrans(enemy *enemy.Enemy, timer []int) []bool
}

type CharWrapper struct {
	Index int
	f     *int // current frame
	Character
	Active bool
	//events event.Eventter

	// base characteristics
	Base    info.CharacterBase
	Weapon  info.WeaponProfile
	Talents info.TalentProfile
	//NormalCon int
	//SkillCon  int
	//BurstCon  int
	//HasArkhe  bool

	Equip struct {
		Weapon info.WeaponProfile
		Sets   info.Sets
	}

	// current status
	ParticleDelay int // character custom particle delay
	Energy        float64
	EnergyMax     float64

	// needed so that start hp is not influenced by hp mods added during team initialization
	StartHP      int
	StartHPRatio int

	// normal attack counter
	NormalHitNum  int // how many hits in a normal combo
	NormalCounter int
	SkillCounter  int

	// tags
	Tags      map[string]int
	BaseStats [attributes.EndStatType]float64

	// mods
	Mods       []modifier.Mods
	StatusMods []modifier.Base

	//角色身上元素附着
	ElementAttachment attributes.Element

	//技能CD品鉴
	SkillFrame       int   // 上一个技能释放的帧数
	SkillGroupFrames []int // 每一个技能组里面的技能释放帧数，比如迪卢克，刻晴，公子
	BurstFrame       int   // 大招
	BurstGroupFrames []int // 大招技能组，比如迪希雅之类

	// dash cd: keeps track of remaining cd frames for off-field chars
	RemainingDashCD int
	DashLockout     bool

	// hitlag stuff
	timePassed   int // how many frames have passed since start of sim
	frozenFrames int // how many frames are we still frozen for

	// 附魔
	Infusion []infusion.WeaponInfusion
}
