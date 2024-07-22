package attribute

type artifact_attribute int

const (
	Hp_percent artifact_attribute = iota
	Hp
	Atk_percent
	Atk
	Def_percent
	Def
	Elemental_mastery
	Crit_rate
	Crit_damage
	Healing_bonus
	Incoming_healing_bonus
	Energy_recharge
	Cd_reduction
	Shield_strength
	Pyro_dmg
	Pyro_res
	Hydro_dmg
	Hydro_res
	Dendro_dmg
	Dendro_res
	Electro_dmg
	Elector_res
	Anemo_dmg
	Anemo_res
	Cryo_dmg
	Cryo_res
	Geo_dmg
	Geo_res
	Phisycal_dmg
	Phisycal_res
	End_artifact_attribute
)
