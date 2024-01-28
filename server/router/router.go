// router
package router

import (
	"server/internal/user"
	"server/internal/ws"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	r.Use(cors.New(cors.Config{

		AllowAllOrigins: true,

		//AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Origin", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// Allows usage of popular browser extensions schemas
		AllowBrowserExtensions: true,

		// Allows usage of WebSocket protocol
		AllowWebSockets: true,

		// Allows usage of file:// schema (dangerous!) use it only when you 100% sure it's needed
		AllowFiles: true,

		// Allows to pass custom OPTIONS response status code for old browsers / clients
		OptionsResponseStatusCode: 200,

		/*AllowOriginFunc: func(origin string) bool {
			return true
			//return origin == os.Getenv("FRONTEND_URL")
		},*/
		//MaxAge: 12 * time.Hour,
	}))

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("/ws/createRoom", wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/getRooms", wsHandler.GetRooms)
	r.GET("/ws/getClients/:roomId", wsHandler.GetClients)
}

func Start(addr string) error {
	return r.Run(addr)
}
