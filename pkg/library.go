package pkg

// 库类型
const (
	ObservationLibrary = "观察库"
	InductiveLibrary   = "纳管库"
	TraceLibrary       = "跟踪库"
	ARCHIVISTLibrary   = "归档库"
)

var LibraryLevel = map[string]int{
	TraceLibrary:       1,
	ObservationLibrary: 2,
	InductiveLibrary:   3,
	ARCHIVISTLibrary:   4,
}
