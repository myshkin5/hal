package hal_test

import (
	"encoding/json"

	"github.com/myshkin5/hal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EmbeddedContainer", func() {
	var (
		sWPtr        sprocketWPtr
		s            sprocket
		sWOOmitEmpty sprocketWOOmitEmpty
	)

	Describe("zero value", func() {
		BeforeEach(func() {
			sWPtr = sprocketWPtr{
				Size: "large",
			}
			s = sprocket{
				Size:     "large",
				Embedded: hal.NewEmbeddedContainer(),
			}
			sWOOmitEmpty = sprocketWOOmitEmpty{
				Size:     "large",
				Embedded: hal.NewEmbeddedContainer(),
			}
		})

		It("for pointers, doesn't marshal the embedded map", func() {
			data, err := json.Marshal(sWPtr)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(MatchJSON(`{
				"size": "large"
			}`))
		})

		It("for values, doesn't marshal an empty embedded map", func() {
			data, err := json.Marshal(s)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(MatchJSON(`{
				"size": "large"
			}`))
		})

		It("without the omitempty tag, marshals an empty embedded map", func() {
			data, err := json.Marshal(sWOOmitEmpty)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(MatchJSON(`{
				"size": "large",
				"_embedded": {}
			}`))
		})
	})

	Describe("everything", func() {
		BeforeEach(func() {
			s = sprocket{
				Size:     "large",
				Embedded: hal.NewEmbeddedContainer(),
			}

			w1 := widget{
				Links: hal.NewLinksContainer("/widgets/1000"),
				Color: "blue",
			}
			w2 := widget{
				Links: hal.NewLinksContainer("/widgets/1001"),
				Color: "red",
			}
			s.Embedded.AppendList("widget", []interface{}{w1, w2})

			f1 := widget{
				Links: hal.NewLinksContainer("/fidgets/2000"),
				Color: "neon",
			}
			s.Embedded.Append("fidget", f1)
			f2 := widget{
				Links: hal.NewLinksContainer("/fidgets/2001"),
				Color: "heliotrope",
			}
			s.Embedded.Append("fidget", f2)
		})

		It("has HAL-compliant JSON", func() {
			data, err := json.Marshal(s)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(MatchJSON(`{
				"size": "large",
				"_embedded": {
					"widget": [
						{
							"_links": {
								"self": { "href": "/widgets/1000" }
							},
							"color": "blue"
						},
						{
							"_links": {
								"self": { "href": "/widgets/1001" }
							},
							"color": "red"
						}
					],
					"fidget": [
						{
							"_links": {
								"self": { "href": "/fidgets/2000" }
							},
							"color": "neon"
						},
						{
							"_links": {
								"self": { "href": "/fidgets/2001" }
							},
							"color": "heliotrope"
						}
					]
				}
			}`))
		})
	})
})

type sprocketWPtr struct {
	Size     string                 `json:"size"`
	Embedded *hal.EmbeddedContainer `json:"_embedded,omitempty"`
}

type sprocket struct {
	Size     string                `json:"size"`
	Embedded hal.EmbeddedContainer `json:"_embedded,omitempty"`
}

type sprocketWOOmitEmpty struct {
	Size     string                `json:"size"`
	Embedded hal.EmbeddedContainer `json:"_embedded"`
}
