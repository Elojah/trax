package cookie

type ErrMissingSecureKeys struct{}

func (err ErrMissingSecureKeys) Error() string {
	return "cookie keys not found"
}
