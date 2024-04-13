package banners

import "github.com/labstack/echo/v4"

type Handlers interface {
	Create() echo.HandlerFunc
	GetContent() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc

	// Небольшой middleware
	AdminAuth(next echo.HandlerFunc) echo.HandlerFunc
	UserAuth(next echo.HandlerFunc) echo.HandlerFunc
}
