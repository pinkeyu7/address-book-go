package service

import (
	"address-book-go/dto/apireq"
	"address-book-go/dto/apires"
	"address-book-go/dto/model"
	"address-book-go/internal/system/sys_account"
	"address-book-go/internal/token"
	tokenLibrary "address-book-go/internal/token/library"
	"address-book-go/pkg/er"
	"address-book-go/pkg/helper"
	"net/http"
)

type Service struct {
	sysAccRepo sys_account.Repository
}

func NewService(sar sys_account.Repository) token.Service {
	return &Service{
		sysAccRepo: sar,
	}
}

func (s *Service) GenToken(req *apireq.GetSysAccountToken) (*apires.SysAccountToken, error) {
	// Check Account Exist
	acc, err := s.sysAccRepo.FindOne(&model.SysAccount{Account: req.Account})
	if err != nil {
		findErr := er.NewAppErr(http.StatusInternalServerError, er.UnknownError, "find account error.", err)
		return nil, findErr
	}
	if acc == nil || acc.IsDisable {
		authErr := er.NewAppErr(http.StatusUnauthorized, er.UnauthorizedError, "", nil)
		return nil, authErr
	}

	// Password not matched
	pw := helper.ScryptStr(req.Password)
	if acc.Password != pw {
		authErr := er.NewAppErr(http.StatusUnauthorized, er.UnauthorizedError, "", nil)
		return nil, authErr
	}

	oToken, expiredAt, err := tokenLibrary.GenToken(acc.Id)
	if err != nil {
		tokenErr := er.NewAppErr(http.StatusUnauthorized, er.UnauthorizedError, "", err)
		return nil, tokenErr
	}

	// TODO 補後踢前
	//_ = tokenCache.SetCodeCertTokenSession(account.Id, iat)

	mapData := map[string]interface{}{}
	mapData["name"] = acc.Name
	mapData["email"] = acc.Email

	res := apires.SysAccountToken{
		Token:     oToken,
		ExpiredAt: expiredAt,
		Data:      mapData,
	}

	return &res, nil
}
