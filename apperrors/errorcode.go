package apperrors

type ErrCode string

const (
	Unknown               ErrCode = "U000"
	ReqBodyDecodeFailed   ErrCode = "R001"
	ReqBadParam           ErrCode = "R002"
	AuthUnauthorized      ErrCode = "A001"
	AuthSessionSaveFailed ErrCode = "A002"
	AuthLogoutFailed      ErrCode = "A003"
	InsertDataFailed      ErrCode = "S001"
	GetDataFailed         ErrCode = "S002"
	NAData                ErrCode = "S003"
	NoTargetData          ErrCode = "S004"
	UpdateDataFailed      ErrCode = "S005"
	DuplicateData         ErrCode = "S006"
	DeleteDataFailed      ErrCode = "S007"
	GenerateHashFailed    ErrCode = "T001"
	GenerateTokenFailed   ErrCode = "T001"
	SendMailFailed        ErrCode = "M001"
	SendEmailFailed       ErrCode = "M001"
	SaveLocalFileFailed   ErrCode = "F001"
)

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
