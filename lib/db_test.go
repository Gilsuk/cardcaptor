package lib_test

import (
	"cardcaptor/lib"
	"testing"
)

func TestKeywordInsert(t *testing.T) {

	keyword := lib.Keyword{
		ID:   9999,
		Slug: "slug",
		Name: "name",
		Ref:  "ref",
		Text: "text",
	}

	err := keyword.Insert(db)

	if err != nil {
		t.Errorf("Insert Fail on %+v\n%w", keyword, err)
	}

	keyword.Delete(db)
}
