package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"github.com/sirupsen/logrus"
	"golang_user/models"
	_ "strconv"
	"time"
)

type UserController struct {
	ctx    iris.Context
	Logger *logrus.Logger
}

func (c *UserController) Init(ctx iris.Context) {
	c.ctx = ctx
}

func (c *UserController) Ctx() iris.Context {
	return c.ctx
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *UserController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/", "List")
	b.Handle("GET", "/{id:uint}", "GetUserByID")
	b.Handle("POST", "/", "CreateUser")
	b.Handle("PUT", "/{id:uint}", "UpdateUser")
	b.Handle("DELETE", "/{id:uint}", "DeleteUser")
}

func (c *UserController) List(context.Context) mvc.Response {
	fmt.Sprintln("TEST")
	users, err := models.GetList()
	if err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError,
			Object: iris.Map{
				"message": err.Error(),
			},
		}
	}
	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, UserResponse{
			ID:        user.ID,
			Name:      user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return mvc.Response{
		Code:   iris.StatusOK,
		Object: userResponses,
	}
}

func (c *UserController) GetUserByID(id uint) mvc.Result {
	user, err := models.GetUserByID(id)
	if err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError,
			Object: iris.Map{
				"message": err.Error(),
			},
		}
	}
	userResponse := UserResponse{
		ID:        user.ID,
		Name:      user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return mvc.Response{
		Code:   iris.StatusOK,
		Object: userResponse,
	}
}

func (c *UserController) CreateUser(context.Context) mvc.Response {
	c.Init(c.ctx)
	//fmt.Printf("request body: %v\n", c.ctx.Request().Body)
	var user models.User
	if c == nil || c.ctx == nil {
		if c.ctx == nil {
			// c.ctx為nil，回傳錯誤
			return mvc.Response{
				Code: iris.StatusInternalServerError,
				Object: iris.Map{
					"message": "UserController Ctx is nil",
				},
			}
		}
		// c或c.ctx為nil，回傳錯誤
		return mvc.Response{
			Code: iris.StatusInternalServerError,
			Object: iris.Map{
				"message": "UserController or Ctx is nil",
			},
		}
	}
	if err := c.ctx.ReadJSON(&user); err != nil {
		return mvc.Response{
			Code: iris.StatusBadRequest,
			Object: iris.Map{
				"message": err.Error(),
			},
		}
	}

	if err := user.Create(); err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError,
			Object: iris.Map{
				"message": err.Error(),
			},
		}
	}

	return mvc.Response{
		Code:   iris.StatusOK,
		Object: user,
	}
}

func (c *UserController) UpdateUser(id uint) mvc.Result {
	var user models.User
	if err := c.ctx.ReadJSON(&user); err != nil {
		return mvc.Response{
			Code: iris.StatusBadRequest,
			Object: iris.Map{
				"message": err.Error(),
			},
		}
	}

	user.ID = id
	if err := user.Update(); err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError,
			Object: iris.Map{
				"message": err.Error(),
			},
		}
	}

	return mvc.Response{
		Code:   iris.StatusOK,
		Object: user,
	}
}

func (c *UserController) DeleteUser(id uint) mvc.Result {
	if err := models.DeleteUserByID(id); err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError,
			Object: iris.Map{
				"message": err.Error(),
			},
		}
	}

	return mvc.Response{
		Code: iris.StatusOK,
		Object: iris.Map{
			"message": "User deleted successfully",
		},
	}
}
