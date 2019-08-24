package evaluation

type Error string

func (err Error) Error() string {
	return string(err)
}

const (
	ErrNotImplemented = Error("evaluation: not implemented")
	ErrNotValidated   = Error("evaluation: not validated")
	ErrNotCompiled    = Error("evaluation: not compiled")
)
