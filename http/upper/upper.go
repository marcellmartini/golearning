package upper

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func UpperCaseHandler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request")
		return
	}

	word := query.Get("word")
	if len(word) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "missing word")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, strings.ToUpper(word))
}
