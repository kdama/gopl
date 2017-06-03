package eval

// Equals は、expr と等しいかどうかを返します。
func (v Var) Equals(expr Expr) bool {
	switch expr := expr.(type) {
	case Var:
		return string(v) == string(expr)
	default:
		return false
	}
}

func (l literal) Equals(expr Expr) bool {
	switch expr := expr.(type) {
	case literal:
		return float64(l) == float64(expr)
	default:
		return false
	}
}

func (u unary) Equals(expr Expr) bool {
	switch expr := expr.(type) {
	case unary:
		return u.op == expr.op && u.x.Equals(expr.x)
	default:
		return false
	}
}

func (b binary) Equals(expr Expr) bool {
	switch expr := expr.(type) {
	case binary:
		return b.op == expr.op && b.x.Equals(expr.x) && b.y.Equals(expr.y)
	default:
		return false
	}
}

func (t ternary) Equals(expr Expr) bool {
	switch expr := expr.(type) {
	case ternary:
		return t.op1 == expr.op1 && t.op2 == expr.op2 &&
			t.x.Equals(expr.x) && t.y.Equals(expr.y) && t.z.Equals(expr.z)
	default:
		return false
	}
}

func (c call) Equals(expr Expr) bool {
	switch expr := expr.(type) {
	case call:
		if c.fn != expr.fn {
			return false
		} else if len(c.args) != len(expr.args) {
			return false
		}
		for i := range c.args {
			if !c.args[i].Equals(expr.args[i]) {
				return false
			}
		}
		return true
	default:
		return false
	}
}
