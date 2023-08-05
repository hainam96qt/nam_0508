package midleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"

	"nam_0508/internal/model"
	error2 "nam_0508/pkg/error"
	"nam_0508/pkg/util/response"
)

const (
	ContextKey = "IDENTITY"

	StructName = "Identity"
)

type (
	contextKey struct {
		name string
	}

	AuthenticateMiddleware struct {
		Algorithm jwa.SignatureAlgorithm
		SecretKey string
	}
)

var (
	AuthenticateMW *AuthenticateMiddleware
)

func NewAuthenticateMiddleware(algorithm string, secretKey string) *AuthenticateMiddleware {
	return &AuthenticateMiddleware{
		Algorithm: jwa.SignatureAlgorithm(algorithm),
		SecretKey: secretKey,
	}
}

var (
	TokenCtxKey    = &contextKey{"Token"}
	RawTokenCtxKey = &contextKey{"RawToken"}
	ErrorCtxKey    = &contextKey{"Error"}
)

func (a *AuthenticateMiddleware) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		token, err := a.VerifyToken(TokenFromHeader(r))
		if err != nil {
			response.JSON(w, error2.NewXError("invalid token"))
			return
		}

		if err := jwt.Validate(token); err != nil {
			if errors.Is(err, jwt.ErrTokenExpired()) {
				response.JSON(w, error2.NewXError("invalid token"))
				return
			}
			response.JSON(w, error2.NewXError("invalid token"))
			return
		}

		claims, err := GetClaimsToken(r.Context(), token)
		if err != nil || token == nil {
			response.JSON(w, error2.NewXError("invalid token"))
			return
		}

		ctx, err = SetIdentityToContext(ctx, claims)
		if err != nil {
			response.JSON(w, error2.NewXError("invalid token"))
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TokenFromHeader(r *http.Request) string {
	// Get token from authorization header.
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}

func (a *AuthenticateMiddleware) VerifyToken(tokenString string) (jwt.Token, error) {
	verifier := jwt.WithVerify(a.Algorithm, []byte(a.SecretKey))
	return jwt.Parse([]byte(tokenString), verifier)
}

func GetClaimsToken(ctx context.Context, token jwt.Token) (map[string]interface{}, error) {
	var err error
	var claims map[string]interface{}

	if token != nil {
		claims, err = token.AsMap(context.Background())
		if err != nil {
			return nil, err
		}
	} else {
		claims = map[string]interface{}{}
	}

	err, _ = ctx.Value(ErrorCtxKey).(error)

	return claims, err
}

func SetIdentityToContext(ctx context.Context, claims map[string]interface{}) (context.Context, error) {
	jsBody, err := json.Marshal(claims)
	if err != nil {
		return ctx, fmt.Errorf("failed to encode claims to json %s", err)
	}
	var identity model.Identity
	if err := json.Unmarshal(jsBody, &identity); err != nil {
		return ctx, fmt.Errorf("failed to parse encoded claims to identity %s", err)
	}
	if reflect.DeepEqual(identity, model.Identity{}) {
		return ctx, fmt.Errorf("failed to set identity to context, identity empty")
	}
	return context.WithValue(ctx, ContextKey, identity), nil
}

func GetIdentityFromContext(ctx context.Context) (*model.Identity, error) {
	identityFromCtx := ctx.Value(ContextKey)
	if reflect.TypeOf(identityFromCtx) == nil {
		return nil, errors.New("identity not found")
	}

	typeName := reflect.TypeOf(identityFromCtx).Name()
	if !strings.EqualFold(typeName, StructName) {
		return nil, errors.New("identity not found")
	}

	result := identityFromCtx.(model.Identity)
	return &result, nil
}
