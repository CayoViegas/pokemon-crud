package models

type Pokemon struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" binding:"required" gorm:"unique:not null"`
	EvolvesAt int    `json:"evolves_at" binding:"required"`
}
