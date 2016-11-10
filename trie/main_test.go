package main

import (
	"testing"
	"unicode/utf8"
)

func TestAll(t *testing.T) {
	// Pulled from my local machine:
	// shuf /usr/share/dict/american-english | head -10 | sed -e 's/^/"/g' -e 's/$/",/g'
	// Some Chinese thrown in there to test locales.
	dict := []string{
		"scollop",
		"nudists",
		"boringly",
		"aloha",
		"obligation",
		"weaker",
		"croak's",
		"tripling",
		"falsely",
		"pigsty",
		"pig",
		"虽然",
		"字典",
		"的",
		"地位",
		"很",
		"大",
		"大便",
		"程度",
		"网络",
	}
	d := newTrie()
	// Add every word to our trie.
	for _, w := range dict {
		d.add(w)
	}
	type test struct {
		word       string
		shouldFind bool
	}
	tests := []test{
		test{"Beethoven", false},
		test{"croak's", true},
		test{"c", false},
		test{"pig", true},
		test{"piggy", false},
		test{"pigsty", true},
		test{"字典", true},
		test{"大便", true},
		test{"大", true},
		test{"子", false},
		test{"典", false},
	}
	for _, tst := range tests {
		found, got := d.search(tst.word)
		if found && !tst.shouldFind {
			t.Errorf("Found %s, but shouldn't have", tst.word)
			if exp := utf8.RuneCountInString(tst.word) + 1; exp != got {
				t.Errorf("Should have performed %d lookups, but performed %d", exp, got)
			}
		} else if !found && tst.shouldFind {
			t.Errorf("Failed to find %s", tst.word)
		}
	}
}
