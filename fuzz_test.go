package ServShared

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

type mockValidatable struct {
	User  string `json:"user"`
	Email string `json:"email"`
}

func (m *mockValidatable) Validate() error {
	return nil
}

func FuzzDecodeAndValidateJSON(f *testing.F) {
	f.Add([]byte(`{"user": "alice", "email": "alice@example.com"}`))
	f.Add([]byte(`{"user": "", "email": ""}`))
	f.Add([]byte(`invalid json`))

	f.Fuzz(func(t *testing.T, data []byte) {
		req := httptest.NewRequest("POST", "/", bytes.NewReader(data))
		w := httptest.NewRecorder()
		var m mockValidatable
		DecodeAndValidateJSON(w, req, &m)
	})
}

func FuzzSanitizeLog(f *testing.F) {
	f.Add("password=secret123")
	f.Add("bearer token-value")
	f.Add("normal message")

	f.Fuzz(func(t *testing.T, data string) {
		SanitizeLog(data)
	})
}
