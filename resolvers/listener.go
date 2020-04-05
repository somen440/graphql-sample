package resolvers

import (
	"context"

	"github.com/somen440/graphql-sample/dataloaders"
	"github.com/somen440/graphql-sample/generated"
	"github.com/somen440/graphql-sample/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type listenerResolver struct{ *Resolver }

var _ generated.ListenerResolver = (*listenerResolver)(nil)

func (r *listenerResolver) ID(ctx context.Context, obj *models.Listener) (string, error) {
	return obj.ListenerID, nil
}

func (r *listenerResolver) Followed(ctx context.Context, obj *models.Listener) ([]*models.Channel, error) {
	return ctx.Value(dataloaders.ChannelSliceLoader).(*generated.ChannelSliceLoader).
		Load(obj.ListenerID)
}

func (r *mutationResolver) AddListener(ctx context.Context) (*models.Listener, error) {
	return nil, errNotImplemented
}

func (r *queryResolver) Listener(ctx context.Context, id string) (*models.Listener, error) {
	return models.Listeners(qm.Where("listener_id = ?", id)).One(ctx, r.Db)
}

func (r *queryResolver) Liteners(ctx context.Context) ([]*models.Listener, error) {
	return models.Listeners().All(ctx, r.Db)
}
