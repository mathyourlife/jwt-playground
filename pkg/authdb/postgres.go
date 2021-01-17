package authdb

import (
// "github.com/lib/pq" // Postgres db bindings
)

func (db *DB) BasicAuth(username, password string) (int64, error) {
	return 1234, nil
}

// func (db *DB) User(u *User, op int) (*User, error) {

// 	if op == OpCreate {
// 		nstmt, err := pg.PrepareNamed(`INSERT into users (username, password, salt) VALUES (:username, :password, :salt) returning *`)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer nstmt.Close()

// 		err = nstmt.QueryRow(u).StructScan(u)
// 		if err, ok := err.(*pq.Error); ok {
// 			switch err.Code.Name() {
// 			case "unique_violation":
// 				// &Error{"unique_violation", "hub exists"}
// 				return nil, err
// 			default:
// 				// &Error{err.Code.Name(), "pq error"}
// 				return nil, err
// 			}
// 		}
// 	}
// 	return nil, nil
// }

// func (db *DB) Role(r *Role, op int) (*Role, error) {
// 	return nil, nil
// }

// func (db *DB) UserRole(ur *UserRole, op int) (*UserRole, error) {
// 	return nil, nil
// }

// func (db *DB) RoleHierarchy(rh *RoleHierarchy, op int) (*RoleHierarchy, error) {
// 	return nil, nil
// }

// func (db *DB) Permission(p *Permission, op int) (*Permission, error) {
// 	return nil, nil
// }

// func (db *DB) RolePermission(rp *RolePermission, op int) (*RolePermission, error) {
// 	return nil, nil
// }
