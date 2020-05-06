package main

import (
	"net/http"

	"github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/common/auth"
	"github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/common/server/httperr"
	"github.com/go-chi/render"
)

type HttpServer struct {
	db db
}

func (h HttpServer) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	authUser, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.Unauthorised("no-user-found", err, w, r)
		return
	}

	user, err := h.db.GetUser(r.Context(), authUser.UUID)
	if err != nil {
		httperr.InternalError("cannot-get-user", err, w, r)
		return
	}
	user.Role = authUser.Role
	user.DisplayName = authUser.DisplayName

	render.Respond(w, r, user)
}
