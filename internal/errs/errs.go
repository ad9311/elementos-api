package errs

const (
	PswdConfMismatch = "password confirmation mismatches"
	WrongPswdOrUser  = "incorrect password or username"
	MissingFormField = "%s field is missing"
	EmptyFormField   = "%s field is empty"
	ExpiredDate      = "%s date has expired"
	FormErr          = "an error has occurred in the form"
	InvNotExists     = "invitation code does not exist"
	InternalErr      = "there was an internal error"
	UserNotInserted  = "user could not be created"
)
