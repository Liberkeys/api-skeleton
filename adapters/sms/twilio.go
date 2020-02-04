package sms

type TwilioAdapter struct {
	accountSid string
	authToken  string
}

func NewTwilioAdapter(accountSid string, authToken string) *TwilioAdapter {
	return &TwilioAdapter{
		accountSid: accountSid,
		authToken:  authToken,
	}
}

func (adapter *TwilioAdapter) Send(to string, from string, body string) error {
	return nil
}
