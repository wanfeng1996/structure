package pkg

import "errors"

var (
	QualityErr                = errors.New("质检错误")
	CellNameIsNoll            = errors.New("小区名为空")
	ConditionNotMetErr        = errors.New("条件未通过")
	QualityCellNameIsNollErr  = errors.New(QualityErr.Error() + " : " + CellNameIsNoll.Error())
	QualityConditionNotMetErr = errors.New(QualityErr.Error() + ":" + ConditionNotMetErr.Error())
	//ProposalTimeIsNull        = errors.New(QualityConditionNotMetErr.Error() + ": 制定方案时间为空!")
)
