package router

import (
	"database/sql"
	"fmt"
	"net/http"
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
		fmt.Printf("DeprecatedCall@CC-VOL7: %s %s %s %s X-Forwarded-For=%s\n",
			c.ClientIP(),
			time.Now().Format("02/Jun/2006:15:04:05"),
			c.Request.Method,
			c.Request.URL.String(),
			c.Request.Header.Get("X-Forwarded-For"),
		)
	}
}

// NewRouter returns a new router.
func (s Server) NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORS())

	s.cacheStore = persistence.NewInMemoryStore(time.Second)

	// v0
	v0 := router.Group("/")
	v0.Use(logOldAPIRequest())
	v0.DELETE("/storagePlace", s.V0StorageDeleteByName)
	v0.GET("/storagePlace", cache.CachePage(s.cacheStore, time.Hour, s.V0StorageGet))
	v0.PUT("/storagePlace", s.V0StoragePut)
	v0.POST("/storagePlace", s.V0StoragePost)
	v0.GET("/storagesPlaces", cache.CachePage(s.cacheStore, time.Hour, s.V0StorageGetCursor))

	// v1
	v1 := router.Group("/v1")
	v1.Use(logOldAPIRequest())
	v1.DELETE("/storagePlace", s.V1StorageDeleteByName)
	v1.GET("/storagePlace", cache.CachePage(s.cacheStore, time.Hour, s.V1StorageGet))
	v1.PUT("/storagePlace", s.V1StoragePut)
	v1.POST("/storagePlace", s.V1StoragePost)
	v1.GET("/storagesPlaces", cache.CachePage(s.cacheStore, time.Hour, s.V1StorageGetCursor))

	// v2
	v2 := router.Group("/v2")
	v1.Use(logOldAPIRequest())
	v2.DELETE("/storagePlace", s.V2StorageDeleteByName)
	v2.GET("/storagePlace", cache.CachePage(s.cacheStore, time.Hour, s.V2StorageGet))
	v2.PUT("/storagePlace", s.V2StoragePut)
	v2.POST("/storagePlace", s.V2StoragePost)
	v2.GET("/storagesPlaces", cache.CachePage(s.cacheStore, time.Hour, s.V2StorageGetCursor))
	v2.GET("/storagePlacesForArticleID", cache.CachePage(s.cacheStore, time.Hour, s.V2StorageGetPlacesForArticleID))

	// v3
	authorizedV3 := router.Group("/v3", gin.BasicAuth(gin.Accounts{
		"user": "pass",
	}))
	authorizedV3.DELETE("/storagePlace", s.V3StorageDeleteByName)
	authorizedV3.GET("/storagePlace", cache.CachePage(s.cacheStore, time.Hour, s.V3StorageGet))
	authorizedV3.PUT("/storagePlace", s.V3StoragePut)
	authorizedV3.POST("/storagePlace", s.V3StoragePost)
	authorizedV3.GET("/storagesPlaces", cache.CachePage(s.cacheStore, time.Hour, s.V3StorageGetCursor))
	authorizedV3.GET("/storagePlacesForArticleID", cache.CachePage(s.cacheStore, time.Hour, s.V3StorageGetPlacesForArticleID))

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	openapiURL := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, openapiURL))

	return router
}
