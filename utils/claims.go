package utils

import (
	reqModel "bms-go/model/request"

	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetClaims(c *gin.Context) (*reqModel.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	calims, err := j.ParseToken(token)
	if err != nil {
		zap.L().Error("ParseToken from gin.Context failed", zap.Error(err))
	}
	return calims, err
}

func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.ID
		}
	} else {
		u := claims.(*reqModel.CustomClaims)
		return u.ID
	}
}

func GetUserUuid(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		u := claims.(*reqModel.CustomClaims)
		return u.UUID
	}
}
