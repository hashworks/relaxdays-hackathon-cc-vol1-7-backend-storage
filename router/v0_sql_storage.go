package router

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/hashworks/relaxdays-hackathon-cc-vol1-7-backend-storage/models"
)

func (s Server) v0SelectStorageByName(name string) (int, models.V0Storage) {
	storageRow, err := s.DotSelect.QueryRow(s.DB, "select-storage-by-name", name)

	if err != nil {
		log.Print(err.Error())
		return http.StatusInternalServerError, models.V0Storage{}
	}

	storage := models.V0Storage{
		Name: name,
	}

	err = storageRow.Scan(&storage.ArticleID, &storage.Stock, &storage.Capacity)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, storage
		}
		log.Print(err.Error())
		return http.StatusInternalServerError, storage
	}

	return http.StatusOK, storage
}

func (s Server) v0InsertOrUpdateStorage(storage models.V0Storage) int {
	_, err := s.DotAlter.Exec(s.DB, "insert-or-update-storage", storage.Name, storage.ArticleID, storage.Stock, storage.Capacity)
	if err != nil {
		log.Print(err.Error())
		return http.StatusInternalServerError
	}

	s.cacheStore.Flush()

	return http.StatusNoContent
}

func (s Server) v0DeleteStorageByName(name string) int {
	_, err := s.DotAlter.Exec(s.DB, "delete-storage", name)
	if err != nil {
		log.Print(err.Error())
		return http.StatusInternalServerError
	}

	s.cacheStore.Flush()

	return http.StatusNoContent
}

func (s Server) v0GetStoragesByCursor(x string, n string) (int, []models.V0Storage) {
	allStorages := make([]models.V0Storage, 0)

	storageRows, err := s.DotSelect.Query(s.DB, "select-storage-by-cursor", x, n)
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
