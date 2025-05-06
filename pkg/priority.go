package pkg

// 枚举值：高、中、低 优先级（高：三维度中任一维度为满分；三维度综合评分>=70。中：40<=三维度综合评分<70。低：三维度综合评分<40。
const (
	PriorityHigh   = "高"
	PriorityMedium = "中"
	PriorityLow    = "低"
)

// PriorityScore 优先级 高优先级:18,中优先级:15,低优先级:10）
var PriorityScore = map[string]int{
	PriorityHigh:   Source18,
	PriorityMedium: Source15,
	PriorityLow:    Source10,
}
