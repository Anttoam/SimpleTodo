package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Anttoam/golang-htmx-todos/dto"
	"github.com/Anttoam/golang-htmx-todos/views/error_page"
	"github.com/Anttoam/golang-htmx-todos/views/todo"
	"github.com/Anttoam/golang-htmx-todos/views/user"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/rbcervilla/redisstore/v9"
)

type UserUsecase interface {
	SignUp(ctx context.Context, req dto.SignUpRequest) error
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
	FindUserByID(ctx context.Context, userID int) (*dto.FindByIDUserResponse, error)
	EditUser(ctx context.Context, user dto.UpdateUserRequest) error
	EditPassword(ctx context.Context, user dto.UpdatePasswordRequest) error
}

type UserController struct {
	Usecase UserUsecase
	Store   *redisstore.RedisStore
}

func NewUserController(e *echo.Echo, uu UserUsecase, store *redisstore.RedisStore) {
	uc := &UserController{Usecase: uu, Store: store}

	api := e.Group("/user")
	api.GET("/signup", uc.SignUp)
	api.POST("/signup", uc.SignUp)
	api.GET("/login", uc.Login)
	api.POST("/login", uc.Login)
	api.GET("/logout", uc.Logout)
	api.GET("/:id", uc.EditUser)
	api.PUT("/:id", uc.EditUser)
	api.GET("/password/:id", uc.EditPassword)
	api.PUT("/password/:id", uc.EditPassword)
}

// Signup godoc
//
//	@Summary		Signup
//	@Description	Create a new user
//	@Tags			user
//
// Accept json
// Produce json
//
//	@Param			request	body		dto.SignUpRequest	true	"Create User Request"
//	@Success		200		{string}	string				"ok"
//	@Router			/user/signup [post]
func (uc *UserController) SignUp(c echo.Context) error {
	if c.Request().Method == http.MethodPost {
		req := dto.SignUpRequest{
			Name:     c.FormValue("name"),
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
		}

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		ctx := c.Request().Context()
		if err := uc.Usecase.SignUp(ctx, req); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.Redirect(http.StatusSeeOther, "/user/login")
	}

	component := user.SignUp()
	handler := echo.WrapHandler(templ.Handler(component))
	return handler(c)
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login user
//	@Tags			user
//
// Accept json
// Produce json
//
//	@Param			request	body	dto.LoginRequest	true	"Login Request"
//	@Router			/user/login [post]
func (uc *UserController) Login(c echo.Context) error {
	if c.Request().Method == http.MethodPost {
		req := dto.LoginRequest{
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
		}
		if err := c.Bind(&req); err != nil {
			log.Println("Bind error: ", err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		ctx := c.Request().Context()
		res, err := uc.Usecase.Login(ctx, req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		sess, _ := uc.Store.Get(c.Request(), "session_id")
		sess.Values["id"] = res.ID
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.Redirect(http.StatusSeeOther, "/todo/")
	}

	component := user.Login()
	handler := echo.WrapHandler(templ.Handler(component))
	return handler(c)
}

// Logout godoc
//
//	@Summary		Logout
//	@Description	Logout user
//	@Tags			user
//
//	@Router			/user/logout [get]
func (uc *UserController) Logout(c echo.Context) error {
	sess, _ := uc.Store.Get(c.Request(), "session_id")
	sess.Values["id"] = nil
	sess.Options.MaxAge = -1
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.Redirect(http.StatusFound, "/user/login")
}

// EditUser godoc
//
//	@Summary		Edit User
//	@Description	Edit user
//	@Tags			user
//
// Accept json
// Produce json
//
//	@Param			id		path	int						true	"User ID"
//	@Param			request	body	dto.UpdateUserRequest	true	"Edit User Request"
//	@Router			/user/:id [put]
func (uc *UserController) EditUser(c echo.Context) error {
	var userID int
	sess, _ := uc.Store.Get(c.Request(), "session_id")
	id := sess.Values["id"]
	if id == nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
		c.Response().WriteHeader(http.StatusUnauthorized)
		if err := error_page.Error401().Render(c.Request().Context(), c.Response().Writer); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
	userID = id.(int)

	var fetch *dto.FindByIDUserResponse

	if c.Request().Method == http.MethodPut {
		req := dto.UpdateUserRequest{
			Name:  c.FormValue("name"),
			Email: c.FormValue("email"),
		}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		sess, _ := uc.Store.Get(c.Request(), "session_id")
		id := sess.Values["id"]
		if id == nil {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
			c.Response().WriteHeader(http.StatusUnauthorized)
			if err := error_page.Error401().Render(c.Request().Context(), c.Response().Writer); err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			return nil
		}
		userID := id.(int)

		idP, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		req.ID = idP

		if err := c.Validate(req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		ctx := c.Request().Context()
		fetch, err := uc.Usecase.FindUserByID(ctx, req.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if fetch.User.ID != userID {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
			c.Response().WriteHeader(http.StatusUnauthorized)
			if err := error_page.Error401().Render(c.Request().Context(), c.Response().Writer); err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			return nil
		}

		if err := uc.Usecase.EditUser(ctx, req); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/user/%d", fetch.User.ID))

	}

	ctx := c.Request().Context()
	fetch, err := uc.Usecase.FindUserByID(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	component := todo.EditUser(strconv.Itoa(userID), fetch.User.Name, fetch.User.Email)
	handler := echo.WrapHandler(templ.Handler(component))
	return handler(c)
}

// EditPassword godoc
//
//	@Summary		Edit Password
//	@Description	Edit password
//	@Tags			user
//
// Accept json
// Produce json
//
//	@Param			id		path	int							true	"User ID"
//	@Param			request	body	dto.UpdatePasswordRequest	true	"Edit Password Request"
//	@Router			/user/password/:id [put]
func (uc *UserController) EditPassword(c echo.Context) error {
	var userID int
	sess, _ := uc.Store.Get(c.Request(), "session_id")
	id := sess.Values["id"]
	if id == nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
		c.Response().WriteHeader(http.StatusUnauthorized)
		if err := error_page.Error401().Render(c.Request().Context(), c.Response().Writer); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
	userID = id.(int)

	if c.Request().Method == http.MethodPut {
		req := dto.UpdatePasswordRequest{
			Password:    c.FormValue("password"),
			NewPassword: c.FormValue("new_password"),
		}
		if err := c.Bind(&req); err != nil {
			return err
		}

		sess, _ := uc.Store.Get(c.Request(), "session_id")
		id := sess.Values["id"]
		if id == nil {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
			c.Response().WriteHeader(http.StatusUnauthorized)
			if err := error_page.Error401().Render(c.Request().Context(), c.Response().Writer); err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			return nil
		}
		userID := id.(int)

		idP, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		req.ID = idP

		if err := c.Validate(req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		ctx := c.Request().Context()
		fetch, err := uc.Usecase.FindUserByID(ctx, req.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if fetch.User.ID != userID {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
			c.Response().WriteHeader(http.StatusUnauthorized)
			if err := error_page.Error401().Render(c.Request().Context(), c.Response().Writer); err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			return nil
		}

		if err := uc.Usecase.EditPassword(ctx, req); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/user/%d", fetch.User.ID))
	}

	component := todo.EditPassword(strconv.Itoa(userID))
	handler := echo.WrapHandler(templ.Handler(component))
	return handler(c)
}
