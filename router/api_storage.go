package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashworks/relaxdays-hackathon-cc-vol1-7-backend-storage/models"
)

// @Summary Returns all storages
// @Success 200 {array} models.Storage
// @Router /storage/ [get]
// @Tags Storage
func (s Server) StorageGet(c *gin.Context) {
	storageRows, err := s.DotSelect.Query(s.DB, "select-storage")
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

// @Summary Save or update a storage
// @Success 200
// @Failure 400 {} {} "Invalid storage"
// @Param storage body models.Storage true "Storage to save"
// @Router /storage/ [put]
// @Tags Storage
func (s Server) StoragePut(c *gin.Context) {
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
// @Router /storage/{name} [delete]
// @Tags Storage
func (s Server) StorageDeleteByName(c *gin.Context) {
	_, err := s.DotAlter.Exec(s.DB, "delete-storage", c.Param("name"))
	if err != nil {
		s.internalServerError(c, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
