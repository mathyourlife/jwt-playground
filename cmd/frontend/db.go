package main

import ()

func getRoleByID(roleID int64) *Role {
	for _, r := range Roles {
		if r.RoleID == roleID {
			return r
		}
	}
	return nil
}

func getRoleHierarchies(rh RoleHierarchy) []*RoleHierarchy {
	rhs := []*RoleHierarchy{}
	for _, item := range RoleHierarchies {
		rhs = append(rhs, item)
	}
	return rhs
}

// func getRolesByParent(r Role) []*Role {
// 	children := []*Role{}
// 	for _, rh := range RoleHierarchies {
// 		if rh.ParentRoleID == roleID {
// 			children = append(children, getRoleByID(rh.ChildRoleID))
// 		}
// 	}
// 	return children
// }
