package router

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashworks/relaxdays-hackathon-cc-vol1-7-backend-storage/models"
)

// API endpoint that returns a specific storage by name
// @Summary Returns a specifc storage by name
// @Success 200 {object} models.V1Storage
// @Failure 404 {} {} "Storage not found"
// @Param x query string true "Name of storage to retrieve"
// @Router /v1/storagePlace [get]
// @Tags V1Storage
func (s Server) V1StorageGet(c *gin.Context) {
	name := c.Query("x")
	storageRow, err := s.DotSelect.QueryRow(s.DB, "select-storage-by-name", name)

	if err != nil {
		s.internalServerError(c, err.Error())
		return
	}

	v0Storage := models.V0Storage{
		Name: name,
	}

	err = storageRow.Scan(&v0Storage.ArticleID, &v0Storage.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			c.Status(http.StatusNotFound)
			return
		}
		s.internalServerError(c, err.Error())
		return
	}

	valid, v1Storage := v0Storage.ToV1Storage()
	if !valid {
		s.internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, v1Storage)
}

// API endpoint that saves a storage
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
// @Summary Update a storage
// @Success 200
// @Failure 400 {} {} "Invalid storage"
// @Param storage body models.V1Storage true "Storage to Update"
// @Router /v1/storagePlace/ [post]
// @Tags V1Storage
func (s Server) V1StoragePost(c *gin.Context) {
	var storage models.V1Storage
	c.BindJSON(&storage)

	if !storage.IsValid() {
		c.Status(http.StatusBadRequest)
		return
	}

	_, err := s.DotAlter.Exec(s.DB, "insert-or-update-storage", storage.V1PropertiesToName(), storage.ArticleID, storage.Stock)
	if err != nil {
		s.internalServerError(c, err.Error())
		return
	}

	s.cacheStore.Flush()

	c.Status(http.StatusOK)
}

// API endpoint that deletes a saved storage by name
// @Summary Delete a storage by name
// @Success 204
// @Param x query string true "Name of storage to delete"
// @Router /v1/storagePlace [delete]
// @Tags V1Storage
func (s Server) V1StorageDeleteByName(c *gin.Context) {
	s.V0StorageDeleteByName(c)
}

// API endpoint returns n storages after a specific one
// @Summary Returns "n" storages lexicographically after storage "name"
// @Success 200 {array} models.V1Storage
// @Param n query int true "Number of storages after the named one"
// @Param x query string false "Storage name where we should start the cursor"
// @Router /v1/storagesPlaces [get]
// @Tags V1Storage
func (s Server) V1StorageGetCursor(c *gin.Context) {
	storageRows, err := s.DotSelect.Query(s.DB, "select-storage-by-cursor", c.Query("x"), c.Query("n"))
	defer storageRows.Close()

	if err != nil {
		s.internalServerError(c, err.Error())
		return
	}

	allStorages := make([]models.V1Storage, 0)

	for storageRows.Next() {
		var v0Storage models.V0Storage
		err := storageRows.Scan(&v0Storage.Name, &v0Storage.ArticleID, &v0Storage.Stock)
		if err != nil {
			s.internalServerError(c, err.Error())
			return
		}
		valid, v1Storage := v0Storage.ToV1Storage()
		if !valid {
			s.internalServerError(c, err.Error())
			return
		}
		allStorages = append(allStorages, v1Storage)
	}

	c.JSON(http.StatusOK, allStorages)
}
