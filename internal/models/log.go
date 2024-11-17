package models

import "github.com/Danya97i/auth/internal/models/consts"

// LogInfo is the log info
type LogInfo struct {
	UserID int64
	Action consts.Action
}
