package book

import (
	"encoding/json"
)

type BookInput struct {
	Title string      `json:"judul" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}
