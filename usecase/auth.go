package usecase

import "errors"

// AuthUsecase は認証のユースケース。
type AuthUsecase interface {
	ValidatePassword(inputPassword string) error
}

// AuthInteractor は AuthUsecase の実装。
// ここでは「パスワード照合」というビジネスルールだけを持ち、設定取得は外側で注入する。
type AuthInteractor struct {
	password string
}

func NewAuthInteractor(password string) *AuthInteractor {
	return &AuthInteractor{password: password}
}

func (u *AuthInteractor) ValidatePassword(inputPassword string) error {
	if inputPassword != u.password {
		return errors.New("invalid password")
	}
	return nil
}
