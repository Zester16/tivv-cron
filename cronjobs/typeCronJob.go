package cronjobs

type NewsLetterStruct struct {
	Date     string `json:"date"`
	NewsBody string `json:"newsbody"`
}

type NewsObject struct {
	Date    string `json:"date"`
	NewsUrl string `json:"newsUrl"`
}
