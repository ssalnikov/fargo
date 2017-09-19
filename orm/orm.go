package orm

type FieldMapper interface {
	GetValue() interface{}
	SetValue(interface{})
}

type Meta struct {
	Primary bool

	alias string
}

type IntField struct {
	Meta
	Value int
}

func (i *IntField) GetValue() interface{} {
	return i.Value
}

func (i *IntField) SetValue(v interface{}) {
	if s, ok := v.(int); ok {
		i.Value = s
	}
}

type CharField struct {
	Meta
	Value string
}

func (i *CharField) GetValue() interface{} {
	return i.Value
}

func (i *CharField) SetValue(v interface{}) {
	if s, ok := v.(string); ok {
		i.Value = s
	}
}

type Fields []FieldMapper
