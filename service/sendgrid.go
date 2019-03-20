package service

type Email struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Personalization struct {
	To      []Email `json:"to"`
	Subject string  `json:"subject"`
}

type Content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
