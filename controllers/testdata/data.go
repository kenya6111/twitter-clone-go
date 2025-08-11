package testdata

import (
	"twitter-clone-go/tutorial"

	"github.com/jackc/pgx/v5/pgtype"
)

var UserTestData = []tutorial.User{
	tutorial.User{
		ID:       1,
		Name:     "kenya",
		Email:    "kenyanke6111@aaa.com",
		Password: "pass",
		IsActive: pgtype.Bool{
			Bool:  true,
			Valid: true,
		},
	},
	tutorial.User{
		ID:       2,
		Name:     "kenya",
		Email:    "tanaka@aaa.com",
		Password: "pass",
		IsActive: pgtype.Bool{
			Bool:  true,
			Valid: true,
		},
	},
	tutorial.User{
		ID:       3,
		Name:     "kenya",
		Email:    "tanaka@aaa.com",
		Password: "pass",
		IsActive: pgtype.Bool{
			Bool:  true,
			Valid: true,
		},
	},
	tutorial.User{
		ID:       4,
		Name:     "kenya",
		Email:    "tanaka@aaa.com",
		Password: "pass",
		IsActive: pgtype.Bool{
			Bool:  true,
			Valid: true,
		},
	},
}
