package router

import (
	"log"
	"net/http"

	"github.com/hashworks/relaxdays-hackathon-cc-vol1-7-backend-storage/models"
)

func (s Server) v3GetStoragesByCursorAndPlace(l string, x string, n string) (int, []models.V0Storage) {
	allStorages := make([]models.V0Storage, 0)

	storageRows, err := s.DotSelect.Query(s.DB, "select-storage-by-cursor", l, x, n)
	defer storageRows.Close()

	if err != nil {
		log.Print(err.Error())
		return http.StatusInternalServerError, allStorages
	}

	for storageRows.Next() {
		var storage models.V0Storage
		err := storageRows.Scan(&storage.Name, &storage.ArticleID, &storage.Stock, &storage.Capacity)
		if err != nil {
			log.Print(err.Error())
			return http.StatusInternalServerError, allStorages
		}
		allStorages = append(allStorages, storage)
	}

	return http.StatusOK, allStorages
}
