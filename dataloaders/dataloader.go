package dataloaders

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

const ChannelSliceLoader = "channelSliceLoader"

func setLoader(db *sqlx.DB, dataloader func(db *sqlx.DB, w http.ResponseWriter, r *http.Request, next http.Handler)) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dataloader(db, w, r, next)
		})
	}
}

func NewMiddleware(session *sqlx.DB) []func(handler http.Handler) http.Handler {
	return []func(handler http.Handler) http.Handler{
		setLoader(session, ChannelSlice),
	}
}
