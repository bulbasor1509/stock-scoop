package config

import (
	"github.com/gin-contrib/cors"
	"time"
)

func CorsConfig(cfg *cors.Config) {
	// Allow only localhost:3000
	cfg.AllowOrigins = []string{"http://localhost:3000"}

	// Other settings can remain as you want
	cfg.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	cfg.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	cfg.ExposeHeaders = []string{"Content-Length"}
	cfg.AllowCredentials = true
	cfg.MaxAge = 12 * time.Hour

}
