package pkg

const (
	MacroStation        = "宏站"
	VentricularDivision = "室分"
)

// 是否排序用
const (
	LevelVentricularDivision = iota
	LevelMacroStation
)

var StationLevel = map[string]int{
	MacroStation:        LevelMacroStation,
	VentricularDivision: LevelVentricularDivision,
}
