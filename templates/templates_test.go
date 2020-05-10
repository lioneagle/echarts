package templates

import (
	"fmt"
	//"net/http"
	"testing"

	"github.com/lioneagle/goutil/src/test"

	"github.com/gobuffalo/packr"
)

func TestInitTemplatesOk(t *testing.T) {
	for i, v := range tpls {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			tpl, ok := GetTemplate(v)

			test.EXPECT_TRUE(t, ok, "")
			test.EXPECT_TRUE(t, len(tpl) > 0, "tpl \"%s\" is empty", v)
		})
	}
}

func TestGetTemplatesNOk(t *testing.T) {
	testdata := []string{
		"x1", "x2",
	}
	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			_, ok := GetTemplate(v)

			test.EXPECT_FALSE(t, ok, "")
		})
	}
}

func TestLoadTemplatesOk(t *testing.T) {
	box := packr.NewBox(".")

	err := LoadTemplates(box)
	test.ASSERT_EQ(t, err, nil, "")

	for i, v := range tpls {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			tpl, ok := GetTemplate(v)

			test.EXPECT_TRUE(t, ok, "")
			test.EXPECT_TRUE(t, len(tpl) > 0, "tpl \"%s\" is empty", v)
		})
	}
}

func TestLoadTemplatesNOk(t *testing.T) {
	testdata := []string{
		"x1", "x2",
	}

	box := packr.NewBox(".")

	err := loadTemplates(testdata, box)
	test.ASSERT_NE(t, err, nil, "")
}
