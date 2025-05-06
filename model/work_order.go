package model

import (
	"github.com/wanfeng1996/structure/constant"
	"slices"
	"strings"
	"time"
)

type Query struct {
	Ids           []int
	OwningLibrary []string
}

type BaseCity struct {
	City   string `json:"city" gorm:"column:city"`     // 地市
	County string `json:"county" gorm:"column:county"` // 区县
}

type BaseParam struct {
	City            string   `json:"city" gorm:"column:city"`                         // 地市
	County          string   `json:"county" gorm:"column:county"`                     // 区县
	Priority        string   `json:"priority" gorm:"column:priority"`                 // 优先级
	DarkSpotName    string   `json:"dark_spot_name" gorm:"column:dark_spot_name"`     // 黑点名称
	DarkSpotSource  string   `json:"dark_spot_source" gorm:"column:dark_spot_source"` // 黑点来源
	MainCell        *string  `json:"main_cell" gorm:"column:main_cell"`               // 投诉主覆盖小区
	CenterLongitude *float64 `json:"center_longitude" gorm:"column:center_longitude"` // 中心经度
	CenterLatitude  *float64 `json:"center_latitude" gorm:"column:center_latitude"`   // 中心维度
	BoundaryPoints  *string  `json:"boundary_points" gorm:"column:boundary_points"`   // 边界经纬度
}

