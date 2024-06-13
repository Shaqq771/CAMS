package repository

import (
	"backend-nabati/domain/rule/model"
	"backend-nabati/infrastructure/database"
	"context"
)

type RuleRepository interface {
	GetListOfRuleRepository(ctx context.Context) (rules []model.Rule, err error)
	GetRuleByIdRepository(ctx context.Context, id int) (rule []model.Rule, err error)
	// GetRequestListsWithFiltersRepository(ctx context.Context, filter *shared_model.Filter, offset int) (requests []model.Request, err error)
	// GetTotalRequestWithFiltersRepository(ctx context.Context, filter *shared_model.Filter) (count int, err error)
	// GetTotalRequestWithConditionsRepository(ctx context.Context, conditions string) (count int, err error)
	// GetTotalRequestRepository(ctx context.Context) (count int, err error)
	// GetRequestListsRepository(ctx context.Context, limit, offset int, sortby, search string) (requests []model.Request, err error)
}

type ruleRepository struct {
	Database *database.Database
}

func NewRuleRepository(db *database.Database) RuleRepository {
	return &ruleRepository{
		Database: db,
	}
}