package authdb

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq" // Postgres db bindings
)

// PostgresConfig ...
type PostgresConfig struct {
	Options map[string]string
}

// NewPostgresConfig ...
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
func NewPostgresConfig() (PostgresConfig, error) {
	c := PostgresConfig{
		Options: map[string]string{},
	}
	c.Options["host"] = "localhost"
	c.Options["port"] = "5432"
	c.Options["dbname"] = "postgres"
	c.Options["sslmode"] = "disable"

	return c, nil
}

// OptionsStr ...
func (pgc PostgresConfig) OptionsStr() string {
	opts := ""
	for k, v := range pgc.Options {
		opts += k + "=" + v + " "
	}
	return opts
}

type Postgres struct {
	*sqlx.DB
}

// Postgres ...
func NewPostgres(config PostgresConfig) (*Postgres, error) {
	db, err := sqlx.Connect("postgres", config.OptionsStr())
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to postgres %s:%s (%s)", config.Options["host"], config.Options["port"], err)
	}
	return &Postgres{db}, nil
}

func (pg *Postgres) CreateAppDB() error {

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
		_, err := pg.Exec(sql)
		if err != nil {
			return err
		}
	}
	return nil
}

func (pg *Postgres) SetupAppDB() error {
	sqls := []string{
		`CREATE TABLE users (
	user_id BIGSERIAL,
	username varchar(256),
	password varchar(256),
	salt varchar(256),
  created timestamp without time zone DEFAULT now(),
  updated timestamp without time zone DEFAULT now(),
 	CONSTRAINT user_id PRIMARY KEY(user_id)
)`,
		`CREATE UNIQUE INDEX index_users_on_username ON users USING btree (username)`,
	}

	for _, sql := range sqls {
		_, err := pg.Exec(sql)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *User) Create(db *sqlx.DB) error {
	nstmt, err := db.PrepareNamed(`INSERT into users (username, password, salt) VALUES (:username, :password, :salt) returning *`)
	if err != nil {
		return err
	}
	defer nstmt.Close()

	return nstmt.QueryRow(u).StructScan(u)

}

func (pg *Postgres) BasicAuth(username, password string) (int64, error) {
	return 1234, nil
}

func (pg *Postgres) User(u *User, op int) (*User, error) {

	if op == OpCreate {
		nstmt, err := pg.PrepareNamed(`INSERT into users (username, password, salt) VALUES (:username, :password, :salt) returning *`)
		if err != nil {
			return nil, err
		}
		defer nstmt.Close()

		err = nstmt.QueryRow(u).StructScan(u)
		if err, ok := err.(*pq.Error); ok {
			switch err.Code.Name() {
			case "unique_violation":
				// &Error{"unique_violation", "hub exists"}
				return nil, err
			default:
				// &Error{err.Code.Name(), "pq error"}
				return nil, err
			}
		}
	}
	return nil, nil
}

func (pg *Postgres) Role(r *Role, op int) (*Role, error) {
	return nil, nil
}

func (pg *Postgres) UserRole(ur *UserRole, op int) (*UserRole, error) {
	return nil, nil
}

func (pg *Postgres) RoleHierarchy(rh *RoleHierarchy, op int) (*RoleHierarchy, error) {
	return nil, nil
}

func (pg *Postgres) Permission(p *Permission, op int) (*Permission, error) {
	return nil, nil
}

func (pg *Postgres) RolePermission(rp *RolePermission, op int) (*RolePermission, error) {
	return nil, nil
}
