package access_rule

import (
	"context"
	"errors"

	"github.com/Danya97i/platform_common/pkg/db"

	"github.com/Danya97i/auth/internal/models/consts"
)

// CheckRuleExist check rule exist
func (r *repo) CheckRuleExist(ctx context.Context, role consts.Role, endpoint string) error {
	rawQuery := `SELECT EXISTS(SELECT 1 FROM access_rules WHERE role = $1 AND endpoint = $2)`

	query := db.Query{RawQuery: rawQuery}
	var result bool
	if err := r.db.DB().ScanOneContext(ctx, &result, query, role, endpoint); err != nil {
		return err
	}

	if !result {
		return errors.New("rule does not exist")
	}
	return nil
}
