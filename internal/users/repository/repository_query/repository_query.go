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
)
