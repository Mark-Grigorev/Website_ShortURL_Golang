package model

type DataOfUser struct {
	FirstName  string
	MiddleName string
	LastName   string
}
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
type URLReq struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type InsertDataForURLs struct {
	UserID   int64
	LongURL  string
	ShortURL string
}
