package setting

type MysqlConn struct {
	MYSQL_USERNAME        string `validate:"max=255,min=1"`
	MYSQL_PASSWORD        string `validate:"max=255,min=1"`
	MYSQL_ADDRESS         string `validate:"max=255,min=1"`
	MYSQL_DB_NAME         string `validate:"max=255,min=1"`
	MYSQL_MaxIdleConns    int    `validate:"numeric"`
	MYSQL_MaxOpenConns    int    `validate:"numeric"`
	MYSQL_ConnMaxIdleTime int    `validate:"numeric"`
	MYSQL_ConnMaxLifetime int    `validate:"numeric"`
}
