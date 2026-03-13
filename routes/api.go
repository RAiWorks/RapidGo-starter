package routes

import (
	"github.com/raiworks/rapidgo/v2/core/router"
	"github.com/raiworks/rapidgo-starter/http/controllers"
)

// RegisterAPI defines API routes under the /api prefix.
func RegisterAPI(r *router.Router) {
	api := r.Group("/api")
	api.APIResource("/posts", &controllers.PostController{})
}