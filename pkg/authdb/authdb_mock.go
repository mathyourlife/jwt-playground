package authdb

import (
	"container/list"
	"fmt"
	"log"
	"sync"
)

type MockDB struct {
	usersLock sync.Mutex
	users     *list.List
}

func NewMockDB() *MockDB {
	return &MockDB{
		users: list.New(),
	}
}

func (m *MockDB) BasicAuth(username, password string) (int64, error) {
	u, err := m.User(&User{Username: username, Password: password}, OpRead)
	if err != nil {
		return 0, err
	}
	return u.UserID, nil
}

func (m *MockDB) User(u *User, op int) (*User, error) {
	m.usersLock.Lock()
	defer m.usersLock.Unlock()
	if op == OpCreate {
		if m.users.Len() == 0 {
			u.UserID = 1
		} else {
			user := m.users.Back()
			u.UserID = user.Value.(*User).UserID + 1
		}
		m.users.PushBack(u)
		return u, nil
	} else if op == OpRead {
		log.Println(u)
		user := m.users.Front()
		if user == nil {
			return nil, fmt.Errorf("unable to locate user %s", u)
		}
		for {
			if user.Value.(*User).UserID == u.UserID {
				return user.Value.(*User), nil
			} else if user.Value.(*User).Username == u.Username && user.Value.(*User).Password == u.Password {
				return user.Value.(*User), nil
			}
			user = user.Next()
			if user == nil {
				return nil, fmt.Errorf("unable to locate user %s", u)
			}
		}
	} else if op == OpUpdate {

	} else if op == OpDelete {
		user := m.users.Front()
		for {
			if user.Value.(*User).UserID == u.UserID {
				m.users.Remove(user)
				return nil, nil
			}
			user = user.Next()
			if user == nil {
				return nil, fmt.Errorf("unable to locate user %s", u)
			}
		}
	}
	user := m.users.Front()
	for {
		log.Println(user.Value)
		user = user.Next()
		if user == nil {
			break
		}
	}
	return u, nil
}
func (m *MockDB) Role(r *Role, op int) (*Role, error) {
	return nil, nil
}
func (m *MockDB) UserRole(ur *UserRole, op int) (*UserRole, error) {
	return nil, nil
}
func (m *MockDB) RoleHierarchy(rh *RoleHierarchy, op int) (*RoleHierarchy, error) {
	return nil, nil
}
func (m *MockDB) Permission(p *Permission, op int) (*Permission, error) {
	return nil, nil
}
func (m *MockDB) RolePermission(rp *RolePermission, op int) (*RolePermission, error) {
	return nil, nil
}
