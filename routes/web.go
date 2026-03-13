package routes

import (
	"github.com/raiworks/rapidgo/v2/core/router"
	"github.com/raiworks/rapidgo-starter/http/controllers"
)

// RegisterWeb defines web (HTML) routes.
func RegisterWeb(r *router.Router) {
	r.Get("/", controllers.Home)
}