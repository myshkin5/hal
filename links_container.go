package hal

type LinksContainer struct {
	Self Self `json:"self"`
}

type Self struct {
	HRef string `json:"href"`
}

func NewLinksContainer(selfURL string) LinksContainer {
	return LinksContainer{
		Self: Self{
			HRef: selfURL,
		},
	}
}
