package diluc

import (
	"dps/src/core/action"
	"dps/src/core/attributes"
	"dps/src/core/combat"
	"dps/src/core/frames"
	character3 "dps/src/core/player/character"
	enemy "dps/src/enemies"
	"dps/src/energy"
	character2 "dps/src/template/character"
	"dps/src/template/modifier"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	skillFrames       [][]int
	skillHitmarks     = []int{24, 28, 46}
	skillHitlagStages = []float64{.12, .12, .16}
	skillHitboxes     = [][]float64{{3, 3.5}, {2.2}, {3.5, 4}}
	skillOffsets      = []float64{0, 1.2, -0.3}
	skillFanAngles    = []float64{360, 300, 360}
)

var (
	SkillCD      = 10 * 60
	SkillGroupCD = 4 * 60
)

const particleICDKey = "diluc-particle-icd"

func init() {
	skillFrames = make([][]int, 3)

	// skill (1st) -> x
	skillFrames[0] = frames.InitAbilSlice(32)
	skillFrames[0][action.ActionSkill] = 31
	skillFrames[0][action.ActionDash] = 24
	skillFrames[0][action.ActionJump] = 24
	skillFrames[0][action.ActionSwap] = 30

	// skill (2nd) -> x
	skillFrames[1] = frames.InitAbilSlice(38)
	skillFrames[1][action.ActionSkill] = 37
	skillFrames[1][action.ActionBurst] = 37
	skillFrames[1][action.ActionDash] = 28
	skillFrames[1][action.ActionJump] = 31
	skillFrames[1][action.ActionSwap] = 36

	// skill (3rd) -> x
	// TODO: missing counts for skill -> skill
	skillFrames[2] = frames.InitAbilSlice(66)
	skillFrames[2][action.ActionAttack] = 58
	skillFrames[2][action.ActionSkill] = 57 // uses burst frames
	skillFrames[2][action.ActionBurst] = 57
	skillFrames[2][action.ActionDash] = 47
	skillFrames[2][action.ActionJump] = 48
}

func (c Diluc) SkillPress(cfg *[4]character3.CharWrapper, charindex int, enemy *enemy.Enemy, frame int) []combat.AttackInfo {
	var ai []combat.AttackInfo
	return ai
}

func (c Diluc) SkillHold(cfg *[4]character3.CharWrapper, charindex int, enemy *enemy.Enemy, frame int) []combat.AttackInfo {
	var ais []combat.AttackInfo
	return ais
}

