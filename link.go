package hal

type Link struct {
	HRef  string `json:"href"`
	Title string `json:"title,omitempty"`
}

func NewLink(href string) Link {
	return Link{
		HRef: href,
	}
}

func (l Link) SetTitle(title string) Link {
	l.Title = title
	return l
}
