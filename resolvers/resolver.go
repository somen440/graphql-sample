package resolvers

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/somen440/graphql-sample/generated"
)

var (
	errNotImplemented = errors.New("not implemented")
)

type Resolver struct {
	Db *sqlx.DB
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Listener() generated.ListenerResolver {
	return &listenerResolver{r}
}
func (r *Resolver) Channel() generated.ChannelResolver {
	return &channelResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
