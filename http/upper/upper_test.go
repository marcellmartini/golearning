package upper

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpperCaseHandler(t *testing.T) {
	tests := []struct {
		name string
		uri  string
		want string
	}{
		{
			name: "Upper correct",
			uri:  "/upper?word=abc",
			want: "ABC",
		},
		{
			name: "Missing Word",
			uri:  "/upper",
			want: "missing word",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.uri, nil)
			w := httptest.NewRecorder()

			UpperCaseHandler(w, req)

			res := w.Result()
			defer res.Body.Close()

			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}

			if string(data) != tt.want {
				t.Errorf("upperCaseHandler(w, %v) = %v, got: %v", tt.uri, tt.want, string(data))
			}
		})

	}
}
