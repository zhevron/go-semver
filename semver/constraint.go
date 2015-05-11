// Package semver provides an implementation of Semantic Versioning.
package semver

import (
	"errors"
	"strings"
)

var (
	// ErrInvalidOperator is returned when the parser operator is not valid.
	ErrInvalidOperator = errors.New("semver: invalid operator")
)

// Constraint represents a version constraint for a semver version number.
type Constraint struct {
	Operator string
	Version  *Version
}

// NewConstraint returns a new constraint using the '=' operator and a default
// Version object. See NewVersion for the default version number.
func NewConstraint() *Constraint {
	return &Constraint{
		Operator: "=",
		Version:  NewVersion(),
	}
}

// ParseConstraint attempts to parse a semver constraint containing an operator
// and a semver string.
// If the operator is not valid, ErrInvalidOperator is returned.
// If the string is of an invalid semver format, ErrInvalidFormat is returned.
//
// Examples:
//		 =2.0.0 (Equals 2.0.0)
//		 >2.0.0 (Greater than 2.0.0)
//		 <2.0.0 (Less than 2.0.0)
//		>=2.0.0 (Greater than or equal to 2.0.0)
//		<=2.0.0 (Less than or equal to 2.0.0)
func ParseConstraint(str string) (*Constraint, error) {
	c := NewConstraint()
	c.Operator = ""

	operator := ""
	for _, r := range str {
		if !isOperator(r) {
			break
		}
		operator += string(r)
	}
	str = str[len(operator):]

	operators := []string{"=", ">", "<", ">=", "<="}
	for _, o := range operators {
		if o == operator {
			c.Operator = o
		}
	}

	if len(c.Operator) == 0 {
		return nil, ErrInvalidOperator
	}

	v, err := ParseVersion(str)
	if err != nil {
		return nil, err
	}
	c.Version = v

	return c, nil
}

// Match attempts to match the given Version with the constraint.
// Valid operators are:
//		 = means equal
//		 > means greater than
//		 < means less than
//		>= means greater than or equal
//		<= means less than or equal
func (c *Constraint) Match(v *Version) bool {
	switch c.Operator {
	case "=":
		return v.Equals(c.Version)

	case ">":
		return v.GreaterThan(c.Version)

	case "<":
		return v.LessThan(c.Version)

	case ">=":
		return v.GreaterThan(c.Version) || v.Equals(c.Version)

	case "<=":
		return v.LessThan(c.Version) || v.Equals(c.Version)
	}

	return false
}

func isOperator(ch rune) bool {
	operators := "=><"
	return strings.ContainsRune(operators, ch)
}
