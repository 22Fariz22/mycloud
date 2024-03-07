package repository

import (
	"context"
	"log"

	"github.com/22Fariz22/mycloud/server/internal/entity"
	userService "github.com/22Fariz22/mycloud/server/proto"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// UserRepository User repository
type UserRepository struct {
	db *sqlx.DB
}

// NewUserPGRepository User repository constructor
func NewUserPGRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create new user
func (r *UserRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	createdUser := &entity.User{}
	if err := r.db.QueryRowxContext(
		ctx,
		createUserQuery,
		user.Login,
		user.Password,
	).StructScan(createdUser); err != nil {
		return nil, errors.Wrap(err, "Create.QueryRowxContext")
	}

	return createdUser, nil
}

// FindByLogin Find by user email address
func (r *UserRepository) FindByLogin(ctx context.Context, login string) (*entity.User, error) {
	log.Println("repo FindByLogin()")

	user := &entity.User{}
	if err := r.db.GetContext(ctx, user, findByLoginQuery, login); err != nil {
		return nil, errors.Wrap(err, "FindByLogin.GetContext")
	}

	return user, nil
}

// FindByID Find user by uuid
func (r *UserRepository) FindByID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	user := &entity.User{}
	if err := r.db.GetContext(ctx, user, findByIDQuery, userID); err != nil {
		return nil, errors.Wrap(err, "FindByID.GetContext")
	}

	return user, nil
}

// AddBinary add binary data
func (r *UserRepository) AddBinary(ctx context.Context, userID string, request *userService.AddBinaryRequest) error {
	_, err := r.db.ExecContext(ctx, addBinaryQuery, userID, request.Title, request.Data)
	if err != nil {
		log.Println("err repo AddBinary in r.db.ExecContext", err)
	}
	return nil
}

// GetByTitle find data by title
func (r *UserRepository) GetByTitle(ctx context.Context, userID string, request *userService.GetByTitleRequest) ([]string, error) {
	var everythingByTitle []string

	binaries := []entity.Binary{}

	//get binaries
	err := r.db.SelectContext(ctx, &binaries, getByTitleBinaryQuery, userID, request.Title)
	if err != nil {
		log.Println("err GetByBinary:", err)
	}
	for _, v := range binaries {
		everythingByTitle = append(everythingByTitle, "Binary-> "+string(v.Data))
	}

	return everythingByTitle, nil
}

// GetFullList find all type of data
func (r *UserRepository) GetFullList(ctx context.Context, userID string) ([]string, error) {
	//all data
	var everythingFullList []string

	binaries := []entity.Binary{}

	//get binaries
	err := r.db.SelectContext(ctx, &binaries, getByFullListBinaryQuery, userID)
	if err != nil {
		log.Println("err getByFullListBinaryQuery:", err)
	}

	for _, v := range binaries {
		everythingFullList = append(everythingFullList, "Binary-> "+"title:"+v.Title+"data:"+string(v.Data))
	}

	return everythingFullList, nil
}
