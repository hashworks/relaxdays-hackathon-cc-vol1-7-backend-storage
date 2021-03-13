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
// @Router /v2/storagePlace [get]
// @Tags V2Storage
func (s Server) V2StorageGet(c *gin.Context) {
	code, v0Storage := s.v0SelectStorageByName(c.Query("x"))

	if code == http.StatusOK {
		valid, v1Storage := v0Storage.ToV1Storage()
		if !valid {
			log.Print("Failed to convert v0Storage to v1Storage")
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(code, v1Storage.ToV2Storage())
	} else {
		c.Status(code)
	}
}

// API endpoint that saves a storage
//
// @Summary Save a storage
// @Success 200
// @Failure 400 {} {} "Invalid storage"
// @Param storage body models.V2Storage true "Storage to save"
// @Router /v2/storagePlace/ [put]
// @Tags V2Storage
func (s Server) V2StoragePut(c *gin.Context) {
	// Currently this function is redundant, it is only used for swagger annotations.
	s.V2StoragePost(c)
}

// API endpoint that updates a storage
//
// @Summary Update a storage
// @Success 200
// @Failure 400 {} {} "Invalid storage"
// @Param storage body models.V2Storage true "Storage to Update"
// @Router /v2/storagePlace/ [post]
// @Tags V2Storage
func (s Server) V2StoragePost(c *gin.Context) {
	var v2Storage models.V2Storage
	c.BindJSON(&v2Storage)

	if !v2Storage.IsValid() {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(s.v0InsertOrUpdateStorage(v2Storage.ToV1Storage().ToV0Storage()))
}

// API endpoint that deletes a saved storage by name
//
// @Summary Delete a storage by name
// @Success 204
// @Param x query string true "Name of storage to delete"
// @Router /v2/storagePlace [delete]
// @Tags V2Storage
func (s Server) V2StorageDeleteByName(c *gin.Context) {
	s.V0StorageDeleteByName(c)
}

// API endpoint returns n storages after a specific one
//
// @Summary Returns "n" storages lexicographically after storage "name"
// @Success 200 {array} models.V2Storage
// @Param n query int true "Number of storages after the named one"
// @Param x query string false "Storage name where we should start the cursor"
// @Router /v2/storagesPlaces [get]
// @Tags V2Storage
func (s Server) V2StorageGetCursor(c *gin.Context) {
	code, v0Storages := s.v0GetStoragesByCursor(c.Query("x"), c.Query("n"))

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

// API endpoint returns n storages after a specific one
//
// @Summary Returns all storages that contain an article
// @Success 200 {array} models.V2Storage
// @Param x query int false "Article ID"
// @Router /v2/storagePlacesForArticleID [get]
// @Tags V2Storage
func (s Server) V2StorageGetPlacesForArticleID(c *gin.Context) {
	code, v2Storages := s.v2GetStoragesByArticleID(c.Query("x"))

	if code == http.StatusOK {
		c.JSON(code, v2Storages)
	} else {
		c.Status(code)
	}
}
