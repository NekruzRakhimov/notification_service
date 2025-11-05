package driven

type AuthServiceClient interface {
	GetAllEmails() ([]string, error)
}
