package apperrors

type ErrCode string

const (
	Unknown ErrCode = "U000"

	InsertDataFailed    ErrCode = "S001"
	GetDataFailed       ErrCode = "S002"
	NAData              ErrCode = "S003"
	NoTargetData        ErrCode = "S004"
	UpdateDataFailed    ErrCode = "S005"
	DuplicateData       ErrCode = "S006"
	GenerateHashFailed  ErrCode = "S007"
	GenerateTokenFailed ErrCode = "S008"
	DeleteDataFailed    ErrCode = "S009"

	ReqBodyDecodeFailed ErrCode = "R001"
	BadParam            ErrCode = "R002"
)

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
