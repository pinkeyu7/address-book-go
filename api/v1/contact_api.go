package v1

import (
	"address-book-go/dto/apireq"
	contactRepo "address-book-go/internal/contact/repository"
	contactSrv "address-book-go/internal/contact/service"
	"address-book-go/pkg/er"
	"address-book-go/pkg/valider"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddContact(c *gin.Context) {
	req := apireq.AddContact{}
	err := c.BindJSON(&req)
	if err != nil {
		paramErr := er.NewAppErr(http.StatusBadRequest, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(paramErr)
		return
	}

	// 參數驗證
	err = valider.Validate.Struct(req)
	if err != nil {
		paramErr := er.NewAppErr(http.StatusBadRequest, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(paramErr)
		return
	}

	cr := contactRepo.NewRepository()
	cs := contactSrv.NewService(cr)
	err = cs.Add(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}

func ListContact(c *gin.Context) {
	req := apireq.ListContact{}
	err := c.Bind(&req)
	if err != nil {
		paramErr := er.NewAppErr(http.StatusBadRequest, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(paramErr)
		return
	}

	// 參數驗證
	err = valider.Validate.Struct(req)
	if err != nil {
		paramErr := er.NewAppErr(http.StatusBadRequest, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(paramErr)
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

func GetContact(c *gin.Context) {
	contactIdStr := c.Param("id")
	contactId, err := strconv.Atoi(contactIdStr)
	if err != nil {
		paramErr := er.NewAppErr(http.StatusBadRequest, er.ErrorParamInvalid, "contact id format error.", err)
		_ = c.Error(paramErr)
		return
	}

	cr := contactRepo.NewRepository()
	cs := contactSrv.NewService(cr)
	res, err := cs.Get(contactId)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func EditContact(c *gin.Context) {
	contactIdStr := c.Param("id")
	contactId, err := strconv.Atoi(contactIdStr)
	if err != nil {
		paramErr := er.NewAppErr(http.StatusBadRequest, er.ErrorParamInvalid, "contact id format error.", err)
		_ = c.Error(paramErr)
		return
	}

	req := apireq.EditContact{}
	err = c.BindJSON(&req)
	if err != nil {
		paramErr := er.NewAppErr(http.StatusBadRequest, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(paramErr)
		return
	}

	// 參數驗證
	err = valider.Validate.Struct(req)
	if err != nil {
		paramErr := er.NewAppErr(http.StatusBadRequest, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(paramErr)
		return
	}

	cr := contactRepo.NewRepository()
	cs := contactSrv.NewService(cr)
	err = cs.Edit(contactId, &req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}
