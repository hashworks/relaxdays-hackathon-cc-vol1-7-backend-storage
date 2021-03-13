package models

// A struct representing an inventory of an article in a storage (v1)
type V2Storage struct {
	Location  string `json:"standort,omitempty"`
	Segment   int    `json:"lagerabschnitt,omitempty"`
	Row       int    `json:"reihe,omitempty"`
	Spot      int    `json:"platz,omitempty"`
	Height    int    `json:"hoehe,omitempty"`
	ArticleID int    `json:"articleID,omitempty"`
	Stock     int    `json:"bestand,omitempty"`
	Capacity  int    `json:"kapazitaet,omitempty"`
}

func (s V2Storage) IsValid() bool {
	return len(s.Location) > 0 &&
		s.Segment >= 0 &&
		s.Row >= 0 &&
		s.Spot >= 0 &&
		s.Height >= 0 &&
		s.ArticleID > 0 &&
		s.Stock >= 0 &&
		s.Capacity > 0
}

func (s V2Storage) ToV1Storage() V1Storage {
	return V1Storage{
		Location:  s.Location,
		Segment:   s.Segment,
		Row:       s.Row,
		Spot:      s.Spot,
		Height:    s.Height,
		ArticleID: s.ArticleID,
		Stock:     s.Stock,
		Capacity:  s.Capacity,
	}
}
