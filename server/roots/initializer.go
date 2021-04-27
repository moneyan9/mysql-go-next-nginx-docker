package roots

import (
	"github.com/labstack/echo/v4"
)

func InitializeRoots(e *echo.Echo) {

	AddUsersRoots(e)
}
