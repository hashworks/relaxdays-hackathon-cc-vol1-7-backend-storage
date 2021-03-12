package router

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashworks/relaxdays-hackathon-cc-vol1-7-backend-storage/models"
)

// API endpoint that returns a specific storage by name
// @Summary Returns a specifc storage by name
// @Success 200 {object} models.Storage
// @Failure 404 {} {} "Storage not found"
// @Param name path string true "Name of storage to retrieve"
// @Router /storagePlace/{name} [get]
// @Tags Storage
func (s Server) StorageGet(c *gin.Context) {
	name := c.Param("name")
	storageRow, err := s.DotSelect.QueryRow(s.DB, "select-storage-by-name", name)

	if err != nil {
		s.internalServerError(c, err.Error())
		return
	}

	storage := models.Storage{
		Name: name,
	}

	err = storageRow.Scan(&storage.ArticleID, &storage.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			c.Status(http.StatusNotFound)
			return
		}
		s.internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, storage)
}

// API endpoint that saves a storage
// @Summary Save a storage
// @Success 200
// @Failure 400 {} {} "Invalid storage"
// @Param storage body models.Storage true "Storage to save"
// @Router /storagePlace/ [put]
// @Tags Storage
func (s Server) StoragePut(c *gin.Context) {
	// Currently this function is redundant, it is only used for swagger annotations.
	s.StoragePost(c)
}

// API endpoint that updates a storage
// @Summary Update a storage
// @Success 200
// @Failure 400 {} {} "Invalid storage"
// @Param storage body models.Storage true "Storage to Update"
// @Router /storagePlace/ [post]
// @Tags Storage
func (s Server) StoragePost(c *gin.Context) {
	var storage models.Storage
	c.BindJSON(&storage)

	if !storage.IsValid() {
		c.Status(http.StatusBadRequest)
		return
	}

	_, err := s.DotAlter.Exec(s.DB, "insert-or-update-storage", storage.Name, storage.ArticleID, storage.Stock)
	if err != nil {
		s.internalServerError(c, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// API endpoint that deletes a saved storage by name
// @Summary Delete a storage by name
// @Success 204
// @Param name path string true "Name of storage to delete"
// @Router /storagePlace/{name} [delete]
// @Tags Storage
func (s Server) StorageDeleteByName(c *gin.Context) {
	_, err := s.DotAlter.Exec(s.DB, "delete-storage", c.Param("name"))
	if err != nil {
		s.internalServerError(c, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

// API endpoint returns n storages after a specific one
// @Summary Returns "n" storages lexicographically after storage "name"
// @Success 200 {array} models.Storage
// @Param n path int true "Number of storages after the named one"
// @Param name path string true "Storage name where we should start the cursor"
// @Router /storagesPlaces/{n}/{name} [get]
// @Tags Storage
func (s Server) StorageGetCursor(c *gin.Context) {
	storageRows, err := s.DotSelect.Query(s.DB, "select-storage-by-cursor", c.Param("name"), c.Param("n"))
	defer storageRows.Close()

	if err != nil {
		s.internalServerError(c, err.Error())
		return
	}

	allStorages := make([]models.Storage, 0)

	for storageRows.Next() {
		var storage models.Storage
		err := storageRows.Scan(&storage.Name, &storage.ArticleID, &storage.Stock)
		if err != nil {
			s.internalServerError(c, err.Error())
			return
		}
		allStorages = append(allStorages, storage)
	}

	c.JSON(http.StatusOK, allStorages)
}
