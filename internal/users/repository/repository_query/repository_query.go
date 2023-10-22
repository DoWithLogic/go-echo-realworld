package repository_query

import _ "embed"

var (
	//go:embed users/insert.sql
	InsertUsers string
	//go:embed users/select.sql
	GetUserByID string
	//go:embed users/get_detail_by_email.sql
	GetUserByEmail string
	//go:embed users/check_is_user_exist.sql
	IsUserExist string
	//go:embed users/update.sql
	UpdateUser string

	//go:embed profiles/get_by_username.sql
	GetProfileByUserName string
	//go:embed profiles/check_follow.sql
	IsUserFollowed string
	//go:embed profiles/insert.sql
	InsertNewProfile string
)
