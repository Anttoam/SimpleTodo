package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Anttoam/golang-htmx-todos/dto"
	"github.com/Anttoam/golang-htmx-todos/views/error_page"
	"github.com/Anttoam/golang-htmx-todos/views/todo"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/rbcervilla/redisstore/v9"
)

type TodoUsecase interface {
	Create(ctx context.Context, req dto.CreateTodoRequest) error
	FindAll(ctx context.Context, userID int) (*dto.FindAllTodoResponse, error)
	FindByID(ctx context.Context, todoID int) (*dto.FindByIDTodoResponse, error)
	Update(ctx context.Context, req dto.UpdateTodoRequest) error
	Delete(ctx context.Context, todoID int) error
	IsDone(ctx context.Context, todoID int) error
	IsNotDone(ctx context.Context, todoID int) error
}

type TodoController struct {
	Usecase TodoUsecase
	Store   *redisstore.RedisStore
}

func NewTodoController(e *echo.Echo, tu TodoUsecase, store *redisstore.RedisStore) {
	t := &TodoController{Usecase: tu, Store: store}

	api := e.Group("/todo")
	api.GET("/create", t.Create)
	api.POST("/create", t.Create)
	api.GET("/", t.FindAll)
	api.GET("/:id", t.FindByID)
	api.PUT("/:id", t.Update)
	api.DELETE("/:id", t.Delete)
	api.PUT("/done/:id", t.IsDone)
	api.PUT("/notdone/:id", t.IsNotDone)
}

// Create godoc
//
//	@Summary		Create
//	@Description	Create a new todo
//	@Tags			todo
//
// Accept json
// Produce json
//
//	@Param			request	body		dto.SignUpRequest	true	"Create User Request"
//	@Success		303		{string}	string				"ok"
//	@Router			/todo/create [post]
func (t *TodoController) Create(c echo.Context) error {
	sess, _ := t.Store.Get(c.Request(), "session_id")
	id := sess.Values["id"]
	if id == nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
		c.Response().WriteHeader(http.StatusUnauthorized)
		if err := error_page.Error401().Render(c.Request().Context(), c.Response().Writer); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return nil
	}

	if c.Request().Method == http.MethodPost {
		req := dto.CreateTodoRequest{
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
		}
		if err := c.Bind(&req); err != nil {
			return err
		}

		req.UserID = id.(int)

		if err := c.Validate(req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		ctx := c.Request().Context()
		if err := t.Usecase.Create(ctx, req); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.Redirect(http.StatusSeeOther, "/todo/")
	}

	component := todo.Create()
	handler := echo.WrapHandler(templ.Handler(component))
	return handler(c)
}

// FindAll godoc
//
//	@Summary		FindAll
//	@Description	Find all todo
//	@Tags			todo
//
// Accept json
// Produce json
//
//	@Success		200	{string}	string	"ok"
//	@Router			/todo/ [get]
func (t *TodoController) FindAll(c echo.Context) error {
	sess, _ := t.Store.Get(c.Request(), "session_id")
	id := sess.Values["id"]
	if id == nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
		c.Response().WriteHeader(http.StatusUnauthorized)
		if err := error_page.Error401().Render(c.Request().Context(), c.Response().Writer); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return nil
	}

	ctx := c.Request().Context()
	userID := id.(int)
	res, err := t.Usecase.FindAll(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	component := todo.Page(*res, strconv.Itoa(userID))
	handler := echo.WrapHandler(templ.Handler(component))
	return handler(c)
}

// FindByID godoc
//
//	@Summary		FindByID
//	@Description	Find todo by ID
//	@Tags			todo
//
// Accept json
// Produce json
//
//	@Param			id	path		int		true	"Todo ID"
//	@Success		200	{string}	string	"ok"
//	@Router			/todo/:id [get]
func (t *TodoController) FindByID(c echo.Context) error {
	sess, _ := t.Store.Get(c.Request(), "session_id")
	id := sess.Values["id"]
	if id == nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
		c.Response().WriteHeader(http.StatusUnauthorized)
		if err := error_page.Error401().Render(c.Request().Context(), c.Response().Writer); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return nil
	}

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	todoID := idP

	ctx := c.Request().Context()
	res, err := t.Usecase.FindByID(ctx, todoID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if res.Todo.UserID != id.(int) {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
		c.Response().WriteHeader(http.StatusUnauthorized)
		if err := error_page.Error401().Render(c.Request().Context(), c.Response().Writer); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return nil
	}

	component := todo.EditForm(strconv.Itoa(res.Todo.ID), res.Todo.Title, res.Todo.Description)
	handler := echo.WrapHandler(templ.Handler(component))
	return handler(c)
}

// Update godoc
//
//	@Summary		Update
//	@Description	Update todo
//	@Tags			todo
//
// Accept json
// Produce json
//
//	@Param			id		path		int						true	"Todo ID"
//	@Param			request	body		dto.UpdateTodoRequest	true	"Update Todo Request"
//	@Success		200		{string}	string					"ok"
//	@Router			/todo/:id [put]
func (t *TodoController) Update(c echo.Context) error {
	req := dto.UpdateTodoRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
	}
	if err := c.Bind(&req); err != nil {
		return err
	}

	sess, _ := t.Store.Get(c.Request(), "session_id")
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
	fetch, err := t.Usecase.FindByID(ctx, req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if fetch.Todo.UserID != userID {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
		c.Response().WriteHeader(http.StatusUnauthorized)
		if err := error_page.Error401().Render(c.Request().Context(), c.Response().Writer); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
	if err := t.Usecase.Update(ctx, req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := t.Usecase.FindAll(ctx, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	component := todo.List(*res)
	handler := echo.WrapHandler(templ.Handler(component))
	return handler(c)
}

// Delete godoc
//
//	@Summary		Delete
//	@Description	Delete todo
//	@Tags			todo
//
// Accept json
// Produce json
//
//	@Param			id	path		int		true	"Todo ID"
//	@Success		200	{string}	string	"ok"
//	@Router			/todo/:id [delete]
func (t *TodoController) Delete(c echo.Context) error {
	sess, _ := t.Store.Get(c.Request(), "session_id")
	id := sess.Values["id"]
	if id == nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
		c.Response().WriteHeader(http.StatusUnauthorized)
		if err := error_page.Error401().Render(c.Request().Context(), c.Response().Writer); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return nil
	}

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	ctx := c.Request().Context()
	res, err := t.Usecase.FindByID(ctx, idP)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if res.Todo.UserID != id.(int) {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
		c.Response().WriteHeader(http.StatusUnauthorized)
		if err := error_page.Error401().Render(c.Request().Context(), c.Response().Writer); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return nil
	}

	if err := t.Usecase.Delete(ctx, res.Todo.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	component := todo.Delete(strconv.Itoa(res.Todo.ID))
	handler := echo.WrapHandler(templ.Handler(component))
	return handler(c)
}

func (t *TodoController) IsDone(c echo.Context) error {
	return t.changeStatus(c, t.Usecase.IsDone)
}

func (t *TodoController) IsNotDone(c echo.Context) error {
	return t.changeStatus(c, t.Usecase.IsNotDone)
}
