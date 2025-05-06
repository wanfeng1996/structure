package pkg

const (
	Is = "是"
	No = "否"
)

// 是否排序用
const (
	LevelNo = iota
	LevelIs
)

var IsLevel = map[string]int{
	Is: LevelIs,
	No: LevelNo,
}
