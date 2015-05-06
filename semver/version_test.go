package semver

import "testing"

func TestNewVersion(t *testing.T) {
	v := NewVersion()
	if v == nil {
		t.Error("expected non-nil, got nil")
	}
}

func TestParseVersion(t *testing.T) {
	v, err := ParseVersion("1.2.3-beta.4+20150505.1")
	if err != nil {
		t.Errorf("expected nil, got %#q", err)
	}

	if v.Major != 1 {
		t.Errorf("expected %d, got %d", 1, v.Major)
	}
	if v.Minor != 2 {
		t.Errorf("expected %d, got %d", 1, v.Minor)
	}
	if v.Patch != 3 {
		t.Errorf("expected %d, got %d", 1, v.Patch)
	}

	if len(v.PreRelease) != 2 {
		t.Errorf("expected %d, got %d", 2, len(v.PreRelease))
		t.FailNow()
	}
	if v.PreRelease[0] != "beta" {
		t.Errorf("expected %#q, got %#q", "beta", v.PreRelease[0])
	}
	if v.PreRelease[1] != "4" {
		t.Errorf("expected %#q, got %#q", "4", v.PreRelease[1])
	}

	if len(v.Metadata) != 2 {
		t.Errorf("expected %d, got %d", 2, len(v.Metadata))
		t.FailNow()
	}
	if v.Metadata[0] != "20150505" {
		t.Errorf("expected %#q, got %#q", "20150505", v.Metadata[0])
	}
	if v.Metadata[1] != "1" {
		t.Errorf("expected %#q, got %#q", "1", v.Metadata[1])
	}
}

func TestParseVersion_InvalidFormat(t *testing.T) {
	_, err := ParseVersion("1.2-beta4")
	if err != ErrInvalidFormat {
		t.Errorf("expected %#q, got %#q", ErrInvalidFormat, err)
	}

	_, err = ParseVersion("1.02.0-beta4")
	if err != ErrInvalidFormat {
		t.Errorf("expected %#q, got %#q", ErrInvalidFormat, err)
	}

	_, err = ParseVersion("1.2.a-beta4")
	if err != ErrInvalidFormat {
		t.Errorf("expected %#q, got %#q", ErrInvalidFormat, err)
	}
}

func TestVersionCompare(t *testing.T) {
	v1 := &Version{2, 3, 5, []string{}, []string{"20150505"}}
	v2 := &Version{2, 3, 5, []string{"beta"}, []string{"20150505"}}
	v3 := &Version{2, 3, 5, []string{"beta", "7"}, []string{"20150505"}}
	v4 := &Version{2, 3, 5, []string{"beta", "9"}, []string{"20150505"}}
	v5 := &Version{2, 3, 5, []string{"beta", "9", "1"}, []string{"20150505"}}
	if v1.Compare(v1) != 0 {
		t.Errorf("expected %d, got %d", 0, v1.Compare(v1))
	}
	if v1.Compare(v2) != 1 {
		t.Errorf("expected %d, got %d", 1, v1.Compare(v2))
	}
	if v2.Compare(v1) != -1 {
		t.Errorf("expected %d, got %d", -1, v2.Compare(v1))
	}
	if v2.Compare(v3) != -1 {
		t.Errorf("expected %d, got %d", -1, v2.Compare(v3))
	}
	if v3.Compare(v4) != -1 {
		t.Errorf("expected %d, got %d", -1, v3.Compare(v4))
	}
	if v4.Compare(v3) != 1 {
		t.Errorf("expected %d, got %d", 1, v4.Compare(v3))
	}
	if v5.Compare(v4) != 1 {
		t.Errorf("expected %d, got %d", 1, v5.Compare(v4))
	}
}

func TestVersionEquals(t *testing.T) {
	v1 := &Version{2, 3, 5, []string{"beta7"}, []string{"20150505"}}
	v2 := &Version{2, 3, 5, []string{"beta8"}, []string{"20150505"}}
	v3 := &Version{2, 5, 3, []string{"beta7"}, []string{"20150505"}}
	if !v1.Equals(v1) {
		t.Error("expected true, got false")
	}
	if v1.Equals(v2) {
		t.Error("expected false, got true")
	}
	if v1.Equals(v3) {
		t.Error("expected false, got true")
	}
}

func TestVersionGreaterThan(t *testing.T) {
	v1 := &Version{2, 3, 5, []string{"beta7"}, []string{"20150505"}}
	v2 := &Version{2, 3, 5, []string{"beta8"}, []string{"20150505"}}
	v3 := &Version{2, 5, 3, []string{"beta7"}, []string{"20150505"}}
	if !v2.GreaterThan(v1) {
		t.Error("expected true, got false")
	}
	if v1.GreaterThan(v3) {
		t.Error("expected false, got true")
	}
}

func TestVersionLessThan(t *testing.T) {
	v1 := &Version{2, 3, 5, []string{"beta7"}, []string{"20150505"}}
	v2 := &Version{2, 3, 5, []string{"beta8"}, []string{"20150505"}}
	v3 := &Version{2, 5, 3, []string{"beta7"}, []string{"20150505"}}
	if !v1.LessThan(v2) {
		t.Error("expected true, got false")
	}
	if v3.LessThan(v1) {
		t.Error("expected false, got true")
	}
}

func TestVersionSlice(t *testing.T) {
	v := &Version{2, 3, 5, []string{"beta7"}, []string{"20150505"}}
	s := v.Slice()
	if len(s) != 3 {
		t.Errorf("expected %d, got %d", 3, len(s))
		t.FailNow()
	}
	if s[0] != v.Major {
		t.Errorf("expected %d, got %d", v.Major, s[0])
	}
	if s[1] != v.Minor {
		t.Errorf("expected %d, got %d", v.Minor, s[1])
	}
	if s[2] != v.Patch {
		t.Errorf("expected %d, got %d", v.Patch, s[2])
	}
}

func TestVersionString(t *testing.T) {
	v := &Version{2, 3, 5, []string{}, []string{"20150505"}}
	if v.String() != "2.3.5+20150505" {
		t.Errorf("expected %#q, got %#q", "2.3.5+20150505", v.String())
	}

	v = &Version{2, 3, 5, []string{"beta7"}, []string{}}
	if v.String() != "2.3.5-beta7" {
		t.Errorf("expected %#q, got %#q", "2.3.5-beta7", v.String())
	}

	v = &Version{2, 3, 5, []string{"beta7"}, []string{"20150505"}}
	if v.String() != "2.3.5-beta7+20150505" {
		t.Errorf("expected %#q, got %#q", "2.3.5-beta7+20150505", v.String())
	}

	v = &Version{2, 3, 5, []string{"beta7"}, []string{"+-*/"}}
	if v.String() != "2.3.5-beta7" {
		t.Errorf("expected %#q, got %#q", "2.3.5-beta7", v.String())
	}
}
