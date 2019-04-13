package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
	"github.com/souz9/errlist"
)

func list(upstream string, stg *Storage) http.Handler {
	cl := gorequest.New()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mod, _ := moduleVersionFromPath(r.URL.Path)
		list, err := stg.list(mod, func(req string) ([]string, error) {
			var ret []string
			res, _, err := cl.
				Get(fmt.Sprintf("%s/%s/@v/list", upstream, mod)).
				EndStruct(&ret)
			if err != nil {
				return nil, errors.WithStack(errlist.Error(err))
			}
			if res.StatusCode != http.StatusOK {
				return nil, errors.WithStack(
					fmt.Errorf("upstream returned code %d", res.StatusCode),
				)
			}
			return ret, nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte(strings.Join(list, "\n")))
	})
}
