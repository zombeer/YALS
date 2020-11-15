package models

import (
	"fmt"
	"yals/utils"
)

// Item :: link mapping structure.
type Item struct {
	URL      string `json:"url"`
	Shortcut string `json:"shortcut"`
}

func (i *Item) String() string {
	return fmt.Sprintf("mapping %s -> %s", i.URL, i.Shortcut)
}

// GenerateHASH :: Creates link mapping.
func (i *Item) GenerateHASH() {
	i.Shortcut = utils.RandString(5)
}
