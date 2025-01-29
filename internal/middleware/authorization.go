package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/stierprogrammer/goapi/api"
	"github.com/stierprogrammer/goapi/internal/tools"
);

var UnauthorizedError = errors.New("Invalid username or token!");

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username");
		var token = r.Header.Get("Authorization");
		var err error;

		if username == "" || token == "" {
			log.Error(UnauthorizedError);
			api.RequestErrorHandler(
				w, UnauthorizedError);
			return;
		};

		var database *tools.DatabaseInterface;
		database, err = tools.NewDatabase();

		if err != nil {
			api.InternalErrorHandler(w);
			return;
		};

		var loginDetails *tools.LoginDetails;
		loginDetails = (*database).GetUserLoginDetails(username);

		if (loginDetails == nil || token != (*loginDetails).AuthToken) {
			log.Error(UnauthorizedError);
			api.RequestErrorHandler(w, UnauthorizedError);
			return;
		};

		next.ServeHTTP(w, r);
	});
};