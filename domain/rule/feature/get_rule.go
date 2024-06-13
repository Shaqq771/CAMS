package feature

import (
	"backend-nabati/domain/rule/constant"
	"backend-nabati/domain/rule/model"
	Error "backend-nabati/domain/shared/error"
	"context"
	"strconv"
)

func (rf ruleFeature) GetRuleFeature(ctx context.Context, id string) (response model.RuleListNoFilter, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	result, err := rf.ruleRepo.GetRuleByIdRepository(ctx, idInt)
	if err != nil {
		return
	}

	var ruleIds []int
	for _, approver := range result {
		ruleIds = append(ruleIds, approver.Id)
	}

	response = model.RuleListNoFilter{
		Rule: result,
	}
	return
}
