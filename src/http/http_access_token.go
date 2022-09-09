package http

import (
	"net/http"

	atDomain "github.com/PreetSIngh8929/bookstore_oauth-api/src/domain/access_token"
	"github.com/PreetSIngh8929/bookstore_oauth-api/src/services/access_token"
	"github.com/PreetSIngh8929/boookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{service: service}
}

func (handler *accessTokenHandler) GetByID(c *gin.Context) {

	accessToken, err := handler.service.GetById(c.Param("access_token_id"))

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, accessToken)

}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("inavlid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	accessToken, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}