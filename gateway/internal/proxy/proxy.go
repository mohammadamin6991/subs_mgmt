package proxy

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	utils "gitlab.amin.run/general/project/subs-mgmt/gateway/pkg"
)

func SubscriptionProxy() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if os.Getenv("SUBSCRIPTION_ENDPOINT") == "" {
			utils.ErrorJSON(w, errors.New("the SUBSCRIPTION_ENDPOINT environment variable is missing"), http.StatusBadRequest)
			return
		}

		url, err := url.Parse(os.Getenv("SUBSCRIPTION_ENDPOINT"))
		if err != nil {
			utils.ErrorJSON(w, fmt.Errorf("can not parse SUBSCRIPTION_ENDPOINT value as url: %s", err), http.StatusBadRequest)
			return
		}
		httpProxy := httputil.NewSingleHostReverseProxy(url)
		httpProxy.ServeHTTP(w, r)
	})
}

func SvcS3Proxy() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if os.Getenv("SVC_S3_ENDPOINT") == "" {
			utils.ErrorJSON(w, errors.New("the SVC_S3_ENDPOINT environment variable is missing"), http.StatusBadRequest)
			return
		}

		url, err := url.Parse(os.Getenv("SVC_S3_ENDPOINT"))
		if err != nil {
			utils.ErrorJSON(w, fmt.Errorf("can not parse SVC_S3_ENDPOINT value as url: %s", err), http.StatusBadRequest)
			return
		}
		httpProxy := httputil.NewSingleHostReverseProxy(url)
		httpProxy.ServeHTTP(w, r)
	})
}

func AuthAPIProxy() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if os.Getenv("AUTH_ENDPOINT") == "" {
			utils.ErrorJSON(w, errors.New("the AUTH_ENDPOINT environment variable is missing"), http.StatusBadRequest)
			return
		}

		url, err := url.Parse(os.Getenv("AUTH_ENDPOINT"))
		if err != nil {
			utils.ErrorJSON(w, fmt.Errorf("can not parse AUTH_ENDPOINT value as url: %s", err), http.StatusBadRequest)
			return
		}

		httpProxy := httputil.NewSingleHostReverseProxy(url)
		httpProxy.ServeHTTP(w, r)
	})
}
