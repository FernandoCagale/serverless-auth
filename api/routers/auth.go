package routers

import (
	"os"

	"github.com/FernandoCagale/serverless-auth/api/handlers"
	"github.com/FernandoCagale/serverless-auth/pkg/auth"
	"github.com/FernandoCagale/serverless-auth/pkg/auth/repository"
	"github.com/FernandoCagale/serverless-auth/pkg/auth/usecase"
	"github.com/gorilla/mux"
)

//MakeHandlers make url handlers
func MakeHandlers(r *mux.Router) {
	service := makeGorm()

	r.Handle("/v1/api/login", handlers.Login(service)).Methods("POST")
}

//makeGorm database postgres
func makeGorm() auth.UseCase {
	return usecase.NewService(repository.NewGormRepository(os.Getenv("DATASTORE_URL")))
}
