package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashworks/relaxdays-hackathon-cc-vol1-7-backend-storage/models"
)

// API endpoint that returns a specific storage by name
//
// @Summary Returns a specifc storage by name
// @Success 200 {object} models.V2Storage
// @Failure 404 {} {} "Storage not found"
// @Param x query string true "Name of storage to retrieve"
// @Router /v3/storagePlace [get]
// @Tags V3Storage
// @Security BasicAuth
func (s Server) V3StorageGet(c *gin.Context) {
	// Currently this function is redundant, it is only used for swagger annotations.
	s.V2StorageGet(c)
}

// API endpoint that saves a storage
//
// @Summary Save a storage
// @Success 200
// @Failure 400 {} {} "Invalid storage"
// @Param storage body models.V2Storage true "Storage to save"
// @Router /v3/storagePlace/ [put]
// @Tags V3Storage
// @Security BasicAuth
func (s Server) V3StoragePut(c *gin.Context) {
	// Currently this function is redundant, it is only used for swagger annotations.
	s.V2StoragePost(c)
}

// API endpoint that updates a storage
//
// @Summary Update a storage
// @Success 200
// @Failure 400 {} {} "Invalid storage"
// @Param storage body models.V2Storage true "Storage to Update"
// @Router /v3/storagePlace/ [post]
// @Tags V3Storage
// @Security BasicAuth
func (s Server) V3StoragePost(c *gin.Context) {
	// Currently this function is redundant, it is only used for swagger annotations.
	s.V2StoragePost(c)
}

// API endpoint that deletes a saved storage by name
//
// @Summary Delete a storage by name
// @Success 204
// @Param x query string true "Name of storage to delete"
// @Router /v3/storagePlace [delete]
// @Tags V3Storage
// @Security BasicAuth
func (s Server) V3StorageDeleteByName(c *gin.Context) {
	// Currently this function is redundant, it is only used for swagger annotations.
	s.V0StorageDeleteByName(c)
}

// API endpoint returns n storages after a specific one
//
// @Summary Returns "n" storages lexicographically after storage "name"
// @Success 200 {array} models.V2Storage
// @Param n query int true "Number of storages after the named one"
// @Param x query string false "Storage name where we should start the cursor"
// @Router /v3/storagesPlaces [get]
// @Tags V3Storage
// @Security BasicAuth
func (s Server) V3StorageGetCursor(c *gin.Context) {
	// Currently this function is redundant, it is only used for swagger annotations.
	s.V2StorageGetCursor(c)
}

// API endpoint returns n storages after a specific one
//
// @Summary Returns all storages that contain an article
// @Success 200 {array} models.V2Storage
// @Param x query int false "Article ID"
// @Router /v3/storagePlacesForArticleID [get]
// @Tags V3Storage
// @Security BasicAuth
func (s Server) V3StorageGetPlacesForArticleID(c *gin.Context) {
	// Currently this function is redundant, it is only used for swagger annotations.
	s.V2StorageGetPlacesForArticleID(c)
}

// API endpoint returns n storages after a specific one
//
// @Summary Returns "n" storages lexicographically after storage "x" at place "l"
// @Success 200 {array} models.V2Storage
// @Param n query int true "Number of storages after the named one"
// @Param l query string true "Place to query"
// @Param x query string false "Storage name where we should start the cursor"
// @Router /v3/storagePlacesAtLocation [get]
// @Tags V3Storage
// @Security BasicAuth
func (s Server) V3StorageGetPlacesAtLocationCursor(c *gin.Context) {
	code, v0Storages := s.v3GetStoragesByCursorAndPlace(c.Query("l"), c.Query("x"), c.Query("n"))

	if code == http.StatusOK {
		v2Storages := make([]models.V2Storage, 0)
		for _, v0Storage := range v0Storages {
			valid, v1Storage := v0Storage.ToV1Storage()
			if !valid {
				log.Print("Failed to convert v0Storage to v1Storage")
				c.Status(http.StatusInternalServerError)
				return
			}
			v2Storages = append(v2Storages, v1Storage.ToV2Storage())
		}
		c.JSON(code, v2Storages)
	} else {
		c.Status(code)
	}
}
