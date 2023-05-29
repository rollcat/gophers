package gophers

import (
	"errors"
	"testing"
)

var successfulError error = errors.New("success")

func recoverSuccess(t *testing.T) {
	err := recover()
	if err != successfulError {
		t.Fail()
		panic(err)
	}
}

func TestAssertNil(t *testing.T) {
	Assert(nil)
}

func TestAssertErr(t *testing.T) {
	defer recoverSuccess(t)
	Assert(successfulError)
}

func TestMustOK(t *testing.T) {
	newInt := func() (int, error) {
		return 42, nil
	}
	i := Must(newInt())
	if i != 42 {
		t.Fail()
	}
}

func TestMustErr(t *testing.T) {
	defer recoverSuccess(t)
	newInt := func() (int, error) {
		return 42, successfulError
	}
	Must(newInt())
}

func TestMust2OK(t *testing.T) {
	newInts := func() (int, int, error) {
		return 42, 47, nil
	}
	i1, i2 := Must2(newInts())
	if i1 != 42 || i2 != 47 {
		t.Fail()
	}
}

func TestMust2Err(t *testing.T) {
	defer recoverSuccess(t)
	newInts := func() (int, int, error) {
		return 42, 47, successfulError
	}
	Must2(newInts())
}
