package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	var tests = []struct {
		name       string
		header     string
		headerType string
		want       string
	}{
		{"successful key get", "ApiKey testey", "Authorization", "testKey"},
		{"missing authorization header", "ApiKey testKey", "", ""},
		{"no ApiKey header", "bearer testKey", "Authorization", ""},
		{"nothing after header", "ApiKey", "Authorization", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testHeader := http.Header{}
			testHeader.Set(test.headerType, test.header)
			key, _ := GetAPIKey(testHeader)

			if key != test.want {
				t.Fatalf("expected %v, got %v", test.want, key)
			}
		})
	}
}
