package pkg

import "strings"

/*
 乡镇一般区域			= 	乡镇
 乡镇核心区域  		= 	乡镇
 农村及其他 		 	= 	农村
 县城城区一般区域		=	县城
 县城城区核心区域		=	县城
 地级城市一般城区		=	城区
 地级城市主城区 		=	城区
*/

const (
	City      = "城区"
	County    = "县城"
	Town      = "乡镇"
	Village   = "农村"
	HighSpeed = "高速"
)

// 行政区域等级
//const (
//	LevelCity    = 10
//	LevelCounty  = 9
//	LevelTown    = 8
//	LevelVillage = 7
//)
//
//var AreaLevelMap = map[string]int{
//	City:    LevelCity,
//	County:  LevelCounty,
//	Town:    LevelTown,
//	Village: LevelVillage,
//}

// AreaLevelMap  （城区:12,县城:10,乡镇:8,农村:6） 场景分
var AreaLevelMap = map[string]int{
	City:    Source12,
	County:  Source10,
	Town:    Source8,
	Village: Source6,
}

func AreaMapping(area string) string {
	switch {
	case strings.Contains(area, City):
		return City
	case strings.Contains(area, County):
		return County
	case strings.Contains(area, Town):
		return Town
	case strings.Contains(area, Village):
		return Village
	}
	return ""
}
