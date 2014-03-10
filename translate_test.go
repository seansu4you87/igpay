package igpay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var translateTests = []struct {
	in  string
	out string
}{
	{
		in:  "dog",
		out: "ogday",
	},
	{
		in:  "cat",
		out: "atcay",
	},
	{
		in:  "hat",
		out: "athay",
	},
	{
		in:  "egg",
		out: "eggday",
	},
	{
		in:  "Dog",
		out: "Ogday",
	},
	{
		in:  "Egg",
		out: "Eggday",
	},
	{
		in:  "Today is gonna be a Great Day",
		out: "Odaytay isday onnagay ebay aday Reatgay Ayday",
	},
	{
		in:  "Do not cross Me You Fool",
		out: "Oday otnay rosscay Emay Ouyay Oolfay",
	},
}

func TestTranslate(t *testing.T) {
	for i, test := range translateTests {
		actual := Translate(test.in)
		assert.Equal(t, test.out, actual, "Test %d", i)
	}
}

func TestIsCapitalized(t *testing.T) {
	assert.Equal(t, isCapitalized("t"), false)
	assert.Equal(t, isCapitalized("T"), true)
	assert.Equal(t, isCapitalized("1"), false)
	assert.Equal(t, isCapitalized("Test"), true)
	assert.Equal(t, isCapitalized("test"), false)
	assert.Equal(t, isCapitalized("tEST"), false)
	assert.Equal(t, isCapitalized("R2D2"), true)
	assert.Equal(t, isCapitalized("c3PO"), false)
}

func TestCapitalize(t *testing.T) {
	assert.Equal(t, capitalize("t"), "T")
	assert.Equal(t, capitalize("T"), "T")
	assert.Equal(t, capitalize("1"), "1")
	assert.Equal(t, capitalize("Test"), "Test")
	assert.Equal(t, capitalize("test"), "Test")
	assert.Equal(t, capitalize("tEST"), "TEST")
	assert.Equal(t, capitalize("R2D2"), "R2D2")
	assert.Equal(t, capitalize("c3PO"), "C3PO")
}

func TestIsVowel(t *testing.T) {
	assert.Equal(t, startsWithVowel("Hello"), false)
	assert.Equal(t, startsWithVowel("hello"), false)
	assert.Equal(t, startsWithVowel("hELLO"), false)
	assert.Equal(t, startsWithVowel("HELLO"), false)
	assert.Equal(t, startsWithVowel("Ello"), true)
	assert.Equal(t, startsWithVowel("ello"), true)
	assert.Equal(t, startsWithVowel("eLLO"), true)
	assert.Equal(t, startsWithVowel("ELLO"), true)
}

func TestSuffix(t *testing.T) {
	assert.Equal(t, suffix("dog"), "day")
	assert.Equal(t, suffix("Dog"), "day")
	assert.Equal(t, suffix("home"), "hay")
	assert.Equal(t, suffix("Home"), "hay")
	assert.Equal(t, suffix("HOME"), "hay")
	assert.Equal(t, suffix("hOME"), "hay")
	assert.Equal(t, suffix("egg"), "day")
	assert.Equal(t, suffix("Egg"), "day")
	assert.Equal(t, suffix("EGG"), "day")
	assert.Equal(t, suffix("eGG"), "day")
}

func TestBase(t *testing.T) {
	assert.Equal(t, base("dog"), "og")
	assert.Equal(t, base("Dog"), "Og")
	assert.Equal(t, base("home"), "ome")
	assert.Equal(t, base("Home"), "Ome")
	assert.Equal(t, base("HOME"), "Ome")
	assert.Equal(t, base("hOME"), "ome")
	assert.Equal(t, base("egg"), "egg")
	assert.Equal(t, base("Egg"), "Egg")
	assert.Equal(t, base("EGG"), "Egg")
	assert.Equal(t, base("eGG"), "egg")
}
