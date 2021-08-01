package endpoint

import "context"

type IService interface {
	Create(ctx context.Context, name, gender string, birthYear int, avatar string) (uint, error)
	Verify(ctx context.Context, id uint) error
}
