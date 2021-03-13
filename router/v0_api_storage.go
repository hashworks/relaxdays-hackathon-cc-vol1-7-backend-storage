package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashworks/relaxdays-hackathon-cc-vol1-7-backend-storage/models"
)

// API endpoint that returns a specific storage by name
//
// @Summary Returns a specifc storage by name
// @Success 200 {object} models.V0Storage
// @Failure 404 {} {} "Storage not found"
// @Param x query string true "Name of storage to retrieve"
// @Router /storagePlace [get]
// @Tags V0Storage
func (s Server) V0StorageGet(c *gin.Context) {
	code, storage := s.v0SelectStorageByName(c.Query("x"))

	if code == http.StatusOK {
		c.JSON(code, storage)
	} else {
		c.Status(code)
	}
}

// API endpoint that saves a storage
//
// @Summary Save a storage
// @Success 204
// @Failure 400 {} {} "Invalid storage"
// @Param storage body models.V0Storage true "Storage to save"
// @Router /storagePlace/ [put]
// @Tags V0Storage
func (s Server) V0StoragePut(c *gin.Context) {
	// Currently this function is redundant, it is only used for swagger annotations.
	s.V0StoragePost(c)
}

// API endpoint that updates a storage
//
// @Summary Update a storage
// @Success 204
// @Failure 400 {} {} "Invalid storage"
// @Param storage body models.V0Storage true "Storage to Update"
// @Router /storagePlace/ [post]
// @Tags V0Storage
func (s Server) V0StoragePost(c *gin.Context) {
	var storage models.V0Storage
	c.BindJSON(&storage)

	// Bei Aufrufen an legacy Endpunkte soll diese den Wert vom Feld "bestand" annehmen.
	storage.Capacity = storage.Stock

	if !storage.IsValid() {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(s.v0InsertOrUpdateStorage(storage))
}

// API endpoint that deletes a saved storage by name
//
// @Summary Delete a storage by name
// @Success 204
// @Param x query string true "Name of storage to delete"
// @Router /storagePlace [delete]
// @Tags V0Storage
func (s Server) V0StorageDeleteByName(c *gin.Context) {
	c.Status(s.v0DeleteStorageByName(c.Query("x")))
}

// API endpoint returns n storages after a specific one
//
// @Summary Returns "n" storages lexicographically after storage "name"
// @Success 200 {array} models.V0Storage
// @Param n query int true "Number of storages after the named one"
// @Param x query string false "Storage name where we should start the cursor"
// @Router /storagesPlaces [get]
// @Tags V0Storage
func (s Server) V0StorageGetCursor(c *gin.Context) {
	code, storages := s.v0GetStoragesByCursor(c.Query("x"), c.Query("n"))

	if code == http.StatusOK {
		c.JSON(code, storages)
	} else {
		c.Status(code)
	}
}
