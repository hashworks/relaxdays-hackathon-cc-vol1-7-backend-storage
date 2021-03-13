package router

import (
	"log"
	"net/http"

	"github.com/hashworks/relaxdays-hackathon-cc-vol1-7-backend-storage/models"
)

func (s Server) v2GetStoragesByArticleID(x string) (int, []models.V2Storage) {
	allStorages := make([]models.V2Storage, 0)

	storageRows, err := s.DotSelect.Query(s.DB, "select-storage-by-articleID", x)
	defer storageRows.Close()

	if err != nil {
		log.Print(err.Error())
		return http.StatusInternalServerError, allStorages
	}

	for storageRows.Next() {
		var v0Storage models.V0Storage
		err := storageRows.Scan(&v0Storage.Name, &v0Storage.ArticleID, &v0Storage.Stock, &v0Storage.Capacity)
		if err != nil {
			log.Print(err.Error())
			return http.StatusInternalServerError, allStorages
		}
		valid, v1Storage := v0Storage.ToV1Storage()
		if !valid {
			log.Print("Failed to convert v0Storage to v1Storage")
			return http.StatusInternalServerError, allStorages
		}
		allStorages = append(allStorages, v1Storage.ToV2Storage())
	}

	return http.StatusOK, allStorages
}
