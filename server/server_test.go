package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEmployee(t *testing.T) {
	testCases := []struct {
		desc          string
		server        *httptest.Server
		resp          *Employee
		expectedError error
	}{
		{
			desc: "Get employee",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"id": 1, "name": "John", "age": 30}`))
			})),
			resp:          &Employee{Id: 1, Name: "John", Age: 30},
			expectedError: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			defer tc.server.Close()
			resp, err := GetEmployee(tc.server.URL)
			assert.Equal(t, resp, tc.resp)
			assert.Nil(t, err)
		})
	}
}
