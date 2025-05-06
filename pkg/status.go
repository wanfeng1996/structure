package pkg

const (
	ScoreEntry   = "评分入库"
	Design       = "方案设计"
	Construction = "施工中"
	Result       = "结果质检"
	Archive      = "归档"
)

type LevelStatus int

const (
	ArchiveLevel LevelStatus = iota
	ResultLevel
	ConstructionLevel
	DesignLevel
	ScoreEntryLevel
)

//对流程排序用

var StatusLevel = map[string]int{
	Archive:      int(ArchiveLevel),
	Result:       int(ResultLevel),
	Construction: int(ConstructionLevel),
	Design:       int(DesignLevel),
	ScoreEntry:   int(ScoreEntryLevel),
}
