package main

import (
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
	"github.com/souz9/errlist"
)

func latest(upstream string, stg *Storage) http.Handler {
	cl := gorequest.New()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mod, _ := moduleVersionFromPath(r.URL.Path)
		latest, err := stg.latest(mod, func(req string) (string, error) {
			var ret string
			res, _, err := cl.
				Get(fmt.Sprintf("%s/%s/@latest", upstream, mod)).
				EndStruct(&ret)
			if err != nil {
				return "", errors.WithStack(errlist.Error(err))
			}
			if res.StatusCode != http.StatusOK {
				return "", errors.WithStack(
					fmt.Errorf("upstream returned code %d", res.StatusCode),
				)
			}
			return ret, nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte(latest))
	})
}
