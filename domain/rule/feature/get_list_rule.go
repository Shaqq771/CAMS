package feature

import (
	"backend-nabati/domain/rule/model"
	"context"
)

func (rf ruleFeature) GetListOfRuleFeature(ctx context.Context) (response model.RuleListNoFilter, err error) {

	listRule, err := rf.ruleRepo.GetListOfRuleRepository(ctx)
	if err != nil {
		return
	}
	response = model.RuleListNoFilter{
		Rule: listRule,
	}

	return
}