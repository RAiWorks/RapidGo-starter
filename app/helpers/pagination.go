package helpers

import (
	"github.com/raiworks/rapidgo/v2/database"
	"gorm.io/gorm"
)

// PaginateResult is an alias for the framework's PageResult.
type PaginateResult = database.PageResult

// CursorResult is an alias for the framework's CursorResult.
type CursorResult = database.CursorResult

// Paginate delegates to the framework's offset-based pagination.
func Paginate(db *gorm.DB, page, perPage int, dest interface{}) (*PaginateResult, error) {
	return database.Paginate(db, page, perPage, dest)
}

// CursorPaginate delegates to the framework's cursor-based pagination.
func CursorPaginate(db *gorm.DB, cursor, orderCol string, perPage int, direction string, dest interface{}) (*CursorResult, error) {
	return database.CursorPaginate(db, cursor, orderCol, perPage, direction, dest)
}
