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
// @Success 200 {object} models.V1Storage
// @Failure 404 {} {} "Storage not found"
// @Param x query string true "Name of storage to retrieve"
// @Router /v1/storagePlace [get]
// @Tags V1Storage
func (s Server) V1StorageGet(c *gin.Context) {
	code, v0Storage := s.v0SelectStorageByName(c.Query("x"))

	if code == http.StatusOK {
		valid, v1Storage := v0Storage.ToV1Storage()
		if !valid {
			log.Print("Failed to convert v0Storage to v1Storage")
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(code, v1Storage)
	} else {
		c.Status(code)
	}
}

// API endpoint that saves a storage
//
// @Summary Save a storage
// @Success 200
// @Failure 400 {} {} "Invalid storage"
// @Param storage body models.V1Storage true "Storage to save"
// @Router /v1/storagePlace/ [put]
// @Tags V1Storage
func (s Server) V1StoragePut(c *gin.Context) {
	// Currently this function is redundant, it is only used for swagger annotations.
	s.V1StoragePost(c)
}

// API endpoint that updates a storage
//
// @Summary Update a storage
// @Success 200
// @Failure 400 {} {} "Invalid storage"
// @Param storage body models.V1Storage true "Storage to Update"
// @Router /v1/storagePlace/ [post]
// @Tags V1Storage
func (s Server) V1StoragePost(c *gin.Context) {
	var v1Storage models.V1Storage
	c.BindJSON(&v1Storage)

	// Bei Aufrufen an legacy Endpunkte soll diese den Wert vom Feld "bestand" annehmen.
	v1Storage.Capacity = v1Storage.Stock

	if !v1Storage.IsValid() {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(s.v0InsertOrUpdateStorage(v1Storage.ToV0Storage()))
}

// API endpoint that deletes a saved storage by name
//
// @Summary Delete a storage by name
// @Success 204
// @Param x query string true "Name of storage to delete"
// @Router /v1/storagePlace [delete]
// @Tags V1Storage
func (s Server) V1StorageDeleteByName(c *gin.Context) {
	s.V0StorageDeleteByName(c)
}

// API endpoint returns n storages after a specific one
//
// @Summary Returns "n" storages lexicographically after storage "name"
// @Success 200 {array} models.V1Storage
// @Param n query int true "Number of storages after the named one"
// @Param x query string false "Storage name where we should start the cursor"
// @Router /v1/storagesPlaces [get]
// @Tags V1Storage
func (s Server) V1StorageGetCursor(c *gin.Context) {
	code, v0Storages := s.v0GetStoragesByCursor(c.Query("x"), c.Query("n"))

	if code == http.StatusOK {
		v1Storages := make([]models.V1Storage, 0)
		for _, v0Storage := range v0Storages {
			valid, v1Storage := v0Storage.ToV1Storage()
			if !valid {
				log.Print("Failed to convert v0Storage to v1Storage")
				c.Status(http.StatusInternalServerError)
				return
			}
			v1Storages = append(v1Storages, v1Storage)
		}
		c.JSON(code, v1Storages)
	} else {
		c.Status(code)
	}
}
