package inkstone

import "time"

type Response struct {
	Id        int        `json:"id,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
