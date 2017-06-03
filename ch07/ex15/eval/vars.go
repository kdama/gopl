package eval

// Vars は、式に含まれる全ての変数を返します。
func (v Var) Vars() []Var {
	return []Var{v}
}

func (l literal) Vars() []Var {
	return []Var{}
}

func (u unary) Vars() []Var {
	return u.x.Vars()
}

func (b binary) Vars() []Var {
	return append(b.x.Vars(), b.y.Vars()...)
}

func (t ternary) Vars() []Var {
	return append(append(t.x.Vars(), t.y.Vars()...), t.z.Vars()...)
}

func (c call) Vars() []Var {
	vars := []Var{}
	for _, arg := range c.args {
		vars = append(vars, arg.Vars()...)
	}
	return vars
}

func appendAsSet(vars []Var, v Var) []Var {
	for _, e := range vars {
		if v == e {
			return vars
		}
	}
	return append(vars, v)
}
