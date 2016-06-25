package routes

import (
	
	"net/http"
	"github.com/gin-gonic/gin"
)


type Router struct {

	 ginrouter *gin.Engine
	 routergroup *gin.RouterGroup
}


func NewRouter() *Router {
	
	r := new(Router)
	r.ginrouter = gin.Default()
	r.routergroup = r.ginrouter.Group("api/v1")
	{
		r.routergroup.GET("/:id", GetCity)
		r.routergroup.DELETE("/", notImplemented)
		r.routergroup.POST("/", notImplemented)
		r.routergroup.PUT("/", notImplemented)
	}
	
	r.ginrouter.Run(":8000")
	return r;
}


func notImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error" : "method not Implemented"})
}

func GetCity( c *gin.Context){
	 city := c.Param("id")
	 c.String(http.StatusOK, "found :", city)
}




