package authdb

import (
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	OpCreate = iota
	OpRead
	OpUpdate
	OpDelete
)

type DB interface {
	Create(DBObject) error
}

type DBObject interface {
	CreateStatement() string
}

type User2 interface {
	Create() error
	Update() error
	SelectByUserID(id int64) error
}

func NewPGUser(db *sqlx.DB) *PGUser {
	return &PGUser{
		db: db,
	}
}

type PGUser struct {
	db       *sqlx.DB
	UserID   int64     `db:"user_id" json:"user_id"`
	Username string    `db:"username" json:"username"`
	Password string    `db:"password" json:"password"`
	Salt     string    `db:"salt" json:"salt"`
	Created  time.Time `db:"created" json:"created"`
	Updated  time.Time `db:"updated" json:"updated"`
}

func (u *PGUser) Create() error {
	nstmt, err := u.db.PrepareNamed(`INSERT into users (username, password, salt) VALUES (:username, :password, :salt) returning *`)
	if err != nil {
		return err
	}
	defer nstmt.Close()

	return nstmt.QueryRow(u).StructScan(u)
}
func (u *PGUser) Update() error {
	return nil
}
func (u *PGUser) SelectByUserID(id int64) error {
	return nil
}

type User struct {
	UserID   int64     `db:"user_id" json:"user_id"`
	Username string    `db:"username" json:"username"`
	Password string    `db:"password" json:"password"`
	Salt     string    `db:"salt" json:"salt"`
	Created  time.Time `db:"created" json:"created"`
	Updated  time.Time `db:"updated" json:"updated"`
}

type Role struct {
	RoleID   int64  `db:"role_id" json:"role_id"`
	RoleName string `db:"role_name" json:"role_name"`
}

type UserRole struct {
	UserRoleID int64 `db:"user_role_id" json:"user_role_id"`
	UserID     int64 `db:"user_id" json:"user_id"`
	RoleID     int64 `db:"role_id" json:"role_id"`
}

type RoleHierarchy struct {
	ParentRoleID int64 `db:"parent_role_id" json:"parent_role_id"`
	ChildRoleID  int64 `db:"child_role_id" json:"child_role_id"`
}

type Permission struct {
	PermissionID   int64  `db:"permission_id" json:"permission_id"`
	PermissionName string `db:"permission_name" json:"permission_name"`
}

type RolePermission struct {
	RolePermissionID int64 `db:"role_permission_id" json:"role_permission_id"`
	RoleID           int64 `db:"role_id" json:"role_id"`
	PermissionID     int64 `db:"permission_id" json:"permission_id"`
}
