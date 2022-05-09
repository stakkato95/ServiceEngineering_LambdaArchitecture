package errs

type AppError struct {
	Code int    `json:"omitempty"`
	Msg  string `json:",msg"`
}
