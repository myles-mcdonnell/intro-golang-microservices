// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"github.com/google/uuid"
)

type Counter struct {
	ID  uuid.UUID `json:"id"`
	Val int32     `json:"val"`
}