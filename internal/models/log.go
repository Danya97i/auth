package models

type Action string

const (
	ActionCreate Action = "create"
	ActionGet    Action = "get"
	ActionUpdate Action = "update"
	ActionDelete Action = "delete"
)

type LogInfo struct {
	UserID int64
	Action Action
}
