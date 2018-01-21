package httpwiki_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	httpwiki "github.com/adamluzsi/httpwiki-go"
	"github.com/stretchr/testify/require"
)

func NewWiki(expectedErrors ...error) (*httpwiki.HTTPWiki, *Storage) {
	storage := NewStorage(expectedErrors)
	templates := Templates{}
	wiki := httpwiki.New(storage, templates)
	return wiki, storage
}

func TestWiki_ServeHTTP_PostNewContent_PersistSucceed(t *testing.T) {
	t.Parallel()

	wiki, storage := NewWiki()
	expectedContent := []byte("# hello world!")

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer(expectedContent))

	wiki.ServeHTTP(w, r)

	actuallySavedData, err := storage.Load("/test")

	require.Nil(t, err)
	require.Equal(t, expectedContent, actuallySavedData)

}

func TestWiki_ServeHTTP_PostNewContent_PersistFail(t *testing.T) {
	t.Parallel()

	wiki, _ := NewWiki(errors.New("Boom!"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/test", bytes.NewBuffer([]byte{}))

	wiki.ServeHTTP(w, r)

	require.Equal(t, 500, w.Code)
	require.NotNil(t, w.Body)
	require.Contains(t, string(w.Body.Bytes()), http.StatusText(http.StatusInternalServerError))
}

func TestWiki_ServeHTTP_GetContentWhenNoWikiPageSavedToTheGivenPath_Redirect(t *testing.T) {
	t.Parallel()

	wiki, _ := NewWiki(errors.New("content not found!"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/test", bytes.NewBuffer([]byte{}))

	wiki.ServeHTTP(w, r)

	require.Equal(t, 307, w.Code)
	require.Equal(t, "/test?edit=TRUE", w.Header().Get("Location"))

}

func TestWiki_ServeHTTP_GetContent_ContentFoundAndServedWithTemplate(t *testing.T) {
	t.Parallel()

	message := "Hello world!"
	data := []byte(message)

	wiki, store := NewWiki()
	store.Save("/test", data)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/test", bytes.NewBuffer([]byte{}))

	wiki.ServeHTTP(w, r)

	require.Equal(t, 200, w.Code)
	require.NotNil(t, w.Body)
	require.Equal(t, w.Body.Bytes(), Templates{}.ShowPage(data))
}

func TestWiki_ServeHTTP_GetContentWithEditQueryParameter_EditPageServed(t *testing.T) {
	t.Parallel()

	message := "Hello world!"
	data := []byte(message)

	wiki, store := NewWiki()
	store.Save("/test", data)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/test?edit=true", bytes.NewBuffer([]byte{}))

	wiki.ServeHTTP(w, r)

	require.Equal(t, 200, w.Code)
	require.NotNil(t, w.Body)
	require.Equal(t, w.Body.Bytes(), Templates{}.EditPage(data))
}
