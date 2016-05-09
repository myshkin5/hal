package hal

import (
	"bytes"
	"encoding/json"
)

type LinksContainer struct {
	Self            Link
	SimpleRelations map[string]Link
	Relations       map[string][]Link
}

func NewLinksContainer(selfURL string) LinksContainer {
	return LinksContainer{
		Self: Link{
			HRef: selfURL,
		},
		SimpleRelations: make(map[string]Link),
		Relations:       make(map[string][]Link),
	}
}

func (l LinksContainer) AddSimpleRelation(relation string, link Link) {
	l.SimpleRelations[relation] = link
}

func (l LinksContainer) AddRelation(relation string, links []Link) {
	l.Relations[relation] = links
}

func (l LinksContainer) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.WriteString(`{"self":`)
	selfBytes, err := json.Marshal(l.Self)
	if err != nil {
		return nil, err
	}
	buf.Write(selfBytes)

	for relation, link := range l.SimpleRelations {
		err = marshalChild(relation, link, buf)
		if err != nil {
			return nil, err
		}
	}

	for relation, link := range l.Relations {
		err = marshalChild(relation, link, buf)
		if err != nil {
			return nil, err
		}
	}

	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func marshalChild(name string, child interface{}, buf *bytes.Buffer) error {
	buf.Write([]byte(`,"`))
	buf.WriteString(name)
	buf.Write([]byte(`":`))
	bytes, err := json.Marshal(child)
	if err != nil {
		return err
	}
	buf.Write(bytes)
	return nil
}
