package controller

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/Mark-Grigorev/Website_ShortURL_Golang/internal/gen"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ValidateToken func(ctx context.Context, req *gen.ValidateTokenRequest, opts ...grpc.CallOption) (*gen.ValidateTokenResponse, error)

type AuthMiddleware struct {
	ValidateToken
}

func NewAuthMiddleware(addr string) (*AuthMiddleware, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return &AuthMiddleware{}, err
	}
	client := gen.NewAuthServiceClient(conn)
	return &AuthMiddleware{
		ValidateToken: client.ValidateToken,
	}, nil
}

func (m *AuthMiddleware) ValidateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("missing_header"))
			return
		}

		parts := strings.Split(token, "")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid_header"))
			return
		}

		resp, err := m.ValidateToken(ctx.Request.Context(), &gen.ValidateTokenRequest{Token: parts[1]})
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, errors.New("cant_validate_token"))
			return
		}

		if !resp.Valid {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("invalid_token"))
			return
		}
		ctx.Set("user_id", resp.UserId)
		ctx.Next()
	}
}
