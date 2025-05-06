package pkg

const (
	Nvwa       = "nvwa"
	NvwaSource = "女娲系统"
	Plan       = "plan"
	PlanSource = "规划站"
	Dark       = "dark"
)

const (
	Source80 = 80
	Source70 = 70
	Source40 = 40
	Source30 = 30
	Source28 = 28
	Source25 = 25
	Source21 = 21
	Source20 = 20
	Source18 = 18
	Source16 = 16
	Source15 = 15
	Source12 = 12
	Source10 = 10
	Source8  = 8
	Source6  = 6
	Source5  = 5
	Source0  = 0
)

const (
	DefaultStatus = iota
	MergedStatus
)

// IsClose
var (
	CloseNot  = 0
	CloseDown = 1
)

/*
select create_by,count(1) from (
select planning_id,create_by,count(1) from plan_station group by planning_id,create_by having count(1) > 2
)tmp group by create_by;
*/

// 分公司：枚举值：低满用户、重复投诉、移动申告、栅格投诉等  重要投诉
// 黑点来源
const (
	LowSatisfactionUserConst = "低满用户投诉"
	RepeatedComplaintConst   = "重复投诉"
	GridComplaintConst       = "栅格投诉"
	VipComplaintConst        = "VIP投诉"
	MovingDeclarationConst   = "移动申告"
	MajorComplaintConst      = "重要投诉"
)

var DarkSpotSourceSlice = []string{
	VipComplaintConst,
	RepeatedComplaintConst,
	LowSatisfactionUserConst,
	GridComplaintConst,
}

// BaseStationNotOpen 质检未通过原因
var (
	BaseStationNotOpen         = "方案中基站未开通"
	FailureOfQualityInspection = "质检不通过"
)
