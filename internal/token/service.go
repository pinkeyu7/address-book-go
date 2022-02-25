package token

import (
	"address-book-go/dto/apireq"
	"address-book-go/dto/apires"
)

type Service interface {
	GenToken(req *apireq.GetSysAccountToken) (*apires.SysAccountToken, error)
}
