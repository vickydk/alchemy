package http

import "github.com/labstack/echo/v4"

type HTTPError struct {
	Message interface{} `json:"message"`
}

type HTTPResponse struct {
	Success bool `json:"success"`
}

func httpError(c echo.Context, code int, err error) error {
	errResp := &HTTPError{Message: err.Error()}
	if err = c.JSON(code, errResp); err != nil {
		return err
	}
	return err
}

func httpOk(c echo.Context, code int, data interface{}) (err error) {
	return c.JSON(code, data)
}

func httpResp(c echo.Context, code int, success bool) (err error) {
	return c.JSON(code, &HTTPResponse{Success: success})
}
