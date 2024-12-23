package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CheckUserRoles(c *gin.Context, role string) (err error) {
	userRole := c.GetString("role")
	if userRole != role {
		err = errors.New("user does not have permission to access this resource")
		return err
	}
	return nil
}

func MatchUserRoleToUid(c *gin.Context, userId uuid.UUID) (err error) {
	userRole := c.GetString("role")
	uid,err := uuid.Parse(c.GetString("uid"))
	if err != nil {
		return err
	}
	if userRole == "user" && userId != uid {
		err = errors.New("user does not have permission to access this resource")
		return err
	}
	err = CheckUserRoles(c, userRole)
	return err
}