package agent

import (
	"context"

	"github.com/kasaderos/notification/internal/model"
)

// AgentRepository defines the interface for agent data operations
type AgentRepository interface {
	Create(ctx context.Context, agent *model.Agent) error
	GetByID(ctx context.Context, id string) (*model.Agent, error)
	GetByUserID(ctx context.Context, userID string) (*model.Agent, error)
	Update(ctx context.Context, agent *model.Agent) error
	Delete(ctx context.Context, id string) error
	Fetch(ctx context.Context, limit, offset int) ([]*model.Agent, error)
}
