package resolvers

import (
	"context"

	"github.com/somen440/graphql-sample/generated"
	"github.com/somen440/graphql-sample/models"
)

type channelResolver struct{ *Resolver }

var _ generated.ChannelResolver = (*channelResolver)(nil)

func (r *channelResolver) ID(ctx context.Context, obj *models.Channel) (string, error) {
	return obj.ChannelID, nil
}
