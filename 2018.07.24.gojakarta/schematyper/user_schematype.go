package entity

// generated by "schematyper --package entity user.json" -- DO NOT EDIT

import "time"

type User struct {
	ID       string    `json:"id,omitempty"`
	JoinDate time.Time `json:"join_date,omitempty"`
	Name     string    `json:"name,omitempty"`
}
