package globalRouters

import (
	"github.com/gin-gonic/gin"
	globalControllers "github.com/putragabrielll/fwg17-cinematix-be/src/controllers/global"
)

func MovieRouter(r *gin.RouterGroup) {
	r.GET("", globalControllers.ListAllMovies)
	r.GET("/:id", globalControllers.DetailMovie)
}
