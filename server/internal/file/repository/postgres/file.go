package postgres

import (
	"github.com/22Fariz22/mycloud/server/internal/file"
	"github.com/jmoiron/sqlx"
)

// GetList()
// Upload()
// Download()
// Share()
// Delete()

type File struct {
	ID     int
	UserID int
	URL    string
	Title  string
}

type pgRepository struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) file.FileRepository {
	return &pgRepository{db: db}
}

func (r *pgRepository) GetList() {

}

func (r *pgRepository) Upload() {

}

func (r *pgRepository) Download() {

}

func (r *pgRepository) Share() {

}

func (r *pgRepository) Delete() {

}
