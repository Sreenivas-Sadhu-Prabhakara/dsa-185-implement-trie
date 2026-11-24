package main

import "testing"

func TestTrie(t *testing.T) {
	tr := NewTrie()
	tr.Insert("apple")
	if !tr.Search("apple") {
		t.Error("inserted word not found")
	}
	if tr.Search("app") {
		t.Error("prefix must not match as a word")
	}
	if !tr.StartsWith("app") {
		t.Error("prefix lookup failed")
	}
	tr.Insert("app")
	if !tr.Search("app") {
		t.Error("inserting the prefix as a word must make it searchable")
	}
	if tr.Search("apples") {
		t.Error("longer word must not match")
	}
	if tr.StartsWith("b") {
		t.Error("unrelated prefix matched")
	}
	if !tr.StartsWith("") {
		t.Error("empty prefix matches everything")
	}
}
