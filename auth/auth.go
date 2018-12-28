package auth

import (
	"fmt"

	"github.com/pengkebao/ftpserver/conf"
)

type Auth struct {
}

func (this *Auth) CheckPasswd(user, pass string) (bool, error) {
	if v, ok := conf.Users[user]; ok {
		if pass == v.Password {
			return true, nil
		}
	}
	return false, fmt.Errorf("用户名或密码不正确")
}
