package authdb

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres db bindings
)

// DBConfig ...
type DBConfig struct {
	Options map[string]string
}

// NewDBConfig ...
// From lib/pq:
// * dbname - The name of the database to connect to
// * user - The user to sign in as
// * password - The user's password
// * host - The host to connect to. Values that start with / are for unix domain sockets. (default is localhost)
// * port - The port to bind to. (default is 5432)
// * sslmode - Whether or not to use SSL (default is require, this is not the default for libpq)
// * fallback_application_name - An application_name to fall back to if one isn't provided.
// * connect_timeout - Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
// * sslcert - Cert file location. The file must contain PEM encoded data.
// * sslkey - Key file location. The file must contain PEM encoded data.
// * sslrootcert - The location of the root certificate file. The file must contain PEM encoded data.
func NewDBConfig() (DBConfig, error) {
	c := DBConfig{
		Options: map[string]string{},
	}
	c.Options["host"] = "localhost"
	c.Options["port"] = "5432"
	c.Options["dbname"] = "postgres"
	c.Options["sslmode"] = "disable"

	return c, nil
}

// OptionsStr ...
func (dbc DBConfig) OptionsStr() string {
	opts := ""
	for k, v := range dbc.Options {
		opts += k + "=" + v + " "
	}
	return opts
}

// NewDB ...
func NewDB(config DBConfig) (*DB, error) {
	db, err := sqlx.Connect("postgres", config.OptionsStr())
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to postgres %s:%s (%s)", config.Options["host"], config.Options["port"], err)
	}
	return &DB{db: db}, nil
}

type DB struct {
	db *sqlx.DB
}

func (db *DB) User() *User {
	return NewUser(db.db)
}

func (db *DB) Setup() error {
	sqls := []string{
		"DROP DATABASE IF EXISTS app_database",
		"DROP USER IF EXISTS app_user",
		"DROP ROLE IF EXISTS app_user_role",
		"CREATE DATABASE app_database ENCODING 'UTF8'",
		"CREATE ROLE app_user_role LOGIN CREATEROLE",
		"CREATE USER app_user PASSWORD 'app_user_password'",
		"GRANT app_user_role to app_user",
		"GRANT ALL PRIVILEGES ON DATABASE app_database to app_user_role",
	}
	for _, sql := range sqls {
		_, err := db.db.Exec(sql)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewUser(db *sqlx.DB) *User {
	return &User{
		db: db,
	}
}

type User struct {
	db       *sqlx.DB
	UserID   int64     `db:"user_id" json:"user_id"`
	Username string    `db:"username" json:"username"`
	Password string    `db:"password" json:"password"`
	Salt     string    `db:"salt" json:"salt"`
	Created  time.Time `db:"created" json:"created"`
	Updated  time.Time `db:"updated" json:"updated"`
}

func (u *User) CreateTable() error {
	sqls := []string{
		`CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated = now();
    RETURN NEW;
END;
$$ language 'plpgsql';`,
		`
CREATE TABLE IF NOT EXISTS users(
   user_id  BIGSERIAL    NOT NULL,
   username VARCHAR(255) NOT NULL,
   password VARCHAR(255) NOT NULL,
   salt     CHAR(50)     NOT NULL,
   created  timestamp    without time zone DEFAULT now(),
   updated  timestamp    without time zone DEFAULT now(),
   PRIMARY KEY( user_id )
)`,
		"CREATE UNIQUE INDEX IF NOT EXISTS index_users_on_username ON users USING btree (username)",
		"DROP TRIGGER IF EXISTS users_updated ON users",
		"CREATE TRIGGER users_updated BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_modified_column()",
	}
	for _, sql := range sqls {
		_, err := u.db.Exec(sql)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *User) Create() error {
	nstmt, err := u.db.PrepareNamed(`INSERT into users (username, password, salt) VALUES (:username, :password, :salt) returning *`)
	if err != nil {
		return err
	}
	defer nstmt.Close()

	return nstmt.QueryRow(u).StructScan(u)

}

// ------------------------------------------------------------------------
// func (u *User) Create() error {
// 	nstmt, err := u.db.PrepareNamed(`INSERT into users (username, password, salt) VALUES (:username, :password, :salt) returning *`)
// 	if err != nil {
// 		return err
// 	}
// 	defer nstmt.Close()

// 	return nstmt.QueryRow(u).StructScan(u)
// }
// func (u *User) Update() error {
// 	return nil
// }
// func (u *User) SelectByUserID(id int64) error {
// 	return nil
// }

// type User struct {
// 	UserID   int64     `db:"user_id" json:"user_id"`
// 	Username string    `db:"username" json:"username"`
// 	Password string    `db:"password" json:"password"`
// 	Salt     string    `db:"salt" json:"salt"`
// 	Created  time.Time `db:"created" json:"created"`
// 	Updated  time.Time `db:"updated" json:"updated"`
// }

// type Role struct {
// 	RoleID   int64  `db:"role_id" json:"role_id"`
// 	RoleName string `db:"role_name" json:"role_name"`
// }

// type UserRole struct {
// 	UserRoleID int64 `db:"user_role_id" json:"user_role_id"`
// 	UserID     int64 `db:"user_id" json:"user_id"`
// 	RoleID     int64 `db:"role_id" json:"role_id"`
// }

// type RoleHierarchy struct {
// 	ParentRoleID int64 `db:"parent_role_id" json:"parent_role_id"`
// 	ChildRoleID  int64 `db:"child_role_id" json:"child_role_id"`
// }

// type Permission struct {
// 	PermissionID   int64  `db:"permission_id" json:"permission_id"`
// 	PermissionName string `db:"permission_name" json:"permission_name"`
// }

// type RolePermission struct {
// 	RolePermissionID int64 `db:"role_permission_id" json:"role_permission_id"`
// 	RoleID           int64 `db:"role_id" json:"role_id"`
// 	PermissionID     int64 `db:"permission_id" json:"permission_id"`
// }
