package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/somen440/graphql-sample/graph/generated"
	"github.com/somen440/graphql-sample/graph/model"
	"github.com/somen440/graphql-sample/models"
)

func (r *queryResolver) Liteners(ctx context.Context) ([]*model.Listener, error) {
	db, err := sqlx.Connect("mysql", "root:root@tcp(localhost:30306)/test_database?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	listeners, err := models.Listeners().All(ctx, db)
	if err != nil {
		return nil, err
	}

	followedChannels, err := models.FollowedChannels().All(ctx, db)
	if err != nil {
		return nil, err
	}

	channels, err := models.Channels().All(ctx, db)
	if err != nil {
		return nil, err
	}

	results := make([]*model.Listener, len(listeners))
	for i, listener := range listeners {
		followed := []*model.Channel{}
		for _, followedChannel := range followedChannels {
			if listener.ListenerID != followedChannel.ListenerID {
				continue
			}
			for _, channel := range channels {
				if followedChannel.ChannelID != channel.ChannelID {
					continue
				}
				followed = append(followed, &model.Channel{
					ID:        channel.ChannelID,
					Name:      channel.Name,
					CreatedAt: channel.CreatedAt,
					UpdatedAt: channel.UpdatedAt,
				})
			}
		}

		results[i] = &model.Listener{
			ID:        listener.ListenerID,
			Name:      listener.Name,
			Followed:  followed,
			CreatedAt: listener.CreatedAt,
			UpdatedAt: listener.UpdatedAt,
		}
	}

	return results, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
