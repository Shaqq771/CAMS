package feature

import (
	"backend-nabati/domain/rule/model"
	repository "backend-nabati/domain/rule/repository"
	"context"
)

type RuleFeature interface {
	GetListOfRuleFeature(ctx context.Context) (response model.RuleListNoFilter, err error)
	GetRuleFeature(ctx context.Context, id string) (response model.RuleListNoFilter, err error)
	// GetRequestFilterFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (requestList model.RequestLists, err error)
}

type ruleFeature struct {
	ruleRepo repository.RuleRepository
}

func NewRuleFeature(ruleRepo repository.RuleRepository) RuleFeature {
	return &ruleFeature{
		ruleRepo: ruleRepo,
	}
}