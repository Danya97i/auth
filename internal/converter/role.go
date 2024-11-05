package converter

import (
	"github.com/Danya97i/auth/internal/models/consts"
	pb "github.com/Danya97i/auth/pkg/user_v1"
)

// ToRoleFromPb конвертирует pb.Role в models.Role
func ToRoleFromPb(r pb.Role) consts.Role {
	return consts.Role(pb.Role_name[int32(r)])
}

// ToRoleFromService конвертирует pb.Role в models.Role
func ToRoleFromService(r consts.Role) pb.Role {
	return pb.Role(pb.Role_value[string(r)])
}
