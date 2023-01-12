package routes

import "github.com/eqto/api-server"

func Logout(ctx api.Context) error {
	ctx.Response().Header().SetCookie(`SESS`, ``, 0)
	return nil
}