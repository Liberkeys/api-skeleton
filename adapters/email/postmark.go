package email

type PostmarkAdapter struct {
	apiKey string
}

func NewPostmarkAdapter(apiKey string) *PostmarkAdapter {
	return &PostmarkAdapter{
		apiKey: apiKey,
	}
}

func (adapter *PostmarkAdapter) Send(email string) error {
	return nil
}
