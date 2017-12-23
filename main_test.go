package main

import "testing"

func TestCheck(t *testing.T) {
	err := check("./testdata/pkgs/badpkg/...", []string{"go/ast", "go/token"})
	if err == nil {
		t.Fatal("expected an error here")
	}
	expect := "./testdata/pkgs/badpkg is using banned dependencies go/ast, go/token"
	got := err.Error()
	if expect != got {
		t.Fatalf("expected %s, got %s", expect, got)
	}
}

func TestCheckGoodCode(t *testing.T) {
	err := check("./testdata/pkgs/goodpkg/...", []string{"go/ast"})
	if err != nil {
		t.Fatal("got an unexpected error:", err)
	}
}
