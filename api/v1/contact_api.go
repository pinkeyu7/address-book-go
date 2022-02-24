package v1

import (
	"address-book-go/api"
	"address-book-go/dto/apireq"
	contactRepo "address-book-go/internal/contact/repository"
	contactSrv "address-book-go/internal/contact/service"
	"address-book-go/pkg/er"
	"address-book-go/pkg/valider"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddContact
// @Summary Add Contact 新增意見回饋
// @Produce json
// @Accept json
// @Tags Contact
// @Security Bearer
// @Param Bearer header string true "JWT Token"
// @Param Body body apireq.AddContact true "Request Add Contact"
// @Success 200 {string} string "{}"
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 401 {object} er.AppErrorMsg "{"code":"400401","message":"Unauthorized"}"
// @Failure 403 {object} er.AppErrorMsg "{"code":"400403","message":"Permission denied"}"
// @Failure 404 {object} er.AppErrorMsg "{"code":"400404","message":"Resource not found"}"
// @Failure 500 {object} er.AppErrorMsg "{"code":"500000","message":"Database unknown error"}"
// @Router /v1/contacts [post]
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

	env := api.GetEnv()
	cr := contactRepo.NewRepository(env.Orm)
	cs := contactSrv.NewService(cr)
	err = cs.Add(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}

// ListContact
// @Summary List Contact 聯絡人列表
// @Produce json
// @Accept json
// @Tags Contact
// @Security Bearer
// @Param Bearer header string true "JWT Token"
// @Param page query string true "Page"
// @Param per_page query string true "Per Page"
// @Success 200 {object} apires.ListContact
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 401 {object} er.AppErrorMsg "{"code":"400401","message":"Unauthorized"}"
// @Failure 403 {object} er.AppErrorMsg "{"code":"400403","message":"Permission denied"}"
// @Failure 404 {object} er.AppErrorMsg "{"code":"400404","message":"Resource not found"}"
// @Failure 500 {object} er.AppErrorMsg "{"code":"500000","message":"Database unknown error"}"
// @Router /v1/contacts [get]
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

	env := api.GetEnv()
	cr := contactRepo.NewRepository(env.Orm)
	cs := contactSrv.NewService(cr)
	res, err := cs.List(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetContact
// @Summary Get Contact 取得聯絡人
// @Produce json
// @Accept json
// @Tags Contact
// @Security Bearer
// @Param Bearer header string true "JWT Token"
// @Param contact_id path int true "聯絡人id e.g. 11"
// @Success 200 {object} model.Contact
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 401 {object} er.AppErrorMsg "{"code":"400401","message":"Unauthorized"}"
// @Failure 403 {object} er.AppErrorMsg "{"code":"400403","message":"Permission denied"}"
// @Failure 404 {object} er.AppErrorMsg "{"code":"400404","message":"Resource not found"}"
// @Failure 500 {object} er.AppErrorMsg "{"code":"500000","message":"Database unknown error"}"
// @Router /v1/contacts/{contact_id} [get]
func GetContact(c *gin.Context) {
	contactIdStr := c.Param("id")
	contactId, err := strconv.Atoi(contactIdStr)
	if err != nil {
		paramErr := er.NewAppErr(http.StatusBadRequest, er.ErrorParamInvalid, "contact id format error.", err)
		_ = c.Error(paramErr)
		return
	}

	env := api.GetEnv()
	cr := contactRepo.NewRepository(env.Orm)
	cs := contactSrv.NewService(cr)
	res, err := cs.Get(contactId)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// EditContact
// @Summary Edit Contact 編輯聯絡人
// @Produce json
// @Accept json
// @Tags Contact
// @Security Bearer
// @Param Bearer header string true "JWT Token"
// @Param contact_id path int true "聯絡人id e.g. 11"
// @Param Body body apireq.EditContact true "Request Edit Contact"
// @Success 200 {string} string "{}"
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 401 {object} er.AppErrorMsg "{"code":"400401","message":"Unauthorized"}"
// @Failure 403 {object} er.AppErrorMsg "{"code":"400403","message":"Permission denied"}"
// @Failure 404 {object} er.AppErrorMsg "{"code":"400404","message":"Resource not found"}"
// @Failure 500 {object} er.AppErrorMsg "{"code":"500000","message":"Database unknown error"}"
// @Router /v1/contacts/{contact_id} [put]
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

	env := api.GetEnv()
	cr := contactRepo.NewRepository(env.Orm)
	cs := contactSrv.NewService(cr)
	err = cs.Edit(contactId, &req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}
