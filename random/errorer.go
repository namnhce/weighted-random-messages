package random

// Error declarations
var (
	errorVersionLengthInvalid = errVersionLengthInvalid{}
	errorWeightInvalid        = errWeightInvalid{}
)

type errVersionLengthInvalid struct{}

func (errVersionLengthInvalid) Error() string {
	return "version length is invalid"
}

type errWeightInvalid struct{}

func (errWeightInvalid) Error() string {
	return "weight is invalid"
}
