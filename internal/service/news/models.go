package news

type News struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
	Roo     string `json:"roo"`
	Org     string `json:"org"`
	Class   string `json:"class"`
	Range   string `json:"range"`
}
