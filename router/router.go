package router

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gchaincl/dotsql"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Server struct {
	DotAlter   *dotsql.DotSql
	DotSelect  *dotsql.DotSql
	DB         *sql.DB
	router     *gin.Engine
	cacheStore *persistence.InMemoryStore
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func logOldAPIRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.String()
		if !strings.HasPrefix(url, "/v2/") {
			fmt.Printf("DeprecatedCall@CC-VOL7: %s %s %s %s X-Forwarded-For=%s\n",
				c.ClientIP(),
				time.Now().Format("02/Jun/2006:15:04:05"),
				c.Request.Method,
				url,
				c.Request.Header.Get("X-Forwarded-For"),
			)
		}
	}
}

// NewRouter returns a new router.
func (s Server) NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORS())
	router.Use(logOldAPIRequest())

	s.cacheStore = persistence.NewInMemoryStore(time.Second)

	// v0
	router.DELETE("/storagePlace", s.V0StorageDeleteByName)
	router.GET("/storagePlace", cache.CachePage(s.cacheStore, time.Hour, s.V0StorageGet))
	router.PUT("/storagePlace", s.V0StoragePut)
	router.POST("/storagePlace", s.V0StoragePost)
	router.GET("/storagesPlaces", cache.CachePage(s.cacheStore, time.Hour, s.V0StorageGetCursor))

	// v1
	router.DELETE("/v1/storagePlace", s.V1StorageDeleteByName)
	router.GET("/v1/storagePlace", cache.CachePage(s.cacheStore, time.Hour, s.V1StorageGet))
	router.PUT("/v1/storagePlace", s.V1StoragePut)
	router.POST("/v1/storagePlace", s.V1StoragePost)
	router.GET("/v1/storagesPlaces", cache.CachePage(s.cacheStore, time.Hour, s.V1StorageGetCursor))

	// v2
	router.DELETE("/v2/storagePlace", s.V2StorageDeleteByName)
	router.GET("/v2/storagePlace", cache.CachePage(s.cacheStore, time.Hour, s.V2StorageGet))
	router.PUT("/v2/storagePlace", s.V2StoragePut)
	router.POST("/v2/storagePlace", s.V2StoragePost)
	router.GET("/v2/storagesPlaces", cache.CachePage(s.cacheStore, time.Hour, s.V2StorageGetCursor))
	router.GET("/v2/storagePlacesForArticleID", cache.CachePage(s.cacheStore, time.Hour, s.V2StorageGetPlacesForArticleID))

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	openapiURL := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, openapiURL))

	return router
}
