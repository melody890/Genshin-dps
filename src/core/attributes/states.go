package attributes

type Stat int

// stat types
const (
	NoStat Stat = iota
	DEF
	DEFP
	HP
	HPP
	ATK
	ATKP
	ER
	EM
	CR
	CD
	Heal
	PyroP
	HydroP
	CryoP
	ElectroP
	AnemoP
	GeoP
	DendroP
	PhyP
	AtkSpd
	DmgP
	DelimBaseStat
	BaseHP
	BaseATK
	BaseDEF
	MaxEnergy
	// delim
	EndStatType
)
