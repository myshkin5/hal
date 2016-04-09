package hal

type Embedded map[string][]interface{}

func NewEmbedded() Embedded {
	return make(Embedded)
}

func (e Embedded) Append(name string, resources ...interface{}) {
	e[name] = append(e[name], resources...)
}

func (e Embedded) AppendList(name string, resources []interface{}) {
	e[name] = append(e[name], resources...)
}