func (c Diluc) Skill(cfg *[4]character3.CharWrapper, charindex int, startFrame int) []combat.AttackInfo {
	//判断CD是不是不好，不好的话报错
	if cfg[charindex].SkillCounter == 1 {
		//第一个e，判断是否在cd里
		if cfg[charindex].SkillFrame == -1 {
			(*cfg)[charindex].SkillFrame = startFrame
			(*cfg)[charindex].SkillGroupFrames[0] = startFrame
		} else if startFrame > cfg[charindex].SkillFrame+SkillCD {
			(*cfg)[charindex].SkillFrame = startFrame
			(*cfg)[charindex].SkillGroupFrames[0] = startFrame
		} else {
			fmt.Println("迪卢克技能冷却中")
			os.Exit(0)
		}
	} else if cfg[charindex].SkillCounter == 2 {
		//第二个e，如果还在初始化状态，判断是否在CD。如果超过了10s那他就是下一个e.否则报错
		if startFrame < cfg[charindex].SkillGroupFrames[0]+SkillGroupCD {
			(*cfg)[charindex].SkillGroupFrames[1] = startFrame
		} else if startFrame > cfg[charindex].SkillGroupFrames[0]+SkillCD {
			(*cfg)[charindex].SkillFrame = startFrame
			(*cfg)[charindex].SkillGroupFrames[0] = startFrame
			(*cfg)[charindex].SkillCounter = 1
		} else {
			fmt.Printf("迪卢克二段e冷却中")
			os.Exit(0)
		}
	} else if cfg[charindex].SkillCounter == 2 {
		//第三个e，如果还在初始化状态，判断是否在CD。如果超过了10s那他就是下一个e.否则报错
		if startFrame < cfg[charindex].SkillGroupFrames[1]+SkillGroupCD {
			(*cfg)[charindex].SkillGroupFrames[2] = startFrame
		} else if startFrame > cfg[charindex].SkillGroupFrames[1]+SkillCD {
			(*cfg)[charindex].SkillFrame = startFrame
			(*cfg)[charindex].SkillGroupFrames[0] = startFrame
			(*cfg)[charindex].SkillCounter = 1
		} else {
			fmt.Printf("迪卢克三段e冷却中")
			os.Exit(0)
		}
	}
	//调整在场角色
	for i := range cfg {
		if i != charindex {
			(*cfg)[i].Active = false
		} else {
			(*cfg)[i].Active = true
		}
	}
	cfg[charindex].Active = true
	ai := combat.AttackInfo{
		ActorIndex:      charindex,
		Abil:            "dilucActionSkill",
		Element:         attributes.Pyro,
		ElementQuantity: 1,
		Mult:            skill[cfg[charindex].SkillCounter-1][cfg[charindex].Talents.Skill-1],
		MultType:        []string{"atk"},
	}
	// 检查四命
	for i := range cfg[charindex].StatusMods {
		if cfg[charindex].StatusMods[i].ModName == "DilucC4Skill1" && cfg[charindex].SkillCounter == 2 && startFrame > cfg[charindex].StatusMods[i].StartFrame && startFrame < cfg[charindex].StatusMods[i].StartFrame+cfg[charindex].StatusMods[i].Dur {
			mod := modifier.Mods{
				Name:       "DilucC4Skill1",
				Modifier:   [attributes.EndStatType]float64{},
				StartFrame: startFrame,
				Dur:        skillFrames[1][action.ActionDelay],
				CharIndex:  charindex,
			}
			mod.Modifier[attributes.DmgP] = 0.4
			(*cfg)[charindex].Mods = append((*cfg)[charindex].Mods, mod)
			// 删除Statesmod
			for j := len(cfg[charindex].StatusMods) - 1; j >= 0; j-- {
				mod1 := cfg[charindex].StatusMods[j]
				if mod1.ModName == "DilucC4Skill1" {
					cfg[charindex].StatusMods = append(cfg[charindex].StatusMods[:j], cfg[charindex].StatusMods[j+1:]...)
				}
			}
		} else if cfg[charindex].StatusMods[i].ModName == "DilucC4Skill2" && cfg[charindex].SkillCounter == 3 && startFrame > cfg[charindex].StatusMods[i].StartFrame && startFrame < cfg[charindex].StatusMods[i].StartFrame+cfg[charindex].StatusMods[i].Dur {
			mod := modifier.Mods{
				Name:       "DilucC4Skill2",
				Modifier:   [attributes.EndStatType]float64{},
				StartFrame: startFrame,
				Dur:        skillFrames[2][action.ActionDelay],
				CharIndex:  charindex,
			}
			mod.Modifier[attributes.DmgP] = 0.4
			(*cfg)[charindex].Mods = append((*cfg)[charindex].Mods, mod)
			// 删除Statesmod
			for j := len(cfg[charindex].StatusMods) - 1; j >= 0; j-- {
				mod1 := cfg[charindex].StatusMods[j]
				if mod1.ModName == "DilucC4Skill2" {
					cfg[charindex].StatusMods = append(cfg[charindex].StatusMods[:j], cfg[charindex].StatusMods[j+1:]...)
				}
			}
		}
	}

	//充能
	count := 1
	rand.Seed(time.Now().UnixNano())
	if rand.Float64() < 0.33 {
		count = 2
	}
	ParticleElement := attributes.Pyro
	RechargeType := "particle"
	elementInfo := character2.EnergyInfo{
		Count: count, ParticleElement: ParticleElement, RechargeType: RechargeType,
	}
	energy.EnergyRecharge(cfg, elementInfo)
	switch cfg[charindex].SkillCounter {
	case 1:
		ai.Abil = "dilucActionSkill1"
	case 2:
		ai.Abil = "dilucActionSkill2"
	case 3:
		ai.Abil = "dilucActionSkill3"
	case 4:
		ai.Abil = "dilucActionSkill4"
	default:
		ai.Abil = "dilucActionSkill"
	}
	var ais []combat.AttackInfo
	ais = append(ais, ai)
	return ais
}

func (c Diluc) SkillMod(cfg *[4]character3.CharWrapper, charIndex int, startframe int, enemy *enemy.Enemy) {
	c1(cfg, enemy, startframe, charIndex)
	// 四命:有节奏地释放逆焰之刃可以大幅提升造成的伤害。施放逆焰之刃地2秒后，使下一段逆焰之刃地伤害提高40%，持续2秒
	// TODO:这里要测试一下，第一段的这个buff能不能不给第二段而给第三段，虽然我猜不行，先按不行来写
	if cfg[charIndex].SkillCounter == 1 {
		character3.AddStatus("DilucC4Skill1", 120, startframe+120, cfg, charIndex, 0, 1)
	} else if cfg[charIndex].SkillCounter == 2 {
		character3.AddStatus("DilucC4Skill2", 120, startframe+120, cfg, charIndex, 0, 1)
	}
}

func (c Diluc) SkillPressMod() {
}

func (c Diluc) SkillHoldMod() {

}

func (c Diluc) SkillElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Diluc) SkillPressElementAttach(enemy *enemy.Enemy) bool {
	return true
}

func (c Diluc) SkillHoldElementAttach(enemy *enemy.Enemy) bool {
	return true
}
