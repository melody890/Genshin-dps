package attributes

type Element int
type ElementType int

const (
	Electro Element = iota
	Pyro
	Cryo
	Hydro
	Dendro
	Quicken // or overdose
	Frozen
	Anemo
	Geo
	NoElement
	Physical
	UnknownElement
	EndEleType
)

const (
	Weak        ElementType = iota //弱元素
	Medium                         //中元素
	Strong                         //强元素
	SuperStrong                    //超强元素
	Clear                          //这个是没有元素的时候
)

func EleToDmgP(e Element) Stat {
	switch e {
	case Anemo:
		return AnemoP
	case Cryo:
		return CryoP
	case Electro:
		return ElectroP
	case Geo:
		return GeoP
	case Hydro:
		return HydroP
	case Pyro:
		return PyroP
	case Dendro:
		return DendroP
	case Physical:
		return PhyP
	}
	return -1
}
