package models

// Action is the type of action
type Action string

const (
	ActionCreate Action = "create" // ActionCreate - add user
	ActionGet    Action = "get"    // ActionGet - get user
	ActionUpdate Action = "update" // ActionUpdate - update user
	ActionDelete Action = "delete" // ActionDelete - delete user
)

// LogInfo is the log info
type LogInfo struct {
	UserID int64
	Action Action
}
