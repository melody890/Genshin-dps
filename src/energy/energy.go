package energy

import (
	"dps/src/core/attributes"
	character2 "dps/src/core/player/character"
	character3 "dps/src/template/character"
)

type EnergyEvent struct {
	Frame       int         `json:"frame"        msg:"frame"`
	Source      string      `json:"source"       msg:"source"`
	FieldStatus FieldStatus `json:"field_status" msg:"field_status"`
	Gained      float64     `json:"gained"       msg:"gained"`
	Wasted      float64     `json:"wasted"       msg:"wasted"`
	Current     float64     `json:"current"      msg:"current"` // this is pre + gained
}

type FieldStatus string
type ReactionModifier string

const (
	OnField  FieldStatus = "on_field"
	OffField FieldStatus = "off_field"

	Melt      ReactionModifier = "melt"
	Vaporize  ReactionModifier = "vaporize"
	Spread    ReactionModifier = "spread"
	Aggravate ReactionModifier = "aggravate"
)

//func ApplyStatesEnergy(cfg character.Charcfg) [4]float64 {
//	var energyStates [4]float64
//	for i := range cfg.Characters {
//		c := find.FindCharC(cfg.Characters[i])
//		var energyState float64
//		if c != nil {
//			energyState = c.InitEnergy()
//		}
//		energyStates[i] = energyState
//	}
//
//	return energyStates
//}

//func ApplyStepEnergy(cfg *[4]character2.CharWrapper, elementInfo character3.EnergyInfo) {
//	for i := range *cfg {
//
//		ae.EnergyStates[i] += energyStates[i]
//		c := find.FindCharC(cfg.Characters[i])
//		if c != nil {
//			if ae.EnergyStates[i] > c.MaxEnergy() {
//				ae.EnergyStates[i] = c.MaxEnergy()
//			}
//		}
//	}
//}

func EnergyRecharge(cfg *[4]character2.CharWrapper, elementInfo character3.EnergyInfo) {
	// 非对应属性，前台角色，元素微粒作为基准，回复1点能量
	baseRecharge := elementInfo.Count
	for i := range *cfg {
		recharge := float64(baseRecharge)
		// 如果是晶球，充能翻三倍
		if elementInfo.RechargeType == "orb" {
			recharge *= 3
		}
		// 如果是同一个元素，充能翻三倍；如果是无属性，充能翻两倍
		charElementType := cfg[i].Base.Element
		if elementInfo.ParticleElement == charElementType {
			recharge *= 3
		} else if elementInfo.ParticleElement == attributes.NoElement {
			recharge *= 2
		}
		// 如果是后台人物，充能变为0.6倍
		if !cfg[i].Active {
			recharge = recharge * 0.6
		}
		cfg[i].Energy += recharge
		if cfg[i].Energy > cfg[i].EnergyMax {
			cfg[i].Energy = cfg[i].EnergyMax
		}
	}
}
