package main

import (
	"testing"
)

func TestChain_Parse_Keylength_0(t *testing.T) {
	toTest := NewChain()

	toTest.Parse("Somestring")

	if len(toTest.index) != 0 {
		t.Fail()
	}
}

func TestChain_Parse_String_shorter_then_keylength(t *testing.T) {
	toTest := NewChain()

	toTest.Parse("Some")

	if len(toTest.index) != 0 {
		t.Fail()
	}
}

func TestChain_Parse_String_equal_to_keylength(t *testing.T) {
	toTest := NewChain()

	toTest.Parse("Some string")

	if len(toTest.index) != 0 {
		t.Fail()
	}
}

func TestChain_Parse_4_words(t *testing.T) {
	toTest := NewChain()

	toTest.Parse("Some string i wrote")

	if len(toTest.index) != 2 {
		t.Fail()
	}
}

func TestChain_Parse_recurring_words(t *testing.T) {
	toTest := NewChain()

	toTest.Parse("Some string i wrote and some string i wrote again and another string i did write")

	if len(toTest.index) != 2 {
		t.Fail()
	}
}