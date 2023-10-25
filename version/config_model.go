package version

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ConfigModel = (*customConfigModel)(nil)

type (
	// ConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customConfigModel.
	ConfigModel interface {
		configModel
		withSession(session sqlx.Session) ConfigModel
	}

	customConfigModel struct {
		*defaultConfigModel
	}
)

// NewConfigModel returns a model for the database table.
func NewConfigModel(conn sqlx.SqlConn) ConfigModel {
	return &customConfigModel{
		defaultConfigModel: newConfigModel(conn),
	}
}

func (m *customConfigModel) withSession(session sqlx.Session) ConfigModel {
	return NewConfigModel(sqlx.NewSqlConnFromSession(session))
}
