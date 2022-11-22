package forms

import "net/http"

type errors map[string][]string

// Add adds an error message for a given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns the first error message
func (e errors) Get(field string) string {
	es := e[field]

	if len(es) == 0 {
		return ""
	}

	return es[0]
}

// Has checks if form field is on post and nor empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)

	if x == "" {
		return false
	}

	return true
}
