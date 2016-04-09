package hal

type Links struct {
	Self Self `json:"self"`
}

type Self struct {
	HRef string `json:"href"`
}

func NewLinks(selfURL string) Links {
	return Links{
		Self: Self{
			HRef: selfURL,
		},
	}
}
