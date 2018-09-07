package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/FernandoCagale/serverless-auth/pkg/auth"
	"github.com/FernandoCagale/serverless-auth/pkg/entity"
	"github.com/FernandoCagale/serverless-infra/errors"
	"github.com/FernandoCagale/serverless-infra/logger"
	"github.com/FernandoCagale/serverless-infra/render"
)

type jwt struct {
	Token string `json:"token"`
}

//Login jwt
func Login(service auth.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var auth *entity.Auth

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&auth); err != nil {
			render.ResponseError(w, errors.AddBadRequestError("Invalid request payload"))
			return
		}

		logger.WithFields(logger.Fields{
			"Username": auth.Username,
		}).Info("login")

		defer r.Body.Close()

		token, err := service.Login(auth)
		if err != nil {
			switch err {
			case entity.ErrInvalidPayload:
				render.ResponseError(w, errors.AddBadRequestError(err.Error()))
			default:
				render.ResponseError(w, errors.AddInternalServerError(err.Error()))
			}
			return
		}

		render.Response(w, jwt{token}, http.StatusOK)
	})
}
