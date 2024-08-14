package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/samuel-alves-chagas/AskMeAnithing/internal/store/pgstore"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{
		q: q,
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/subscribe/{room_id}", a.handleSubscribe)

	r.Route("/api", func(r chi.Router) {
		r.Route("/rooms", func(r chi.Router) {
			r.Post("/", a.handleCreateRoom)
			r.Get("/", a.handleGetRooms)

			r.Route("/{room_id}/messages", func(r chi.Router) {
				r.Post("/", a.handleCreateRoomMessage)
				r.Get("/", a.handleGetRoomMessages)

				r.Route("/{message_id}", func(r chi.Router) {
					r.Get("/", a.handleGetRoomMessage)
					r.Patch("/react", a.handleReactToMessage)
					r.Delete("/react", a.handleRemovingReactFromMessage)
					r.Patch("/answer", a.handleMarkMessageAsAnswered)
				})
			})
		})
	})
	a.r = r
	return a
}

func (h apiHandler) handleSubscribe(r http.ResponseWriter, w *http.Request) {

}

func (h apiHandler) handleCreateRoom(r http.ResponseWriter, w *http.Request) {

}

func (h apiHandler) handleGetRooms(r http.ResponseWriter, w *http.Request) {

}

func (h apiHandler) handleCreateRoomMessage(r http.ResponseWriter, w *http.Request) {

}

func (h apiHandler) handleGetRoomMessages(r http.ResponseWriter, w *http.Request) {

}

func (h apiHandler) handleGetRoomMessage(r http.ResponseWriter, w *http.Request) {

}

func (h apiHandler) handleReactToMessage(r http.ResponseWriter, w *http.Request) {

}

func (h apiHandler) handleRemovingReactFromMessage(r http.ResponseWriter, w *http.Request) {

}

func (h apiHandler) handleMarkMessageAsAnswered(r http.ResponseWriter, w *http.Request) {

}
