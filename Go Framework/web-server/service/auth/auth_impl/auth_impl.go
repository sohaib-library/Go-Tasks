package auth_service

import auth_repo "web-server/repo/auth/auth_impl"

type Auth struct {
	authRepo auth_repo.AuthImpl
}

func NewAuth(authRepo *auth_repo.AuthImpl) *Auth {
	return &Auth{authRepo: *authRepo}
}
