package auth

type ACL_ENUM int

const (
	ADMIN_SUPER ACL_ENUM = iota
	ADMIN
	SPV
	CUSTOMER
	OPERATOR
	HEAD_STORE
)

// *string allow nil
func (e ACL_ENUM) String() *string {
	switch e {
	case ADMIN_SUPER:
		val := "admin_super"
		return &val
	case ADMIN:
		val := "admin"
		return &val
	case SPV:
		val := "spv"
		return &val
	case CUSTOMER:
		val := "customer"
		return &val
	case OPERATOR:
		val := "operator"
		return &val
	case HEAD_STORE:
		val := "head_store"
		return &val
	default:
		return nil
	}
}
