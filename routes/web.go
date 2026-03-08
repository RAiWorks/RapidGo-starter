package routes

import (
	"github.com/RAiWorks/RapidGo/v2/core/router"
	"github.com/RAiWorks/RapidGo-starter/http/controllers"
)

// RegisterWeb defines web (HTML) routes.
func RegisterWeb(r *router.Router) {
	r.Get("/", controllers.Home)
}