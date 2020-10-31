package root_object

import (
	"context"
	"net/http"
	"strings"

	"github.com/taskalla/api/pkg/models/token"
)

func ResolveRootObject(ctx context.Context, r *http.Request) map[string]interface{} {
	auth := r.Header.Get("Authorization")
	split := strings.Split(auth, "Bearer ")

	if len(split) < 2 {
		return map[string]interface{}{}
	}

	t := split[1]
	token, err := token.GetToken(t)
	if err != nil {
		return map[string]interface{}{}
	}

	return map[string]interface{}{
		"token": token,
	}
}
