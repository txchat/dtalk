package midware

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/txchat/dtalk/pkg/api"
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/backend/model"
)

func JWTAuthMiddleWare(flag bool, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if flag == true {
			c.Next()
			return
		}
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			c.Set(api.ReqError, xerror.NewError(xerror.TokenError).SetExtMessage("token is empty"))
			c.Abort()
			return
		}
		stringArray := strings.SplitN(authHeader, " ", 2)
		if !(len(stringArray) == 2 && stringArray[0] == "Bearer") {
			c.Set(api.ReqError, xerror.NewError(xerror.TokenError).SetExtMessage("The token's format is wrong"))
			c.Abort()
			return
		}
		claims, err := parseToken(stringArray[1], key)
		if err != nil {
			c.Set(api.ReqError, err)
			c.Abort()
			return
		}
		c.Set("userName", claims.Username)
		c.Next()
	}
}

func parseToken(tokenString string, key string) (*model.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, xerror.NewError(xerror.TokenError).SetExtMessage("invalid token")
	}
	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, xerror.NewError(xerror.TokenError).SetExtMessage("invalid token")
}
