package models

type Role uint8

const (
	RoleAdmin Role = iota + 1
	RoleEditor
	RoleMember
)
