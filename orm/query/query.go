package query

// Query object
type Query struct {
	queries []Query
	args    []interface{}
	values  []interface{}
}

// Build sql string
func (q *Query) Build() string {
	var query string
	for _, qr := range q.queries {
		query += qr.Build() + "\n"
	}
	return query
}

// Args list of query
func (q *Query) Args() []interface{} {
	if len(q.args) == 0 {
		return nil
	}

	args := make([]interface{}, len(q.args))
	for i := 0; i < len(q.args); i++ {
		args[i] = q.args[i]
	}

	for _, qr := range q.queries {
		args = append(args, qr.Args()...)
	}

	return args
}

// Values list of query
func (q *Query) Values() []interface{} {
	if len(q.values) == 0 {
		return nil
	}

	values := make([]interface{}, len(q.values))
	for i := 0; i < len(q.values); i++ {
		values[i] = q.values[i]
	}

	for _, qr := range q.queries {
		values = append(values, qr.Values()...)
	}

	return values
}
