package dataloaders

import (
	"context"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/somen440/graphql-sample/generated"
	"github.com/somen440/graphql-sample/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func ChannelSlice(db *sqlx.DB, w http.ResponseWriter, r *http.Request, next http.Handler) {
	loader := generated.NewChannelSliceLoader(generated.ChannelSliceLoaderConfig{
		MaxBatch: 100,
		Wait:     1 * time.Millisecond,
		Fetch: func(keys []string) ([][]*models.Channel, []error) {
			convertedKeys := make([]interface{}, len(keys))
			for index, key := range keys {
				convertedKeys[index] = key
			}
			followedChannels, err := models.FollowedChannels(
				qm.Select("distinct channel_id, listener_id"),
				qm.WhereIn("listener_id in ?", convertedKeys...),
			).All(r.Context(), db)

			if err != nil {
				return [][]*models.Channel{}, []error{err}
			}

			channelIDs := make([]interface{}, len(followedChannels))
			for i, fc := range followedChannels {
				channelIDs[i] = fc.ChannelID
			}

			channels, err := models.Channels(
				qm.WhereIn("channel_id in ?", channelIDs...),
			).All(r.Context(), db)
			if err != nil {
				return [][]*models.Channel{}, []error{err}
			}

			resultSet := make([][]*models.Channel, len(keys))
			for i, key := range keys {
				for _, fc := range followedChannels {
					if key != fc.ListenerID {
						continue
					}
					for _, channel := range channels {
						if fc.ChannelID != channel.ChannelID {
							continue
						}
						resultSet[i] = append(resultSet[i], channel)
					}
				}
			}

			return resultSet, nil
		},
	})

	ctx := context.WithValue(r.Context(), ChannelSliceLoader, loader)
	next.ServeHTTP(w, r.WithContext(ctx))
}