// WorkOrder 工单跟踪表
type WorkOrder struct {
	BaseParam
	ID                 int       `json:"id" gorm:"column:id;primaryKey"`                          // 序号
	Status             int       `json:"status" gorm:"column:status"`                             // 状态 0 正常状态 1 已合并
	CreateTime         time.Time `json:"create_time" gorm:"column:create_time;type:timestamptz"`  // 下库时间
	OwningLibrary      string    `json:"owning_library" gorm:"column:owning_library"`             // 所属库
	ComplaintLocation  *string   `json:"complaint_location" gorm:"column:complaint_location"`     // 投诉地点
	ComplaintPhone     *string   `json:"complaint_phone" gorm:"column:complaint_phone"`           // 投诉号码
	ServiceRequestType *string   `json:"service_request_type" gorm:"column:service_request_type"` // 服务请求分类
	ComplaintTime      *string   `json:"complaint_time" gorm:"column:complaint_time"`             // 投诉时间
	CollectionSource   *string   `json:"collection_source" gorm:"column:collection_source"`       // 收集来源
	SourceType         *string   `json:"source_type" gorm:"column:source_type"`                   // 来源分类
	//BoundaryPointsOrb       orb.Polygon `json:"-" gorm:"-"`                                                            // 边界经纬度
	Area                    *float64 `json:"area" gorm:"column:area"`                                               // 面积（平方米）
	Scene                   *string  `json:"scene" gorm:"column:scene"`                                             // 场景
	AdminArea               *string  `json:"admin_area" gorm:"column:admin_area"`                                   // 行政区域
	IsPartyGovernmentArmy   *string  `json:"is_party_government_army" gorm:"column:is_party_government_army"`       // 是否党政军
	CompetitorInfo          *string  `json:"competitor_info" gorm:"column:competitor_info"`                         // 是否属于人有我无或人有我弱
	IsUndergroundOrLiftWell *string  `json:"is_underground_or_lift_well" gorm:"column:is_underground_or_lift_well"` // 是否包含地下停车场或电梯间
	SiteType                *string  `json:"site_type" gorm:"column:site_type"`                                     // 站点类型
	ThisCoverageArea        *int     `json:"this_coverage_area" gorm:"column:this_coverage_area"`                   // 本次覆盖面积
	CorrelationComplaintCT  *int     `json:"correlation_complaint_ct" gorm:"column:correlation_complaint_ct"`       // 关联投诉次数
	ComplaintID             *string  `json:"complaint_id" gorm:"column:complaint_id"`                               // 投诉工单号
	ResponsiblePerson       *string  `json:"responsible_person" gorm:"column:responsible_person"`                   // 分公司上报数据真实性责任人
	OrdinaryComplaint       *int     `json:"ordinary_complaint" gorm:"column:ordinary_complaint"`                   // X普通投诉
	RepeatComplaint         *int     `json:"repeat_complaint" gorm:"column:repeat_complaint"`                       // X重复投诉
	VIPComplaint            *int     `json:"vip_complaint" gorm:"column:vip_complaint"`                             // XVIP投诉
	LowSatisfyComplaint     *int     `json:"low_satisfy_complaint" gorm:"column:low_satisfy_complaint"`             // X低满用户投诉
	GridComplaint           *int     `json:"grid_complaint" gorm:"column:grid_complaint"`                           // XTOP栅格投诉
	//1）当“黑点来源”包含(“VIP投诉”、“重复投诉”、“低满用户投诉”、“TOP栅格投诉”) 其中之一时，赋值 30，优先级 赋值“高”
	//2）当以下各列中其中之一（X重复投诉、XVIP投诉、X低满用户投诉、XTOP栅格投诉）不为空时，赋值 30，优先级 赋值“高”
	//3）当“普通投诉”不为空 且 X重复投诉、XVIP投诉、X低满用户投诉、XTOP栅格投诉 都为空时，赋值 5
	//4）以上条件都不符则为 0"
	ComplaintTypeScore int    `json:"complaint_type_score" gorm:"column:complaint_type_score"` // X投诉类型得分
	IsHighSpeed        string `json:"is_high_speed" gorm:"column:is_high_speed"`               // Y是否高速场景
	CompetitorCovered  string `json:"competitor_covered" gorm:"column:competitor_covered"`     // Y竞争对手是否已覆盖
	SceneLevel         string `json:"scene_level" gorm:"column:scene_level"`                   // Y景区等级
	/*
		当“Y竞争对手是否已覆盖”=“否”
		且 “Y是否高速场景”=“否”
		且 “Y景区等级”不为空 时，
		用“Y景区等级”数据匹配
		（1A:18,2A:21,3A:25,4A:28,5A:30）回填分值，
		例如 “Y景区等级”=2A 则回填 21。以上不符则回填 0"
	*/
	SceneLevelScore int     `json:"scene_level_score" gorm:"column:scene_level_score"`   // Y景区等级得分
	YAdminAreaScene *string `json:"y_admin_area_scene" gorm:"column:y_admin_area_scene"` // Y行政区划分场景
	//YAdminAreaSceneList []string `json:"-" gorm:"-"`                                          // Y行政区划分场景
	/*
		当“Y竞争对手是否已覆盖”=“否”
		且 “Y是否高速场景”=“否”
		且 “Y景区等级”为空
		且 “Y行政区划分场景”不为空，
		用“Y行政区划分场景”数据匹配（城区:12,县城:10,乡镇:8,农村:6）回填分值，
		例如 “Y行政区划分场景”=“县城” 则回填 10。以上不符则回填 0"
	*/
	YAdminAreaSceneScore int     `json:"y_admin_area_scene_score" gorm:"column:y_admin_area_scene_score"` // Y行政区划分场景得分
	OverlayScene         *string `json:"overlay_scene" gorm:"column:overlay_scene"`                       // Y覆盖场景
	ScenarioPriority     string  `json:"scenario_priority" gorm:"column:scenario_priority"`               // Y覆盖场景优先级
	/*
		当“Y竞争对手是否已覆盖”=“否”
		且 “Y是否高速场景”=“否”
		且 “Y景区等级”为空
		且 “Y覆盖场景优先级”不为空，
		用“Y覆盖场景优先级”数据匹配
		（高优先级:18,中优先级:15,低优先级:10）回填分值，
		例如 “Y覆盖场景优先级”=“中优先级” 则回填 15。以上不符则回填 0"
	*/
	ScenarioPriorityScore int `json:"scenario_priority_score" gorm:"column:scenario_priority_score"` // Y覆盖场景优先级得分
	/*
		1）当“Y竞争对手是否已覆盖”=“是” 或 “Y是否高速场景”=“是”时，赋值30，优先级 赋值 “高”
		2）当“Y竞争对手是否已覆盖”=“否” 且 “Y是否高速场景”=“否” 且 “Y景区等级”不为空 时，则为 “Y景区等级得分”
		3）当“Y竞争对手是否已覆盖”=“否” 且 “Y是否高速场景”=“否” 且 “Y景区等级”为空 则为“Y行政区划分场景得分”+“Y覆盖场景优先级得分”
		4）以上条件都不符则为 0"
	*/
	SceneTypeScore  int     `json:"scene_type_score" gorm:"column:scene_type_score"`     // Y场景类型得分
	ZAdminAreaScene *string `json:"z_admin_area_scene" gorm:"column:z_admin_area_scene"` // Z行政区划分场景
	Flow            *int    `json:"flow" gorm:"column:flow"`                             // Z流量GB
	/*
			1）当“Z行政区划分场景”in (""城区"",""县城"") 时，
				根据“Z流量GB”的给予评分（>=40:40,15<= 且 <40:28,<15:16）。
		例如 “Z流量GB”=16 则回填 28
			2）以上都不符则为 0"
	*/
	FlowScore    int `json:"flow_score" gorm:"column:flow_score"`       // Z流量得分
	HouseholdsCT int `json:"households_ct" gorm:"column:households_ct"` // Z居住户数
	/*
		1）当“Z行政区划分场景”=""乡镇"" 时，
			根据“Z居住户数”的给予评分（>=80:40,40<= 且 <80:28,<40:16）。
			例如 “Z居住户数”=46 则回填 28
		2）当“Z行政区划分场景”=""农村"" 时，
			根据“Z居住户数”的给予评分（>=40:40,20<= 且 <40:28,<20:16）。
			例如 “Z居住户数”=25 则回填 28
		3）以上都不符则为 0"
	*/
	HouseholdsCTScore int `json:"households_ct_score" gorm:"column:households_ct_score"` // Z居住户数得分
	/*
		Z流量得分+Z居住户数得分
	*/
	ValueScore int `json:"value_score" gorm:"column:value_score"` // Z价值得分
	/*
		总分 = X投诉类型得分+Y场景类型得分+Z价值得分
	*/
	TotalScore int `json:"total_score" gorm:"column:total_score"` // 总分
	/*
		"评定方法：
		1）当 总分>=70，则为'高'
		2）当 X投诉类型得分=30 或 Y场景类型得分=30 或 Z居住户数得分=40，则为'高'
		3）当 40<=总分<70，则为'中'
		4）以上都不符为'低'"
	*/

	SchemeType              *string    `json:"scheme_type" gorm:"column:scheme_type"`                               // 解决方案类型
	Scheme                  *string    `json:"scheme" gorm:"column:scheme"`                                         // 解决方案
	SchemeDetail            *string    `json:"scheme_detail" gorm:"column:scheme_detail"`                           // 解决方案详细描述
	IsToBeSolved            *string    `json:"is_to_be_solved" gorm:"column:is_to_be_solved"`                       // 是否计划解决
	SolvedPlanTime          *string    `json:"solved_plan_time" gorm:"column:solved_plan_time"`                     // 预计解决时间
	Project                 *string    `json:"project" gorm:"column:project"`                                       // 归属项目
	PlanID                  *string    `json:"plan_id" gorm:"column:plan_id"`                                       // 规划ID  plan：需求编号  	//端到端规划表
	BaseName                *string    `json:"base_name" gorm:"column:base_name"`                                   // 基站名 置为空
	CellName                *string    `json:"cell_name" gorm:"column:cell_name"`                                   // 小区名 置为空
	CellLng                 *string    `json:"cell_lng" gorm:"column:cell_lng"`                                     // 小区经度 plan:
	CellLat                 *string    `json:"cell_lat" gorm:"column:cell_lat"`                                     // 小区纬度 plan:
	WarehousingTime         *string    `json:"warehousing_time" gorm:"column:warehousing_time"`                     // 需求入库时间 plan: 需求审核时间
	PlanningReviewTime      *string    `json:"planning_review_time" gorm:"column:planning_review_time"`             // 初审完成时间
	FinalJudgmentTime       *string    `json:"final_judgment_time" gorm:"column:final_judgment_time"`               // 终审完成时间 plan: 设计审核时间
	ConstructionTime        *string    `json:"construction_time" gorm:"column:construction_time"`                   // 施工完成时间 plan: 默认为空，当“所处阶段”='施工阶段-施工验收' 时回填当前系统时间
	DataLoadTime            *string    `json:"data_load_time" gorm:"column:data_load_time"`                         // 数据加载完成时间 plan: 默认为空，当“所处阶段”='入网阶段-数据加载' 时回填当前系统时间
	SingleReviewTime        *string    `json:"single_review_time" gorm:"column:single_review_time"`                 // 单验完成时间
	AccessThroughTime       *string    `json:"access_through_time" gorm:"column:access_through_time"`               // 入网评估通过时间 plan: 默认为空，当“所处阶段”='入网阶段-入网评估' 时回填当前系统时间
	IsClosed                *int       `json:"is_closed" gorm:"column:is_closed"`                                   // 是否闭环
	Coverage                *float64   `json:"coverage" gorm:"column:coverage"`                                     // 质检 覆盖率 todo 修改为MDT 后覆盖率
	AvgRSRP                 *float64   `json:"avg_rsrp" gorm:"column:avg_rsrp"`                                     // 质检平均RSRP
	NotClosedReason         *string    `json:"not_closed_reason" gorm:"column:not_closed_reason"`                   // 质检测未通过原因
	Remark                  *string    `json:"remark" gorm:"column:remark"`                                         // 备注
	SendOrderTime           *time.Time `json:"send_order_time" gorm:"column:send_order_time"`                       // 派单时间
	ProposalTime            *time.Time `json:"proposal_time" gorm:"column:proposal_time"`                           // 制定方案时间
	ProposalUser            *string    `json:"proposal_user" gorm:"column:proposal_user"`                           // 方案制定人
	QualityTime             *string    `json:"quality_time" gorm:"column:quality_time"`                             // 质检时间
	ProposalDuration        *float64   `json:"proposal_duration" gorm:"column:proposal_duration"`                   // 制定方案流程时长(小时)
	WaveHousingDuration     *float64   `json:"wave_housing_duration" gorm:"column:wave_housing_duration"`           // 需求入库流程时长(天)
	PlanningReviewDuration  *float64   `json:"planning_review_duration" gorm:"column:planning_review_duration"`     // 初审完成时长(天)
	FinalJudgmentDuration   *float64   `json:"final_judgment_duration" gorm:"column:final_judgment_duration"`       // 终审流程时长(天)
	ConstructionDuration    *float64   `json:"construction_duration" gorm:"column:construction_duration"`           // 施工流程时长(天)
	DataLoadDuration        *float64   `json:"data_load_duration" gorm:"column:data_load_duration"`                 // 数据加载流程时长(天)
	SingleReviewDuration    *float64   `json:"single_review_duration" gorm:"column:single_review_duration"`         // 单验完成时长(天)
	AccessThroughDuration   *float64   `json:"access_through_duration" gorm:"column:access_through_duration"`       // 入网评估流程时长(天)
	ProcessDuration         *float64   `json:"process_duration" gorm:"column:process_duration"`                     // 流程总时长
	HistoryProposalType     *string    `json:"history_proposal_type" gorm:"column:history_proposal_type"`           // 历史解决方案类型
	HistoryProposal         *string    `json:"history_proposal" gorm:"column:history_proposal"`                     // 历史解决方案
	HistoryProposalDetail   *string    `json:"history_proposal_detail" gorm:"column:history_proposal_detail"`       // 历史解决方案详细描述
	MDTBeforeSamplingNumber *int       `json:"mdt_before_sampling_number" gorm:"column:mdt_before_sampling_number"` // MDT栅格采样点数(前)
	MDTBeforeCoverage       *float64   `json:"mdt_before_coverage" gorm:"column:mdt_before_coverage"`               // MDT栅格覆盖率(前)
	MDTAfterSamplingNumber  *int       `json:"mdt_after_sampling_number" gorm:"column:mdt_after_sampling_number"`   // MDT栅格采样点数(后)
	MDTAfterCoverage        *float64   `json:"mdt_after_coverage" gorm:"column:mdt_after_coverage"`                 // MDT栅格覆盖率(后)
	ReportLink              *string    `json:"report_link" gorm:"column:reportlink"`                                // 报告链接
	Reporter                *string    `json:"reporter" gorm:"column:reporter"`                                     // 报告上传人
	ReportTime              *time.Time `json:"report_time" gorm:"column:reporttime"`                                // 报告上传时间
	DueStatus               *string    `json:"due_status" gorm:"column:due_status"`                                 // 处理进度
	NetType                 *string    `json:"net_type" gorm:"column:nettype"`                                      // 网络制式 todo 待处理
	ProcessLimited          *float64   `json:"process_limited" gorm:"column:process_limited"`                       // 处理时限 todo
	ElapsedTime             *float64   `json:"elapsed_time" gorm:"column:elapsed_time"`                             // 当前耗时 todo
	IsTimeOut               *string    `json:"is_timeout" gorm:"column:is_timeout"`                                 // 是否超时 todo
}

