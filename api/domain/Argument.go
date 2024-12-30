package domain

type Argument struct {
	UserName            string `json:"UserName"`
	Password            string `json:"Password"`
	ImapUrl             string `json:"ImapUrl"`
	Port                int    `json:"Port"`
	InputFilterFileName string `json:"InputFilterFileName"`
}
