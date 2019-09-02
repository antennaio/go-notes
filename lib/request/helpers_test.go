package request

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-chi/chi"
)

func TestParamIntHelper(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	rctx := chi.NewRouteContext()
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	rctx.URLParams.Add("id", "8")

	id, _ := ParamInt(r, "id")
	if id != 8 {
		t.Fatalf("expected: %v got: %v", 8, id)
	}
}
