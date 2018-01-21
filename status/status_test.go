package status_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adamluzsi/httpwiki-go/status"
	"github.com/stretchr/testify/require"
)

func TestInternalServerError(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	status.InternalServerError(w)

	require.Equal(t, 500, w.Code)
	require.Contains(t, w.Body.String(), http.StatusText(http.StatusInternalServerError))
}

func TestNotFound(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	status.NotFound(w)

	require.Equal(t, 404, w.Code)
	require.Contains(t, w.Body.String(), http.StatusText(http.StatusNotFound))
}
