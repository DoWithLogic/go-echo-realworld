package repository_test

import (
	"context"
	"os"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/DoWithLogic/go-echo-realworld/internal/users/entities"
	"github.com/DoWithLogic/go-echo-realworld/internal/users/repository"
	"github.com/DoWithLogic/go-echo-realworld/internal/users/repository/repository_query"
	"github.com/DoWithLogic/go-echo-realworld/pkg/otel/zerolog"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_repository_GetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	defer db.Close()

	var (
		conn = sqlx.NewDb(db, "sqlmock")
		repo = repository.NewRepository(conn, zerolog.NewZeroLog(context.Background(), os.Stdout))
	)

	currentTime := time.Now()
	userID := int64(1)

	user := entities.Users{
		UserID:    userID,
		Email:     "realworld@example.com",
		Password:  "pwdtesting",
		UserName:  "go-echo-realworld",
		CreatedAt: currentTime,
		CreatedBy: "SYSTEM",
		Bio:       "",
		Image:     "",
	}

	t.Run("GetUserByID_Positive", func(t *testing.T) {
		mock.
			ExpectExec(repository_query.InsertUsers).
			WithArgs(
				user.UserName,
				user.Email,
				user.Password,
				currentTime,
				user.CreatedBy,
			).WillReturnResult(sqlmock.NewResult(userID, 1))

		createdID, err := repo.SaveNewUser(context.Background(), user)
		require.NoError(t, err)
		require.Equal(t, createdID, userID)

		mock.ExpectQuery(repository_query.GetUserByID).
			WithArgs(userID).
			WillReturnRows(sqlmock.NewRows(
				[]string{"id", "email", "username", "bio", "image"}).
				AddRow(user.UserID, user.Email, user.UserName, user.Bio, user.Image))

		_, err = repo.GetUserByID(context.Background(), userID)
		require.NoError(t, err)
	})
}
