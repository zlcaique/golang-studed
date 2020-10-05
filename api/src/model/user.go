package model

import "time"

//User representa um usuario
type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Pass     string    `json:"pass,omitempty"`
	CreateAt time.Time `json:"CreateAt,omitempty"`
}
