package auth_controller

import (
	"emvn/consts"
	"emvn/internal/model"
	auth_usecase "emvn/internal/usecase/auth"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}

type authController struct {
	authUsecase auth_usecase.IAuthUsecase
}

func NewController(authUsecase auth_usecase.IAuthUsecase) IAuthController {
	return &authController{
		authUsecase: authUsecase,
	}
}

// SignUp godoc
//
//	@Summary		Sign up new user
//	@Description	Sign up new user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		SignUpInput	true	"Sign up new user"
//	@Success		200		{object}	SignUpOutput
//	@Router			/auth/signup [post]
func (ctrl *authController) SignUp(c *gin.Context) {
	// validate request
	var in SignUpInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.Set(consts.GinErrorKey, consts.CodeInvalidRequest)
		c.Set(consts.GinDetailErrorKey, err)
		return
	}
	// call usecase

	err := ctrl.authUsecase.SignUp(c, model.User{
		Username: in.Username,
		Password: in.Password,
	})
	if err != nil {
		c.Set(consts.GinErrorKey, err)
		return
	}
	// response
	c.Set(consts.GinResponseKey, SignUpOutput{
		Success: true,
	})
}

// SignIn godoc
//
//	@Summary		Sign in user
//	@Description	Sign in user, return token and exp time
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		SignInInput	true	"Sign in user"
//	@Success		200		{object}	SignInOutput
//	@Router			/auth/signin [post]
func (ctrl *authController) SignIn(c *gin.Context) {
	// validate request
	var in SignInInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.Set(consts.GinErrorKey, consts.CodeInvalidRequest)
		c.Set(consts.GinDetailErrorKey, err)
		return
	}
	// call usecase

	out, err := ctrl.authUsecase.SignIn(c, in.Username, in.Password)
	if err != nil {
		c.Set(consts.GinErrorKey, err)
		return
	}
	// response
	c.Set(consts.GinResponseKey, SignInOutput{
		Token: out.Token,
		Exp:   out.Exp,
	})
}
