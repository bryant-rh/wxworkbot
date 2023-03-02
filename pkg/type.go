package pkg

type message struct {
	MsgType string `json:"msgtype"`
}

//text
type textMessage struct {
	message
	Text Text `json:"text"`
}

type Text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

//markdown
type markdownMessage struct {
	message
	Markdown Markdown `json:"markdown"`
}

type Markdown struct {
	Content string `json:"content"`
}

//image
type imageMessage struct {
	message
	Image Image `json:"image"`
}

type Image struct {
	Base64 string `json:"base64"`
	MD5    string `json:"md5"`
}

//news
type newsMessage struct {
	message
	News News `json:"news"`
}

type News struct {
	Articles []NewsArticle `json:"articles"`
}

type NewsArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

//response
type wxWorkResponse struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}