func (*WorkOrder) TableName() string {
	return "work_order"
}

type WorkOrders []*WorkOrder

func (w *WorkOrders) Merge() *WorkOrder {
	return &WorkOrder{
		BaseParam: BaseParam{
			City:            (*w)[0].City,
			County:          (*w)[0].County,
			Priority:        Default(NewPriority(StructField[string](w, "BaseParam.Priority").Index(), PriorityScore).Max()),
			DarkSpotName:    Default(w.StringMerge("BaseParam.DarkSpotName")),
			DarkSpotSource:  Default(w.StringMerge("BaseParam.DarkSpotSource")),
			MainCell:        w.StringMerge("BaseParam.MainCell"),
			CenterLongitude: (*w)[0].CenterLongitude,
			CenterLatitude:  (*w)[0].CenterLatitude,
			BoundaryPoints:  (*w)[0].BoundaryPoints,
		},
		OwningLibrary:           Default(NewPriority(StructField[string](w, "OwningLibrary").Index(), LibraryLevel).Max()),
		CreateTime:              Default(TimeSort(StructField[time.Time](w, "CreateTime").Index()).Min()),
		ComplaintLocation:       w.StringMerge("ComplaintLocation"),
		ComplaintPhone:          w.StringMerge("ComplaintPhone"),
		ServiceRequestType:      w.StringMerge("ServiceRequestType"),
		ComplaintTime:           w.StringMerge("ComplaintTime"),
		CollectionSource:        w.StringMerge("CollectionSource"),
		SourceType:              w.StringMerge("SourceType"),
		Area:                    Max(StructField[float64](w, "Area").Index()),
		Scene:                   w.StringMerge("Scene"),
		AdminArea:               NewPriority(StructField[string](w, "AdminArea").Index(), AreaLevelMap).Max(),
		IsPartyGovernmentArmy:   NewPriority(StructField[string](w, "IsPartyGovernmentArmy").Index(), IsLevel).Max(),
		CompetitorInfo:          NewPriority(StructField[string](w, "CompetitorInfo").Index(), IsLevel).Max(),
		IsUndergroundOrLiftWell: NewPriority(StructField[string](w, "IsUndergroundOrLiftWell").Index(), IsLevel).Max(),
		SiteType:                NewPriority(StructField[string](w, "SiteType").Index(), StationLevel).Max(),
		ThisCoverageArea:        Max(StructField[int](w, "ThisCoverageArea").Index()),
		CorrelationComplaintCT:  Max(StructField[int](w, "CorrelationComplaintCT").Index()),
		ComplaintID:             w.StringMerge("ComplaintID"),
		ResponsiblePerson:       w.StringMerge("ResponsiblePerson"),
		OrdinaryComplaint:       Max(StructField[int](w, "OrdinaryComplaint").Index()),
		RepeatComplaint:         Max(StructField[int](w, "RepeatComplaint").Index()),
		VIPComplaint:            Max(StructField[int](w, "VIPComplaint").Index()),
		LowSatisfyComplaint:     Max(StructField[int](w, "LowSatisfyComplaint").Index()),
		GridComplaint:           Max(StructField[int](w, "GridComplaint").Index()),
		ComplaintTypeScore:      Default(Max(StructField[int](w, "ComplaintTypeScore").Index())),
		IsHighSpeed:             Default(NewPriority(StructField[string](w, "IsHighSpeed").Index(), IsLevel).Max()),
		CompetitorCovered:       Default(NewPriority(StructField[string](w, "CompetitorCovered").Index(), IsLevel).Max()),
		SceneLevel:              Default(NewPriority(StructField[string](w, "SceneLevel").Index(), ScenicSpotLevel).Max()),
		SceneLevelScore:         Default(Max(StructField[int](w, "SceneLevelScore").Index())),
		YAdminAreaScene:         NewPriority(StructField[string](w, "YAdminAreaScene").Index(), AreaLevelMap).Max(),
		YAdminAreaSceneScore:    Default(Max(StructField[int](w, "YAdminAreaSceneScore").Index())),
		OverlayScene:            w.StringMerge("OverlayScene"),
		ScenarioPriority:        Default(NewPriority(StructField[string](w, "ScenarioPriority").Index(), PriorityScore).Max()),
		ScenarioPriorityScore:   Default(Max(StructField[int](w, "ScenarioPriorityScore").Index())),
		SceneTypeScore:          Default(Max(StructField[int](w, "SceneTypeScore").Index())),
		ZAdminAreaScene:         NewPriority(StructField[string](w, "ZAdminAreaScene").Index(), AreaLevelMap).Max(),
		Flow:                    Max(StructField[int](w, "Flow").Index()),
		FlowScore:               Default(Max(StructField[int](w, "FlowScore").Index())),
		HouseholdsCT:            Default(Max(StructField[int](w, "HouseholdsCT").Index())),
		HouseholdsCTScore:       Default(Max(StructField[int](w, "HouseholdsCTScore").Index())),
		ValueScore:              Default(Max(StructField[int](w, "ValueScore").Index())),
		TotalScore:              Default(Max(StructField[int](w, "TotalScore").Index())),
		SchemeType:              w.StringMerge("SchemeType"),
		Scheme:                  w.StringMerge("Scheme"),
		SchemeDetail:            w.StringMerge("SchemeDetail"),
		IsToBeSolved:            NewPriority(StructField[string](w, "IsToBeSolved").Index(), IsLevel).Max(),
		SolvedPlanTime:          w.StringMerge("SolvedPlanTime"),
		Project:                 w.StringMerge("Project"),
		PlanID:                  w.StringMerge("PlanID"),
		BaseName:                w.StringMerge("BaseName"),
		CellName:                w.StringMerge("CellName"),
		CellLng:                 w.StringMerge("CellLng"),
		CellLat:                 w.StringMerge("CellLat"),
		WarehousingTime:         w.StringMerge("WarehousingTime"),
		FinalJudgmentTime:       w.StringMerge("FinalJudgmentTime"),
		ConstructionTime:        w.StringMerge("ConstructionTime"),
		DataLoadTime:            w.StringMerge("DataLoadTime"),
		AccessThroughTime:       w.StringMerge("AccessThroughTime"),
		IsClosed:                Max(StructField[int](w, "IsClosed").Index()),
		Coverage:                Max(StructField[float64](w, "Coverage").Index()),
		AvgRSRP:                 Max(StructField[float64](w, "AvgRSRP").Index()),
		NotClosedReason:         w.StringMerge("NotClosedReason"),
		Remark:                  w.StringMerge("Remark"),
		SendOrderTime:           TimeSort(StructField[time.Time](w, "SendOrderTime").Index()).Min(),
		ProposalTime:            TimeSort(StructField[time.Time](w, "ProposalTime").Index()).Min(),
		ProposalUser:            w.StringMerge("ProposalUser"),
		QualityTime:             w.StringMerge("QualityTime"),                                                    // 质检时间
		ProposalDuration:        Max(StructField[float64](w, "ProposalDuration").Index()),                        // 制定方案流程时长(小时)
		WaveHousingDuration:     Max(StructField[float64](w, "WaveHousingDuration").Index()),                     // 需求入库流程时长(天)
		FinalJudgmentDuration:   Max(StructField[float64](w, "FinalJudgmentDuration").Index()),                   // 终审流程时长(天)
		ConstructionDuration:    Max(StructField[float64](w, "ConstructionDuration").Index()),                    // 施工流程时长(天)
		DataLoadDuration:        Max(StructField[float64](w, "DataLoadDuration").Index()),                        // 数据加载流程时长(天)
		AccessThroughDuration:   Max(StructField[float64](w, "AccessThroughDuration").Index()),                   // 入网评估流程时长(天)
		PlanningReviewDuration:  Max(StructField[float64](w, "PlanningReviewDuration").Index()),                  // 初审完成时长(天)
		SingleReviewDuration:    Max(StructField[float64](w, "SingleReviewDuration").Index()),                    // 单验完成时长(天)
		ProcessDuration:         Max(StructField[float64](w, "ProcessDuration").Index()),                         // 流程总时长(天)
		HistoryProposalType:     w.StringMerge("HistoryProposalType"),                                            //
		HistoryProposal:         w.StringMerge("HistoryProposal"),                                                //
		HistoryProposalDetail:   w.StringMerge("HistoryProposalDetail"),                                          //
		MDTBeforeSamplingNumber: Max(StructField[int](w, "MDTBeforeSamplingNumber").Index()),                     // MDT栅格采样点数(前)
		MDTBeforeCoverage:       Max(StructField[float64](w, "MDTBeforeCoverage").Index()),                       // MDT栅格覆盖率(前)
		MDTAfterSamplingNumber:  Max(StructField[int](w, "MDTAfterSamplingNumber").Index()),                      // MDT栅格采样点数(后)
		MDTAfterCoverage:        Max(StructField[float64](w, "MDTAfterCoverage").Index()),                        // MDT栅格覆盖率(后)
		ReportLink:              w.StringMerge("ReportLink"),                                                     // 报告链接
		Reporter:                w.StringMerge("Reporter"),                                                       // 报告上传人
		ReportTime:              TimeSort(StructField[time.Time](w, "ReportTime").Index()).Min(),                 // 报告上传时间
		DueStatus:               NewPriority(StructField[string](w, "DueStatus").Index(), StatusLevel).Min(), // 处理进度
		NetType:                 w.StringMerge("NetType"),                                                        // 网络制式
		ProcessLimited:          Max(StructField[float64](w, "ProcessLimited").Index()),                          // 处理时限
	}
}

func (w *WorkOrders) StringMerge(name string) *string {
	var duplicates []string
	wName := StructField[string](w, name).Index()
	for i := range wName {
		duplicates = append(duplicates, strings.Split(wName[i], constant.Slash)...)
	}
	// 使用 DeleteFunc 删除所有空字符串
	duplicates = slices.DeleteFunc(duplicates, func(s string) bool {
		return s == ""
	})

	return Address(strings.Join(RemoveDuplicates(duplicates), constant.Slash))
}
