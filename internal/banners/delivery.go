package banners

import "github.com/labstack/echo/v4"

type Handlers interface {
	Create() echo.HandlerFunc
	GetContent() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc

	// Небольшой middleware
	UserAuth(next echo.HandlerFunc) echo.HandlerFunc
}
