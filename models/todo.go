package models

import (
	"time"
)

// Todo - todo model
type Todo struct {
	ID        *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Title     *string   `json:"title"`
	Precio     *int   `json:"precio"`
	Cantidad     *int   `json:"cantidad"`
	Fecha_Vencimiento     *time.Time   `json:"fecha_vencimiento"`
	Fecha_Venta     *time.Time   `json:"fecha_venta"`
	Categotia     *string   `json:"categoria"`
	Completed *bool     `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}