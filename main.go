// main.go
package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/sirupsen/logrus"
	"golang_user/config"
	"golang_user/controllers"
	"golang_user/models"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.UseGlobal(func(ctx iris.Context) {
		ctx.Application().Logger().Debugf("Method: %s , Path: %s", ctx.Method(), ctx.Path())
		ctx.Next()
	})

	logger := logrus.New()

	// initialize database
	db, err := config.ConnectDatabase()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	// migrate database
	db.AutoMigrate(&models.User{})

	//userController := &controllers.UserController{
	//	Logger: logger,
	//}
	//
	//userApp := mvc.New(app.Party("/users"))
	//userApp.Register(userController)
	//userApp.Handle(new(controllers.UserController))

	userController := &controllers.UserController{
		Logger: logger,
	}

	userApp := mvc.New(app.Party("/users"))
	userApp.Register(userController)
	userApp.Handle(userController)

	app.Run(iris.Addr(":8080"))
}
