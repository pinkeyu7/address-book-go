package v1

import (
	"address-book-go/dto/apireq"
	contactRepo "address-book-go/internal/contact/repository"
	contactSrv "address-book-go/internal/contact/service"
	"address-book-go/pkg/er"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListContact(c *gin.Context) {
	req := apireq.ListContact{}
	err := c.Bind(&req)
	if err != nil {
		err = er.NewAppErr(http.StatusBadRequest, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(err)
		return
	}

	cr := contactRepo.NewRepository()
	cs := contactSrv.NewService(cr)
	res, err := cs.List(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, res)
}
