package ecommerce

import (
	"errors"
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorAuth       = errors.New("auth enabled not initialised")
	ErrorPermission = errors.New("permissions enabled not initialised")
	ErrorEmpty      = errors.New("given some values is empty")
	TZONE, _        = time.LoadLocation(os.Getenv("TIME_ZONE"))
	TenantId, _     = strconv.Atoi(os.Getenv("Tenant_ID"))
)

type EcommerceModel struct {
	DataAccess int
	UserId     int
}

var Ecommercemodel EcommerceModel

func AuthandPermission(ecommerce *Ecommerce) error {

	//check auth enable if enabled, use auth pkg otherwise it will return error
	if ecommerce.AuthEnable && !ecommerce.Auth.AuthFlg {

		return ErrorAuth
	}
	//check permission enable if enabled, use team-role pkg otherwise it will return error
	if ecommerce.PermissionEnable && !ecommerce.Auth.PermissionFlg {

		return ErrorPermission

	}

	return nil
}

// Hass Password
func HashingPassword(pass string) string {

	passbyte, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	if err != nil {

		panic(err)

	}

	return string(passbyte)
}
