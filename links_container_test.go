package hal_test

import (
	"encoding/json"

	"github.com/myshkin5/hal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LinksContainer", func() {
	var (
		w widget
	)

	Describe("zero value", func() {
		BeforeEach(func() {
			w = widget{
				Color: "blue",
			}
		})

		It("has HAL-compliant JSON", func() {
			data, err := json.Marshal(w)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(MatchJSON(`{
				"_links": {
					"self": { "href": "" }
				},
				"color": "blue"
			}`))
		})
	})

	Describe("everything", func() {
		BeforeEach(func() {
			w = widget{
				Links: hal.NewLinksContainer("/widgets/1000"),
				Color: "blue",
			}
		})

		It("has HAL-compliant JSON", func() {
			data, err := json.Marshal(w)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(MatchJSON(`{
				"_links": {
					"self": { "href": "/widgets/1000" }
				},
				"color": "blue"
			}`))
		})
	})
})

type widget struct {
	Links hal.LinksContainer `json:"_links"`
	Color string             `json:"color"`
}
