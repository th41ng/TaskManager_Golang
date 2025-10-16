package commonrepo

import "context"

type CRUDRepo[T any, C any] struct {
	CreateFn func(ctx context.Context, client C, args ...any) (*T, error)
	GetFn    func(ctx context.Context, client C, id int) (*T, error)
	UpdateFn func(ctx context.Context, client C, id int, args ...any) (*T, error)
	DeleteFn func(ctx context.Context, client C, id int) error
	ListFn   func(ctx context.Context, client C) ([]*T, error)
	Client   C
}

func (r *CRUDRepo[T, C]) Create(ctx context.Context, args ...any) (*T, error) {
	return r.CreateFn(ctx, r.Client, args...)
}
func (r *CRUDRepo[T, C]) GetByID(ctx context.Context, id int) (*T, error) {
	return r.GetFn(ctx, r.Client, id)
}
func (r *CRUDRepo[T, C]) Update(ctx context.Context, id int, args ...any) (*T, error) {
	return r.UpdateFn(ctx, r.Client, id, args...)
}
func (r *CRUDRepo[T, C]) Delete(ctx context.Context, id int) error {
	return r.DeleteFn(ctx, r.Client, id)
}
func (r *CRUDRepo[T, C]) List(ctx context.Context) ([]*T, error) { return r.ListFn(ctx, r.Client) }
