package pkg

//SendText Send Text
func (c *CMClient) SendText(webhookurl string, content string, mentioned_list []string, mentioned_mobile_list []string) (*wxWorkResponse, error) {
	res := &wxWorkResponse{}
	msg := Text{
		Content:             content,
		MentionedList:       mentioned_list,
		MentionedMobileList: mentioned_mobile_list,
	}
	data, err := marshalMessage(msg)
	if err != nil {
		return res, err
	}
	_, err = c.R().
		SetBody(data).
		SetSuccessResult(res).Post(webhookurl)
	return res, err
}

//SendMarkdown Send Markdown
func (c *CMClient) SendMarkdown(webhookurl, content string) (*wxWorkResponse, error) {
	res := &wxWorkResponse{}
	msg := Markdown{
		Content: content,
	}
	data, err := marshalMessage(msg)
	if err != nil {
		return res, err
	}
	_, err = c.R().
		SetBody(data).
		SetSuccessResult(res).Post(webhookurl)
	return res, err
}

//SendImage Send Image
func (c *CMClient) SendImage(webhookurl, base64, md5 string) (*wxWorkResponse, error) {
	res := &wxWorkResponse{}
	msg := Image{
		Base64: base64,
		MD5:    md5,
	}
	data, err := marshalMessage(msg)
	if err != nil {
		return res, err
	}
	_, err = c.R().
		SetBody(data).
		SetSuccessResult(res).Post(webhookurl)
	return res, err
}

//SendNews Send News
func (c *CMClient) SendNews(webhookurl, title, description, url, picurl string) (*wxWorkResponse, error) {
	res := &wxWorkResponse{}
	msg := News{
		Articles: []NewsArticle{
			{
				Title:       title,
				Description: description,
				URL:         url,
				PicURL:      picurl,
			},
		},
	}
	data, err := marshalMessage(msg)
	if err != nil {
		return res, err
	}
	_, err = c.R().
		SetBody(data).
		SetSuccessResult(res).Post(webhookurl)
	return res, err
}
