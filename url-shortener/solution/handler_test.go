package urlshortener

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMapHandler(t *testing.T) {
	pathToURL := map[string]string{
		"/example": "https://example.com/",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	})

	tcs := map[string]struct {
		requestURL   string
		responseURL  string
		responseCode int
	}{
		"existing":     {requestURL: "/example", responseURL: "https://example.com/", responseCode: http.StatusFound},
		"non-existing": {requestURL: "/something", responseURL: "", responseCode: http.StatusBadGateway},
	}

	for title, tc := range tcs {
		t.Run(title, func(t *testing.T) {
			handler := MapHandler(pathToURL, mux)

			request := httptest.NewRequest("GET", tc.requestURL, nil)
			response := httptest.NewRecorder()

			handler.ServeHTTP(response, request)

			require.Equal(t, tc.responseURL, response.HeaderMap.Get("Location"))
			require.Equal(t, tc.responseCode, response.Code)
		})
	}
}

func TestYAMLHandler(t *testing.T) {
	yaml := `
- path: /example
  url: "https://example.com/"
`

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	})

	tcs := map[string]struct {
		requestURL   string
		responseURL  string
		responseCode int
	}{
		"existing":     {requestURL: "/example", responseURL: "https://example.com/", responseCode: http.StatusFound},
		"non-existing": {requestURL: "/something", responseURL: "", responseCode: http.StatusBadGateway},
	}

	for title, tc := range tcs {
		t.Run(title, func(t *testing.T) {
			handler, _ := YAMLHandler([]byte(yaml), mux)

			request := httptest.NewRequest("GET", tc.requestURL, nil)
			response := httptest.NewRecorder()

			handler.ServeHTTP(response, request)

			require.Equal(t, tc.responseURL, response.HeaderMap.Get("Location"))
			require.Equal(t, tc.responseCode, response.Code)
		})
	}
}
