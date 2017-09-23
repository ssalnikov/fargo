package orm

type FieldMapper interface {
	GetValue() interface{}
	SetValue(interface{})
}

type Fields []FieldMapper
