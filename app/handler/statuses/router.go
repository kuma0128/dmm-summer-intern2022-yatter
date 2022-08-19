package statuses

import (
	"net/http"

	"yatter-backend-go/app/app"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi"
)

// Implementation of handler
type handler struct {
	app *app.App
}

// Create Handler for `/v1/statuses/`
func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()

	h := &handler{app: app}
	//r.With(middleware.Logger).Post("/", h.Create)
	r.With(auth.Middleware(app)).Post("/", h.Create)
	r.With(auth.Middleware(app)).Delete("/{id}", h.Delete)
	r.Get("/{id}", h.GetStatus)

	return r
}
