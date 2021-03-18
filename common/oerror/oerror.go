package oerror

type OError struct {
	error
}

func (o *OError) Error() string {
	return ""
}
