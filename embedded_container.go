package hal

type EmbeddedContainer map[string][]interface{}

func NewEmbeddedContainer() EmbeddedContainer {
	return make(EmbeddedContainer)
}

func (e EmbeddedContainer) Append(name string, resources ...interface{}) {
	e[name] = append(e[name], resources...)
}

func (e EmbeddedContainer) AppendList(name string, resources []interface{}) {
	e[name] = append(e[name], resources...)
}
