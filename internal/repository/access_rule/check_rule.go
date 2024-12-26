package access_rule

import (
	"context"
	"errors"
	"strings"

	"github.com/Danya97i/platform_common/pkg/db"

	"github.com/Danya97i/auth/internal/models/consts"
)

// CheckRuleExist check rule exist
func (r *repo) CheckRuleExist(ctx context.Context, role consts.Role, endpoint string) error {
	rawQuery := `SELECT EXISTS(SELECT 1 FROM access_rules WHERE role = $1 AND endpoint = $2)`

	query := db.Query{RawQuery: rawQuery}
	var result bool
	if err := r.db.DB().ScanOneContext(ctx, &result, query, strings.ToLower(string(role)), endpoint); err != nil {
		return err
	}

	if !result {
		return errors.New("rule does not exist")
	}
	return nil
}
