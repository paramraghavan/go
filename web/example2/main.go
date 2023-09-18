package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	tm := NewTodoManager()
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		todos := tm.GetAll()

		return c.JSON(http.StatusOK, todos)
	})

	e.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")

		todo, err := tm.GetByID(id)
		if err != nil {
			c.Error(err)
			return err
		}

		return c.JSON(http.StatusAccepted, todo)
	})

	authenticatedGroup := e.Group("/todos", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("authorization")
			if authorization != "auth-token" {
				c.Error(echo.ErrUnauthorized)
				return nil
			}

			next(c)
			return nil
		}
	})

	authenticatedGroup.PATCH("/:id/complete", func(c echo.Context) error {
		id := c.Param("id")

		err := tm.Complete(id)
		if err != nil {
			c.Error(err)
			return err
		}

		return nil
	})

	authenticatedGroup.POST("/create", func(c echo.Context) error {
		requestBody := CreateTodoRequest{}

		err := c.Bind(&requestBody)
		if err != nil {
			return err
		}

		err = requestBody.Validate()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		todo := tm.Create(requestBody)
		return c.JSON(http.StatusCreated, todo)
	})

	authenticatedGroup.DELETE("/:id", func(c echo.Context) error {
		id := c.Param("id")

		err := tm.Remove(id)
		if err != nil {
			c.Error(err)
			return err
		}

		return nil
	})

	authenticatedGroup.Any("/*", func(c echo.Context) error {
		c.Error(echo.ErrNotFound)
		return nil
	})

	e.Any("/*", func(c echo.Context) error {
		c.Error(echo.ErrNotFound)
		return nil
	})

	e.Start(":8888")
}
