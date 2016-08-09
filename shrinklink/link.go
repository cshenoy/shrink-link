package shrinklink

type Link struct {
	LongURL   string `json:"long_url"`
	Hash      string `json:"hash"`
	ShortURL  string `json:"short_url"`
	CreatedAt string `json:"created_at"`
}

type UrlPayload struct {
	LongURL    string `json:"longUrl"`
	CustomHash string `json:"customHash"`
}
