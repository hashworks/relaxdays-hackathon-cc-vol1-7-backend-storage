package models

import "fmt"

// A struct representing an inventory of an article in a storage (v1)
type V1Storage struct {
	Location  string `json:"standort,omitempty"`
	Segment   int    `json:"lagerabschnitt,omitempty"`
	Row       int    `json:"reihe,omitempty"`
	Spot      int    `json:"platz,omitempty"`
	Height    int    `json:"hoehe,omitempty"`
	ArticleID int    `json:"articleID,omitempty"`
	Stock     int    `json:"bestand,omitempty"`
	Capacity  int    `json:"-"`
}

func (s V1Storage) IsValid() bool {
	return len(s.Location) > 0 &&
		s.Segment >= 0 &&
		s.Row >= 0 &&
		s.Spot >= 0 &&
		s.Height >= 0 &&
		s.ArticleID > 0 &&
		s.Stock >= 0
}

func (s V1Storage) ToV0Storage() V0Storage {
	return V0Storage{
		Name:      s.V1PropertiesToName(),
		ArticleID: s.ArticleID,
		Stock:     s.Stock,
		Capacity:  s.Capacity,
	}
}

func (s V1Storage) ToV2Storage() V2Storage {
	return V2Storage{
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

func (s V1Storage) V1PropertiesToName() string {
	return fmt.Sprintf("%s-%d;%d;%d;%d", s.Location, s.Segment, s.Row, s.Spot, s.Height)
}
