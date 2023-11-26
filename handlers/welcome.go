package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Welcome is an example handler function
func (h *Handler) Welcome(c echo.Context) error {
    // Status, Component Name, Data to pass on
     return c.Render(http.StatusOK, "Welcome", map[string]interface{}{})
}
