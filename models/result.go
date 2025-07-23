package models

type Result struct {
	res string
}

func (r *Result) Store(s string) {
	r.res = s
}

func (r *Result) String() string {
	return r.res
}
