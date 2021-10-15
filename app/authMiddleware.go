package app

import (
	"github.com/gorilla/mux"
	"github.com/lozhkindm/banking-lib/errs"
	"github.com/lozhkindm/banking/domain"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

func (a AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")

			if authHeader != "" {
				token := getTokenFromHeader(authHeader)
				isAuthorized := a.repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)

				if isAuthorized {
					next.ServeHTTP(w, r)
				} else {
					err := errs.NewUnauthorizedError("unauthorized")
					writeResponse(w, err.Code, err.AsMessage())
				}
			} else {
				err := errs.NewUnauthorizedError("missing token")
				writeResponse(w, err.Code, err.AsMessage())
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	splitToken := strings.Split(header, "Bearer")

	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}

	return ""
}
