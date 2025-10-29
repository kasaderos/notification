package notification

type Service struct {
	repo     Repository
	ruleRepo RuleRepository
}

func NewService(repo Repository, ruleRepo RuleRepository) *Service {
	return &Service{repo: repo, ruleRepo: ruleRepo}
}
