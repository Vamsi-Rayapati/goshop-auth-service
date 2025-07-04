package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smartbot/auth/pkg/validator"
)

type AuthController struct {
	service AuthService
}

func (ac AuthController) Signup(c *gin.Context) {
	var user SignupRequest

	err := validator.ValidateBody(c, &user)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, signupErr := ac.service.Signup(&user)

	if signupErr != nil {
		c.JSON(signupErr.Code, signupErr)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (ac AuthController) Login(c *gin.Context) {

	var user LoginRequest
	err := validator.ValidateBody(c, &user)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, loginErr := ac.service.Login(&user)

	if loginErr != nil {
		c.JSON(loginErr.Code, loginErr)
		return
	}

	c.JSON(http.StatusCreated, res)

}

func (ac AuthController) RefreshToken(c *gin.Context) {
	var authInfo RefreshRequest
	err := validator.ValidateBody(c, &authInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, refreshTokenErr := ac.service.RefreshToken(authInfo)
	if refreshTokenErr != nil {
		c.JSON(refreshTokenErr.Code, refreshTokenErr)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (ac AuthController) GetJWKS(c *gin.Context) {
	jwks, jwksErr := ac.service.FetchJWKS()

	if jwksErr != nil {
		c.JSON(jwksErr.Code, jwksErr)
		return
	}

	c.JSON(http.StatusOK, jwks)
}
