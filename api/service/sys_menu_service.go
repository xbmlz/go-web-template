package service

type SysMenuService interface{}

type sysMenuService struct{}

func NewSysMenuService() SysMenuService {
	return &sysMenuService{}
}
