package models

import (
	"strconv"
	"strings"
)

// A struct representing an inventory of an article in a storage (v0)
type V0Storage struct {
	Name      string `json:"name,omitempty"`
	ArticleID int    `json:"articleID,omitempty"`
	Stock     int64  `json:"bestand,omitempty"`
}

func (s V0Storage) IsValid() bool {
	return s.ArticleID > 0 && len(s.Name) > 0 && s.Stock >= 0
}

// Returns valid, V1Storage
func (s V0Storage) ToV1Storage() (bool, V1Storage) {
	valid, location, segment, row, spot, height := NameToV1Properties(s.Name)
	if !valid {
		return valid, V1Storage{}
	} else {
		return valid, V1Storage{
			Location:  location,
			Segment:   segment,
			Row:       row,
			Spot:      spot,
			Height:    height,
			ArticleID: s.ArticleID,
			Stock:     s.Stock,
		}
	}
}

// Returns valid, location, segment, row, spot and height for a storage name
func NameToV1Properties(name string) (bool, string, int, int, int, int) {
	split1 := strings.Split(name, "-")
	if len(split1) != 2 {
		return false, "", 0, 0, 0, 0
	}
	split2 := strings.Split(split1[1], ";")
	if len(split2) != 4 {
		return false, "", 0, 0, 0, 0
	}
	segment, segment_err := strconv.Atoi(split2[0])
	row, row_err := strconv.Atoi(split2[1])
	spot, spot_err := strconv.Atoi(split2[2])
	height, height_err := strconv.Atoi(split2[3])
	if segment_err != nil || row_err != nil || spot_err != nil || height_err != nil {
		return false, "", 0, 0, 0, 0
	}
	return true, split1[0], segment, row, spot, height
}
