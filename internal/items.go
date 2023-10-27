package data

import "time"

type Items struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"-"`
	Name       string    `json:"name"`
	Case       int32     `json:"case,omitempty"`
	Rarity     int32     `json:"rarity,omitempty"`
	Conditions []string  `json:"conditions,omitempty"`
	Version    int32     `json:"version"`
}
