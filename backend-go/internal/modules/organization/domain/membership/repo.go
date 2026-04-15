package membership

import "context"

type Repo interface {
	Create(ctx context.Context, membership *Membership) (*Membership, error)
	Exists(ctx context.Context, userID string, organizationID string, role string) (bool, error)
	ListByUser(ctx context.Context, userID string) ([]Membership, error)
	ListByOrganization(ctx context.Context, organizationID string) ([]Membership, error)
}
