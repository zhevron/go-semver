// Package semver provides an implementation of Semantic Versioning.
package semver

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	// ErrInvalidFormat is returned when an invalid semver version is parsed.
	ErrInvalidFormat = errors.New("semver: version number not in semver format")
)

// Version represents a version number in semver format.
type Version struct {
	Major      int64
	Minor      int64
	Patch      int64
	PreRelease []string
	Metadata   []string
}

// NewVersion returns a new Version set to 0.1.0.
func NewVersion() *Version {
	return &Version{
		Major:      0,
		Minor:      1,
		Patch:      0,
		PreRelease: make([]string, 0),
		Metadata:   make([]string, 0),
	}
}

// ParseVersion attempts to parse a semver string according to the semver spec.
// If the string is of an invalid semver format, ErrInvalidFormat is returned.
//
// Pre-release and metadata are optional strings and have to be provided in that
// order. Ex. 1.2.3-beta4+20150505.
func ParseVersion(str string) (*Version, error) {
	v := NewVersion()

	s := strings.SplitN(str, ".", 3)
	if len(s) != 3 {
		return nil, ErrInvalidFormat
	}

	v.Metadata = strings.Split(stripAfter(&s[2], '+'), ".")
	v.PreRelease = strings.Split(stripAfter(&s[2], '-'), ".")

	n := make([]int64, 3, 3)
	for i := 0; i < 3; i++ {
		if hasLeadingZero(s[i]) {
			return nil, ErrInvalidFormat
		}
		v, err := strconv.ParseInt(s[i], 10, 64)
		if err != nil {
			return nil, ErrInvalidFormat
		}
		n[i] = v
	}

	v.Major = n[0]
	v.Minor = n[1]
	v.Patch = n[2]

	return v, nil
}

// Compare compares to Version values and returns a number representing the
// result. It can be one of the following:
//  -1 means v is less than ver
//   0 means v is equal to ver
//   1 means v is greater than ver
func (v *Version) Compare(ver *Version) int {
	v1 := []int64{v.Major, v.Minor, v.Patch}
	v2 := []int64{ver.Major, ver.Minor, ver.Patch}

	for i := 0; i < 3; i++ {
		if v1[i] != v2[i] {
			if v1[i] > v2[i] {
				return 1
			}
			return -1
		}
	}

	return v.comparePreRelease(ver)
}

// Equals checks if two Version values are equal.
func (v *Version) Equals(ver *Version) bool {
	return v.Compare(ver) == 0
}

// GreaterThan checks if a Version is greater than another.
func (v *Version) GreaterThan(ver *Version) bool {
	return v.Compare(ver) == 1
}

// LessThan checks if a Version is less than another.
func (v *Version) LessThan(ver *Version) bool {
	return v.Compare(ver) == -1
}

// Slice returns the major, minor and patch values in a slice.
func (v *Version) Slice() []int64 {
	return []int64{v.Major, v.Minor, v.Patch}
}

// String returns Version formatted as a semver string.
func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d%s%s",
		v.Major,
		v.Minor,
		v.Patch,
		v.strPreRelease(),
		v.strMetadata(),
	)
}

func (v *Version) comparePreRelease(ver *Version) int {
	if len(v.PreRelease) == 0 && len(ver.PreRelease) == 0 {
		return 0
	}
	if len(v.PreRelease) == 0 && len(ver.PreRelease) > 0 {
		return 1
	}
	if len(v.PreRelease) > 0 && len(ver.PreRelease) == 0 {
		return -1
	}

	for i := 0; i < len(v.PreRelease) && i < len(ver.PreRelease); i++ {
		b := []bool{false, false}
		n1, err := strconv.Atoi(v.PreRelease[i])
		if err == nil {
			b[0] = true
		}
		n2, err := strconv.Atoi(ver.PreRelease[i])
		if err == nil {
			b[1] = true
		}
		if b[0] && b[1] {
			if n1 > n2 {
				return 1
			}
			if n1 < n2 {
				return -1
			}
		} else {
			if v.PreRelease[i] > ver.PreRelease[i] {
				return 1
			}
			if v.PreRelease[i] < ver.PreRelease[i] {
				return -1
			}
		}
	}

	if len(v.PreRelease) == len(ver.PreRelease) {
		return 0
	}
	if len(v.PreRelease) > len(ver.PreRelease) {
		return 1
	}
	return -1
}

func (v *Version) strPreRelease() string {
	s := ""
	if len(v.PreRelease) > 0 {
		s = strings.Join(v.PreRelease, ".")
	}
	if len(s) > 0 && isAlphaNumeric(s) {
		return "-" + s
	}
	return ""
}

func (v *Version) strMetadata() string {
	s := ""
	if len(v.Metadata) > 0 {
		s = strings.Join(v.Metadata, ".")
	}
	if len(s) > 0 && isAlphaNumeric(s) {
		return "+" + s
	}
	return ""
}

func hasLeadingZero(str string) bool {
	return len(str) > 1 && str[0] == '0'
}

func isAlphaNumeric(str string) bool {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-1234567890"
	for _, ch := range str {
		if !strings.ContainsRune(chars, ch) {
			return false
		}
	}
	return true
}

func stripAfter(str *string, ch rune) string {
	s := *str
	v := ""
	if i := strings.IndexRune(s, ch); i != -1 {
		v = s[i+1:]
		*str = s[:i]
	}
	return v
}
