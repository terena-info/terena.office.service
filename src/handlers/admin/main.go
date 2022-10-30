package adminHandle

import (
	"github.com/gin-gonic/gin"
	adminRepo "terena.office/src/repositories/admin"
)

type _Handlers interface {
	FindById(*gin.Context)
}

type _Adator struct {
	adminRepo adminRepo.AdminRepositories
}

func New() _Handlers {
	var handle _Handlers = _Adator{
		adminRepo: adminRepo.New(),
	}
	return handle
}
