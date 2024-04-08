package sql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserAccountModel = (*customUserAccountModel)(nil)

type (
	// UserAccountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAccountModel.
	UserAccountModel interface {
		userAccountModel
		withSession(session sqlx.Session) UserAccountModel
	}

	customUserAccountModel struct {
		*defaultUserAccountModel
	}
)

// NewUserAccountModel returns a model for the database table.
func NewUserAccountModel(conn sqlx.SqlConn) UserAccountModel {
	return &customUserAccountModel{
		defaultUserAccountModel: newUserAccountModel(conn),
	}
}

func (m *customUserAccountModel) withSession(session sqlx.Session) UserAccountModel {
	return NewUserAccountModel(sqlx.NewSqlConnFromSession(session))
}
