package public

type ClientInfomationFormat struct {
	Behavior string `json:"behavior,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}

type BackendInformation struct {
	Status  string `json:"behavior,omitempty"`
	Message string `json:"message,omitempty"`
}
