// router

package router

import (
	"fmt"
	"sync"
	"time"

	"server/internal/user"
	"server/internal/ws"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	r           *gin.Engine
	activeUsers = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "active_users",
		Help: "Number of active users",
	})
	oldActiveUsers = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "old_active_users",
		Help: "Old count of active users",
	})
	customRegistry   = prometheus.NewRegistry()
	activeUsersCount float64    // New variable to keep the count
	mu               sync.Mutex // Mutex to protect the count variable
)

func refreshCounterEveryTenSeconds() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		mu.Lock()
		oldActiveUsers.Set(activeUsersCount)
		activeUsersCount = 0
		activeUsers.Set(0)
		mu.Unlock()
	}
}

func init() {
	customRegistry.MustRegister(activeUsers)
	customRegistry.MustRegister(oldActiveUsers)
	go refreshCounterEveryTenSeconds()
}

func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "OK")
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigins := []string{"http://devopschat.xyz", "http://devopschat.xyz/"}
		origin := c.Request.Header.Get("Origin")
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
				break
			}
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			for name, values := range c.Request.Header {
				// Loop over all values for the name.
				for _, value := range values {
					fmt.Printf("Header field %q, Value %q\n", name, value)
				}
			}
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func prometheusActiveUsersMetric() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Exclude /metrics endpoint
		if c.Request.URL.Path != "/metrics" {
			mu.Lock()
			activeUsersCount++
			activeUsers.Inc()
			mu.Unlock()

			c.Next()
		} else {
			c.Next()
		}
	}
}

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()
	//r.Use(Logger())
	r.Use(CORSMiddleware())

	r.Use(prometheusActiveUsersMetric())

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("/ws/createRoom", wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/getRooms", wsHandler.GetRooms)
	r.GET("/ws/getClients/:roomId", wsHandler.GetClients)

	// Add the health check endpoint
	r.GET("/health", HealthCheck())
	r.GET("/", HealthCheck())

	// prometheus metrics
	r.GET("/metrics", gin.WrapH(promhttp.HandlerFor(customRegistry, promhttp.HandlerOpts{})))
}

func Start(addr string) error {
	return r.Run(addr)
}

// code samples
/*func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		t := time.Now()

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(t)

		// Log request details
		log.Printf("Request: %s %s, status: %d, latency: %s\n", c.Request.Method, c.Request.URL, c.Writer.Status(), latency)
	}
}*/

/*r.Use(cors.New(cors.Config{
	AllowOrigins:     []string{"http://devopschat.xyz"},
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
	OptionsResponseStatusCode: 204,
	AllowOriginFunc: func(origin string) bool {
		return origin == "http://devopschat.xyz"
	},
	MaxAge: 12 * time.Hour,
}))*/
