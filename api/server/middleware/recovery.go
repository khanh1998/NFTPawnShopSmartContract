package middleware

import "github.com/gin-gonic/gin"

type APIErrors struct {
	Errors []*APIError `json:"errors"`
}

func (errors *APIErrors) Status() int {
	return errors.Errors[0].Status
}

type APIError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Title   string `json:"title"`
	Details string `json:"details"`
	Href    string `json:"href"`
}

func newAPIError(status int, code string, title string, details string, href string) *APIError {
	return &APIError{
		Status:  status,
		Code:    code,
		Title:   title,
		Details: details,
		Href:    href,
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case *APIError:
					apiError := err.(*APIError)
					apiErrors := &APIErrors{
						Errors: []*APIError{apiError},
					}
					c.JSON(apiError.Status, apiErrors)
				case *APIErrors:
					apiErrors := err.(*APIErrors)
					c.JSON(apiErrors.Status(), apiErrors)
				}
				panic(err)
			}
		}()

		c.Next()
	}
}
