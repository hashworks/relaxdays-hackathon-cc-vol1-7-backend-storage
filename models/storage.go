package models

// The unique name of a storage
type StorageName = string

// A struct representing an inventory of an article in a storage
type Storage struct {
	Name StorageName `json:"name,omitempty"`

	ArticleID int `json:"articleID,omitempty"`

	Stock int64 `json:"bestand,omitempty"`
}

func (s Storage) IsValid() bool {
	return s.ArticleID > 0 && len(s.Name) > 0 && s.Stock >= 0
}
