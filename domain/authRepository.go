package domain

import (
	"encoding/json"
	"github.com/lozhkindm/banking-lib/logger"
	"net/http"
	"net/url"
)

type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}

type RemoteAuthRepository struct{}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {
	u := buildVerifyURL(token, routeName, vars)

	if res, err := http.Get(u); err != nil {
		logger.Error("Error while getting a response from auth server: " + err.Error())
		return false
	} else {
		m := map[string]bool{}

		if err = json.NewDecoder(res.Body).Decode(&m); err != nil {
			logger.Error("Error while decoding a response from auth server: " + err.Error())
			return false
		}

		return m["is_authorized"]
	}
}

func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	u := url.URL{Host: "127.0.0.1:8181", Path: "/auth/verify", Scheme: "http"}
	q := u.Query()

	q.Add("token", token)
	q.Add("route_name", routeName)

	for k, v := range vars {
		q.Add(k, v)
	}

	u.RawQuery = q.Encode()

	return u.String()
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}
