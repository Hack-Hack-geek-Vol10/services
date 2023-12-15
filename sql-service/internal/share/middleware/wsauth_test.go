package middleware

import (
	"testing"

	"github.com/schema-creator/services/sql-service/pkg/logger"
)

const (
	token = "ew0KCSJwcm9qZWN0X2lkIjogImRhZGE3ZmZjLTU5NzQtNGI1OS04NWIzLWQ1YjJlMTQwYmQxNCIsDQoJInVzZXJfaWQiOiAgICAidGVzdCIsDQoJInJvbGUiOiAgICAgICAiaG9nZSIsDQoJImF1dGhfdG9rZW4iOiAiaG9nZSINCn0NCg=="
)

func TestDecodeURI(t *testing.T) {
	l := logger.NewLogger(logger.DEBUG)
	middleware := NewMiddleware(l)
	middleware.decodeURI(token)

}
