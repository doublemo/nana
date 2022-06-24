package validator

import (
	"github.com/doublemo/nana/helper"
	"github.com/doublemo/nana/lib/dfa"
)

// IsBadWord 脏词检查
func IsBadWord(s string) bool {
	dfa := dfa.New()
	dfa.AddBadWords(helper.BadCodes())
	_, _, res := dfa.Check(s)
	return res
}
