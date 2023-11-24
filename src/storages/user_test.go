package storage

import (
	"context"
	"database/sql"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/Hack-Hack-geek-Vol10/services/cmd/config"
	"github.com/Hack-Hack-geek-Vol10/services/pkg/utils"
	"github.com/Hack-Hack-geek-Vol10/services/src/domain"
	"github.com/Hack-Hack-geek-Vol10/services/src/driver/postgres"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
)

var dbconn *sql.DB

var testuser = struct {
	id    string
	name  string
	email string
}{
	id:    utils.RandomString(10),
	name:  utils.RandomString(10),
	email: utils.RandomString(10),
}

func TestMain(m *testing.M) {
	config.LoadEnv()
	conn := postgres.NewConnection()
	defer conn.Close(context.Background())

	db, err := conn.Connection()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	dbconn = db

	os.Exit(m.Run())
}

func TestCreate(t *testing.T) {
	repo := NewUserRepo(dbconn)

	testCases := []struct {
		name        string
		arg         domain.CreateUserParams
		wantErr     error
		wantErrCode string
	}{
		{
			name: "success",
			arg: domain.CreateUserParams{
				UserID: testuser.id,
				Name:   testuser.name,
				Email:  testuser.email,
			},
			wantErr: nil,
		},
		{
			name: "failed-unique-key-violation",
			arg: domain.CreateUserParams{
				UserID: testuser.id,
				Name:   testuser.name,
				Email:  testuser.email,
			},
			wantErr: &pq.Error{Code: pgerrcode.UniqueViolation},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Create(tc.arg)
			if err != nil {
				if tc.wantErr == nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if tc.wantErrCode != "" {
					pqerr, ok := err.(*pq.Error)
					if !ok {
						t.Fatalf("unexpected error type: %T", err)
					}

					if pqerr.Code != pgerrcode.UniqueViolation {
						t.Fatalf("unexpected error code: %v", pqerr.Code)
					}
				}
			}
		})
	}
}

func TestReadOne(t *testing.T) {
	repo := NewUserRepo(dbconn)

	testCases := []struct {
		name    string
		userID  string
		wantErr error
		want    *domain.User
	}{
		{
			name:    "success",
			userID:  testuser.id,
			wantErr: nil,
			want: &domain.User{
				UserID: testuser.id,
				Name:   testuser.name,
				Email:  testuser.email,
			},
		},
		{
			name:    "failed-not-found",
			userID:  "notfound",
			wantErr: sql.ErrNoRows,
			want:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user, err := repo.ReadOne(tc.userID)
			if err != nil {
				if tc.wantErr == nil {
					t.Fatalf("want = %v , got = %v", tc.wantErr, err)
				}
			}

			if !reflect.DeepEqual(user, tc.want) {
				t.Fatalf("want = %v , got = %v", tc.want, user)
			}
		})
	}
}
