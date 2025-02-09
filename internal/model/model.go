package model

type UserData struct {
	Browser        string
	BrowserVersion string
	Model          string
	Platform       string
	OS             string
	IP             string
	ShortURL       string
	Bot            bool
	Mobile         bool
}

type RegistrationUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type URLInfo struct {
	ID        int
	URLID     int
	UserAgent string
	Device    string
	OS        string
	IPAddress string
}

type AuthorizationData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Config struct {
	Host    string
	DBCon   string
	SiteURL string
}

type URLReq struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}
