package meta_list

import (
	"net/http"

	"github.com/MunifTanjim/stremthru/internal/config"
	"github.com/MunifTanjim/stremthru/internal/server"
	"github.com/MunifTanjim/stremthru/internal/shared"
)

var IsMethod = shared.IsMethod
var SendError = shared.SendError
var SendResponse = shared.SendResponse

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := server.GetReqCtx(r)
		ctx.Log = log.With("request_id", ctx.RequestId)
		next.ServeHTTP(w, r)
	})
}

func AddEndpoints(mux *http.ServeMux) {
	router := http.NewServeMux()

	if config.Integration.Letterboxd.IsEnabled() {
		router.HandleFunc("/letterboxd/{list_id}", handleGetLetterboxdListById)
	}

	mux.Handle("/v0/meta/lists/", http.StripPrefix("/v0/meta/lists", commonMiddleware(router)))
}
