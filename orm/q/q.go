package qr

type Querier interface {
}

type Query struct {
}

func And(...Querier) Querier {
	return &Query{}
}
