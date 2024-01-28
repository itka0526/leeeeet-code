package seventyfive

import (
	"testing"
)

func TestCaseMergeStringsAlternatelyOne(t *testing.T) {
	a := "abc"
	b := "pqr"
	e := "apbqcr"
	v := MergeAlternately(a, b)
	if v != e {
		t.Errorf("MergeAlternately(%s, %s) = %s; want %s", a, b, v, e)
	}
}

func TestCaseMergeStringsAlternatelyTwo(t *testing.T) {
	a := "ab"
	b := "pqrs"
	e := "apbqrs"
	v := MergeAlternately(a, b)
	if v != e {
		t.Errorf("MergeAlternately(%s, %s) = %s; want %s", a, b, v, e)
	}
}

func TestCaseMergeStringsAlternatelyThree(t *testing.T) {
	a := "abcd"
	b := "pq"
	e := "apbqcd"
	v := MergeAlternately(a, b)
	if v != e {
		t.Errorf("MergeAlternately(%s, %s) = %s; want %s", a, b, v, e)
	}
}

func TestGcdOfStrings1(t *testing.T) {
	a := "ABCABC"
	b := "ABC"
	e := "ABC"
	v := gcdOfStrings(a, b)
	if v != e {
		t.Errorf("gcdOfStrings(%s, %s) = %s; want %s", a, b, v, e)
	}
}

func TestGcdOfStrings2(t *testing.T) {
	a := "ABABAB"
	b := "ABAB"
	e := "AB"
	v := gcdOfStrings(a, b)
	if v != e {
		t.Errorf("gcdOfStrings(%s, %s) = %s; want %s", a, b, v, e)
	}
}

func TestGcdOfStrings3(t *testing.T) {
	a := "LEET"
	b := "CODE"
	e := ""
	v := gcdOfStrings(a, b)
	if v != e {
		t.Errorf("gcdOfStrings(%s, %s) = %s; want %s", a, b, v, e)
	}
}

func TestGcdOfStrings4(t *testing.T) {
	a := "TAUXXTAUXXTAUXXTAUXXTAUXX"
	b := "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"
	e := "TAUXX"
	v := gcdOfStrings(a, b)
	if v != e {
		t.Errorf("gcdOfStrings(%s, %s) = %s; want %s", a, b, v, e)
	}
}

func TestGcdOfStrings5(t *testing.T) {
	a := "NLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGM"
	b := "NLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGMNLZGM"
	e := "NLZGM"
	v := gcdOfStrings(a, b)
	if v != e {
		t.Errorf("gcdOfStrings(%s, %s) = %s; want %s", a, b, v, e)
	}
}

func TestGcdOfStrings6(t *testing.T) {
	a := "ABABABAB"
	b := "ABAB"
	e := "ABAB"
	v := gcdOfStrings(a, b)
	if v != e {
		t.Errorf("gcdOfStrings(%s, %s) = %s; want %s", a, b, v, e)
	}
}
