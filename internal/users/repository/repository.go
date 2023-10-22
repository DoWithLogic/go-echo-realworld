package repository

import (
	"context"
	"database/sql"

	"github.com/DoWithLogic/go-echo-realworld/internal/users/entities"
	"github.com/DoWithLogic/go-echo-realworld/internal/users/repository/repository_query"
	"github.com/DoWithLogic/go-echo-realworld/pkg/datasource"
	"github.com/DoWithLogic/go-echo-realworld/pkg/otel/zerolog"
	"github.com/DoWithLogic/go-echo-realworld/pkg/utils"
	"github.com/jmoiron/sqlx"
)

type (
	Repository interface {
		Atomic(ctx context.Context, opt *sql.TxOptions, repo func(tx Repository) error) error

		SaveNewUser(context.Context, entities.Users) (int64, error)
		GetUserByID(context.Context, int64, ...entities.LockingOpt) (entities.Users, error)
		GetUserByEmail(context.Context, string) (entities.Users, error)
		IsUserExist(ctx context.Context, email string) bool
		UpdateUser(context.Context, entities.Users) error

		SaveNewProfile(ctx context.Context, req entities.Profile) (int64, error)
		GetUserProfile(ctx context.Context, userName string) (entities.Profile, error)
		IsUserFollowed(ctx context.Context, user_id, follow_user_id int64) bool
	}

	repository struct {
		db   *sqlx.DB
		conn datasource.ConnTx
		log  *zerolog.Logger
	}
)

func NewRepository(c *sqlx.DB, l *zerolog.Logger) Repository {
	return &repository{conn: c, log: l, db: c}
}

// Atomic implements vendor.Repository for transaction query
func (r *repository) Atomic(ctx context.Context, opt *sql.TxOptions, repo func(tx Repository) error) error {
	txConn, err := r.db.BeginTx(ctx, opt)
	if err != nil {
		r.log.Z().Err(err).Msg("[repository]Atomic.BeginTxx")

		return err
	}

	newRepository := &repository{conn: txConn, db: r.db}

	repo(newRepository)

	if err := new(datasource.DataSource).EndTx(txConn, err); err != nil {
		return err
	}

	return nil
}

func (repo *repository) SaveNewUser(ctx context.Context, user entities.Users) (userID int64, err error) {
	args := utils.Array{
		user.UserName,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.CreatedBy,
	}

	err = new(datasource.DataSource).ExecSQL(repo.conn.ExecContext(ctx, repository_query.InsertUsers, args...)).Scan(nil, &userID)
	if err != nil {
		repo.log.Z().Err(err).Msg("users.repository.SaveNewUser.ExecContext")

		return userID, err
	}

	return userID, nil
}

func (repo *repository) GetUserByID(ctx context.Context, userID int64, options ...entities.LockingOpt) (userData entities.Users, err error) {
	args := utils.Array{
		userID,
	}

	row := func(idx int) utils.Array {
		return utils.Array{
			&userData.UserID,
			&userData.Email,
			&userData.UserName,
			&userData.Bio,
			&userData.Image,
		}
	}

	query := repository_query.GetUserByID

	if len(options) >= 1 && options[0].PessimisticLocking {
		query += " FOR UPDATE"
	}

	if err = new(datasource.DataSource).QuerySQL(repo.conn.QueryContext(ctx, query, args...)).Scan(row); err != nil {
		repo.log.Z().Err(err).Msg("users.repository.GetUserByID.QueryContext")
		return userData, err
	}

	return userData, err
}

func (repo *repository) IsUserExist(ctx context.Context, email string) bool {
	args := utils.Array{email}

	var id int64
	row := func(idx int) utils.Array {
		return utils.Array{
			&id,
		}
	}

	err := new(datasource.DataSource).QuerySQL(repo.conn.QueryContext(ctx, repository_query.IsUserExist, args...)).Scan(row)
	if err != nil {
		repo.log.Z().Warn().Err(err).Msg("users.repository.IsUserExist.QueryContext")

		return false
	}

	return id != 0
}

func (repo *repository) GetUserByEmail(ctx context.Context, email string) (userDetail entities.Users, err error) {
	args := utils.Array{
		email,
	}

	row := func(idx int) utils.Array {
		return utils.Array{
			&userDetail.UserID,
			&userDetail.Email,
			&userDetail.Password,
			&userDetail.UserName,
			&userDetail.Bio,
			&userDetail.Image,
		}
	}

	err = new(datasource.DataSource).QuerySQL(repo.conn.QueryContext(ctx, repository_query.GetUserByEmail, args...)).Scan(row)
	if err != nil {
		repo.log.Z().Err(err).Msg("users.repository.GetUserByEmail.QueryContext")
		return entities.Users{}, err
	}

	return userDetail, err
}

func (repo *repository) UpdateUser(ctx context.Context, req entities.Users) error {
	args := utils.Array{
		req.UserName, req.UserName,
		req.Email, req.Email,
		req.Password, req.Password,
		req.Bio, req.Bio,
		req.Image, req.Image,
		req.UpdatedAt,
		req.UpdatedBy,
		req.UserID,
	}

	err := new(datasource.DataSource).ExecSQL(repo.conn.ExecContext(ctx, repository_query.UpdateUser, args...)).Scan(nil, nil)
	if err != nil {
		repo.log.Z().Err(err).Msg("users.repository.UpdateUser.ExecContext")

		return err
	}

	return nil
}

func (repo *repository) SaveNewProfile(ctx context.Context, req entities.Profile) (int64, error) {
	args := utils.Array{
		req.UserID,
		req.FollowUserID,
		req.IsActive,
		req.CreatedAt,
		req.CreatedBy,
	}

	var profileID int64
	err := new(datasource.DataSource).ExecSQL(repo.conn.ExecContext(ctx, repository_query.InsertNewProfile, args...)).Scan(nil, &profileID)
	if err != nil {
		repo.log.Z().Err(err).Msg("users.repository.SaveNewProfile.ExecContext")

		return profileID, err
	}

	return profileID, nil
}

func (repo *repository) GetUserProfile(ctx context.Context, userName string) (entities.Profile, error) {
	args := utils.Array{
		userName,
	}
	var profile entities.Profile
	row := func(idx int) utils.Array {
		return utils.Array{
			&profile.UserID,
			&profile.UserName,
			&profile.Bio,
			&profile.Image,
		}
	}

	err := new(datasource.DataSource).QuerySQL(repo.conn.QueryContext(ctx, repository_query.GetProfileByUserName, args...)).Scan(row)
	if err != nil {
		repo.log.Z().Err(err).Msg("users.repository.GetUserProfile.QueryContext")

		return entities.Profile{}, err
	}

	return profile, nil
}

func (repo *repository) IsUserFollowed(ctx context.Context, user_id, follow_user_id int64) bool {
	args := utils.Array{
		user_id,
		follow_user_id,
	}

	var profile entities.Profile
	row := func(idx int) utils.Array {
		return utils.Array{
			&profile.UserID,
			&profile.FollowUserID,
		}
	}

	err := new(datasource.DataSource).QuerySQL(repo.conn.QueryContext(ctx, repository_query.IsUserFollowed, args...)).Scan(row)
	if err != nil {
		repo.log.Z().Warn().Err(err).Msg("users.repository.IsUserFollowed.QueryContext")

		return false
	}

	return user_id == profile.UserID && follow_user_id == profile.FollowUserID
}
