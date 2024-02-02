package errors

type ErrPermissionForbidden struct{}

func (e ErrPermissionForbidden) Error() string {
	return "forbidden access"
}
