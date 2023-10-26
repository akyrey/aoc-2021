package main

import (
	"reflect"
	"testing"
)

func TestRetrieveConnections(t *testing.T) {
	got := retrieveConnections("./test12.txt")

	expected := map[string][]string{
		"start": {"A", "b"},
		"A":     {"c", "b", "end"},
		"b":     {"d", "end"},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected map %v, got %v", expected, got)
	}

	got = retrieveConnections("./test12_mid.txt")

	expected = map[string][]string{
		"dc":    {"end", "start", "HN"},
		"HN":    {"start", "end"},
		"start": {"kj"},
		"LN":    {"dc"},
		"kj":    {"sa", "HN", "dc"},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected map %v, got %v", expected, got)
	}

	got = retrieveConnections("./test12_large.txt")

	expected = map[string][]string{
		"fs":    {"end", "he", "DX"},
		"he":    {"DX", "WI"},
		"start": {"DX", "pj", "RW"},
		"pj":    {"DX", "he", "RW", "fs"},
		"end":    {"zg"},
		"zg":    {"sl", "pj", "RW", "he"},
		"RW":    {"he"},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected map %v, got %v", expected, got)
	}
}
