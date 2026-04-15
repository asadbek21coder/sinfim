package schoolrequest

import "context"

type Repo interface {
	Create(ctx context.Context, request *SchoolRequest) (*SchoolRequest, error)
	FindOpenDuplicate(ctx context.Context, phoneNumber string, schoolName string) (*SchoolRequest, error)
	List(ctx context.Context, filter Filter) ([]SchoolRequest, error)
	UpdateStatus(ctx context.Context, id string, status string) (*SchoolRequest, error)
}
