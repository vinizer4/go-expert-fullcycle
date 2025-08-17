package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"vinizer4/go-expert-fullcycle/labs/auction/configuration/rest_err"
	"vinizer4/go-expert-fullcycle/labs/auction/internal/usecase/user_usecase"
)

type userController struct {
	userUserCase user_usecase.UserUseCase
}

func NewUserController(userUserCase user_usecase.UserUseCase) *userController {
	return &userController{
		userUserCase: userUserCase,
	}
}

func (u *userController) FindUserById(c *gin.Context) {
	userId := c.Param("userId")

	if err := uuid.Validate(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "userId",
			Message: "Invalid UUID value",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	userData, err := u.userUserCase.FindUserById(c.Request.Context(), userId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(200, userData)
}
