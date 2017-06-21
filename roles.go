package main

import ()

type User struct {
	UserID   int64  `db:"user_id" json:"user_id"`
	First    string `db:"first" json:"first"`
	Last     string `db:"last" json:"last"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Role struct {
	RoleID   int64  `db:"role_id" json:"role_id"`
	RoleName string `db:"role_name" json:"role_name"`
}

type UserRole struct {
	UserRoleID int64 `json:"user_role_id"`
	UserID     int64 `json:"user_id"`
	RoleID     int64 `json:"role_id"`
}

type RoleHierarchy struct {
	ParentRoleID int64 `json:"parent_role_id"`
	ChildRoleID  int64 `json:"child_role_id"`
}

type Permission struct {
	PermissionID   int64  `json:"permission_id"`
	PermissionName string `json:"permission_name"`
}

type RolePermission struct {
	RolePermissionID int64 `json:"role_permission_id"`
	RoleID           int64 `json:"role_id"`
	PermissionID     int64 `json:"permission_id"`
}

var Admin = Role{1, "admin"}
var Teacher = Role{2, "teacher"}
var Student = Role{3, "student"}

var Roles = []Role{Admin, Teacher, Student}

var Christa = User{1, "Christa", "McAuliffe", "cmcauliffe", "c"}
var Sally = User{2, "Sally", "Ride", "sride", "s"}
var Jane = User{3, "Jane", "Austin", "jaustin", "j"}
var Harry = User{4, "Harry", "Truman", "htruman", "h"}

var Users = []User{Christa, Sally, Jane, Harry}

var UserRoles = []UserRole{
	UserRole{1, Christa.UserID, Admin.RoleID},
	UserRole{2, Sally.UserID, Teacher.RoleID},
	UserRole{3, Jane.UserID, Student.RoleID},
	UserRole{3, Harry.UserID, Student.RoleID},
}

var RoleHierarchies = []RoleHierarchy{
	RoleHierarchy{Admin.RoleID, Teacher.RoleID},
	RoleHierarchy{Teacher.RoleID, Student.RoleID},
}

var CreateUsers = Permission{1, "create users"}
var EditGrades = Permission{2, "edit grades"}
var ViewAllGrades = Permission{3, "view all grades"}
var ViewOwnGrades = Permission{4, "view own grades"}

var Permissions = []Permission{CreateUsers, EditGrades, ViewAllGrades, ViewOwnGrades}

var RolePermissions = []RolePermission{
	RolePermission{1, Admin.RoleID, CreateUsers.PermissionID},
	RolePermission{2, Teacher.RoleID, EditGrades.PermissionID},
	RolePermission{3, Teacher.RoleID, ViewAllGrades.PermissionID},
	RolePermission{4, Student.RoleID, ViewOwnGrades.PermissionID},
}
