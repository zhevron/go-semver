package semver

import "testing"

func TestNewConstraint(t *testing.T) {
	c := NewConstraint()
	if c == nil {
		t.Error("expected non-nil, got nil")
	}
}

func TestParseConstraint(t *testing.T) {
	v := &Version{1, 2, 3, []string{"beta4"}, []string{}}
	c, err := ParseConstraint(">=1.2.3-beta4")
	if err != nil {
		t.Errorf("expected nil, got %#q", err)
	}

	if c.Operator != ">=" {
		t.Errorf("expected %#q, got %#q", ">=", c.Operator)
	}
	if !c.Version.Equals(v) {
		t.Errorf("expected true, got false")
	}
}

func TestParseConstraint_InvalidOperator(t *testing.T) {
	_, err := ParseConstraint("<>1.2.3-beta4")
	if err != ErrInvalidOperator {
		t.Errorf("expected %#q, got %#q", ErrInvalidOperator, err)
	}
}

func TestParseConstraint_InvalidFormat(t *testing.T) {
	_, err := ParseConstraint(">=1.2-beta4")
	if err != ErrInvalidFormat {
		t.Errorf("expected %#q, got %#q", ErrInvalidFormat, err)
	}
}

func TestConstraintMatch(t *testing.T) {
	v1 := &Version{1, 2, 3, []string{"beta4"}, []string{}}
	v2 := &Version{1, 3, 2, []string{"beta4"}, []string{}}

	c := &Constraint{"=", v1}
	if c.Match(v2) {
		t.Error("expected false, got true")
	}

	c = &Constraint{">", v1}
	if !c.Match(v2) {
		t.Error("expected true, got false")
	}

	c = &Constraint{"<", v1}
	if c.Match(v2) {
		t.Error("expected false, got true")
	}

	c = &Constraint{">=", v1}
	if !c.Match(v2) {
		t.Error("expected true, got false")
	}

	c = &Constraint{"<=", v1}
	if c.Match(v2) {
		t.Error("expected false, got true")
	}

	c = &Constraint{"<>", v1}
	if c.Match(v2) {
		t.Error("expected false, got true")
	}
}
