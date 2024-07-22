package attribute

type character_attribute int
type Attri struct {
	Max_hp                 float64
	Atk                    float64
	Def                    float64
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
}

const (
	MAX_HP character_attribute = iota
	ATK
	DEF
	ELEMENTAL_MASTERY
	CRIT_RATE
	CRIT_DMG
	HEALING_BONUS
	INCOMING_HEALING_BONUS
	ENERGY_RECHARGE
	CD_REDUCTION
	SHIELD_STRENGTH
	PYRO_DMG
	PYRO_RES
	HYDRO_DMG
	HYDRO_RES
	DENDRO_DMG
	DENDRO_RES
	ELECTRO_DMG
	ELECTOR_RES
	ANEMO_DMG
	ANEMO_RES
	CRYO_DMG
	CRYO_RES
	GEO_DMG
	GEO_RES
	PHISYCAL_DMG
	PHISYCAL_RES
)
