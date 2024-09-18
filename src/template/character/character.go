package character

import "dps/src/core/attributes"

type Character struct {
	Name                   string
	Hp                     float64
	Atk                    float64
	Def                    float64
	Hpp                    float64
	Atkp                   float64
	Defp                   float64
	Elemental_mastery      int
	Crit_rate              float64
	Crit_dmg               float64
	Healing_bonus          float64
	Incoming_healing_bonus float64
	Energy_recharge        float64
	Cd_reduction           float64
	Shield_strength        float64
	Pyro_dmg               float64
	Pyro_res               float64
	Hydro_dmg              float64
	Hydro_res              float64
	Dendro_dmg             float64
	Dendro_res             float64
	Electro_dmg            float64
	Elector_res            float64
	Anemo_dmg              float64
	Anemo_res              float64
	Cryo_dmg               float64
	Cryo_res               float64
	Geo_dmg                float64
	Geo_res                float64
	Phisycal_dmg           float64
	Phisycal_res           float64

	NormalCounter int
}

type EnergyInfo struct {
	Count           int
	ParticleElement attributes.Element
	RechargeType    string
}
