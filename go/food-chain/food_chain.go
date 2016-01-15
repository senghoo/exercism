package foodchain

import "fmt"
import "strings"

const TestVersion = 1

type foodChain struct {
	noun, desc, attr string
}

var foodChains = []foodChain{
	{"fly", "", ""},
	{"spider", "", "wriggled and jiggled and tickled inside her"},
	{"bird", "How absurd to swallow a bird!", ""},
	{"cat", "Imagine that, to swallow a cat!", ""},
	{"dog", "What a hog, to swallow a dog!", ""},
	{"goat", "Just opened her throat and swallowed a goat!", ""},
	{"cow", "I don't know how she swallowed a cow!", ""},
}

func Verse(n int) string {
	if n == 0 {
		return ""
	} else if n == 8 {
		return "I know an old lady who swallowed a horse.\nShe's dead, of course!"
	}
	n -= 1

	verse := fmt.Sprintf("I know an old lady who swallowed a %s.\n%s",
		foodChains[n].noun, foodChains[n].Desc())

	for i := n; i > 0; i-- {
		verse += fmt.Sprintf("She swallowed the %s to catch the %s%s.\n",
			foodChains[i].noun, foodChains[i-1].noun, foodChains[i-1].Attr())
	}

	verse += "I don't know why she swallowed the fly. Perhaps she'll die."
	return verse
}

func Verses(start, end int) string {
	var verses = []string{}
	for v := start; v <= end; v++ {
		verses = append(verses, Verse(v))
	}
	return strings.Join(verses, "\n\n")

}

func Song() string {
	return Verses(1, 8)
}

func (f foodChain) Desc() string {
	if f.desc == "" && f.attr != "" {
		return "It " + f.attr + ".\n"
	}
	if f.desc != "" {
		return f.desc + "\n"
	}
	return ""
}

func (f foodChain) Attr() string {
	if f.attr != "" {
		return " that " + f.attr
	}
	return f.attr
}
