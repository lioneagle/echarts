package types

import (
	"fmt"
	//"net/http"
	"testing"

	"github.com/pkg/errors"

	"github.com/lioneagle/goutil/src/test"
)

func TestStringSetAdd(t *testing.T) {
	testdata := []struct {
		items  []string
		wanted []string
	}{
		{[]string{}, []string{}},
		{[]string{"a"}, []string{"a"}},
		{[]string{"a", "b"}, []string{"a", "b"}},
		{[]string{"a", "b", "a"}, []string{"a", "b"}},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			//set := NewStringSet()
			set := &StringSet{}
			set.Add(v.items...)

			test.EXPECT_EQ(t, set.Len(), len(v.wanted), "")
			test.EXPECT_EQ(t, set.data, v.wanted, "")
		})
	}
}

func TestStringSetForeach(t *testing.T) {
	items := []string{"a", "b", "c"}

	set := NewStringSet()
	set.Add(items...)

	count := 0
	checkFunc1 := func(index int, item *string) (stop bool, err error) {
		count++
		return false, nil
	}

	checkFunc2 := func(index int, item *string) (stop bool, err error) {
		count++
		if index >= 1 {
			return true, nil
		}
		return false, nil
	}

	checkFunc3 := func(index int, item *string) (stop bool, err error) {
		count++
		if index >= 2 {
			return false, errors.Errorf("test err")
		}
		return false, nil
	}

	checkFunc4 := func(index int, item *string) (stop bool, err error) {
		count++
		*item = "host/" + *item
		return false, nil
	}

	count = 0
	err := set.Foreach(nil)
	test.EXPECT_TRUE(t, err != nil, "")

	count = 0
	err = set.Foreach(checkFunc1)
	test.EXPECT_TRUE(t, err == nil, "err = %+v", err)
	test.EXPECT_EQ(t, count, len(items), "")

	count = 0
	err = set.Foreach(checkFunc2)
	test.EXPECT_TRUE(t, err == nil, "err = %+v", err)
	test.EXPECT_EQ(t, count, 2, "")

	count = 0
	err = set.Foreach(checkFunc3)
	test.EXPECT_TRUE(t, err != nil, "")
	test.EXPECT_EQ(t, count, 3, "")

	count = 0
	err = set.Foreach(checkFunc4)
	test.EXPECT_TRUE(t, err == nil, "")
	test.EXPECT_EQ(t, count, 3, "")
	test.EXPECT_EQ(t, set.data, []string{"host/a", "host/b", "host/c"}, "")
}

func TestStringSetSort(t *testing.T) {
	testdata := []struct {
		items  []string
		wanted []string
	}{
		{[]string{"a", "c", "b"}, []string{"a", "b", "c"}},
		{[]string{"ab", "c", "b"}, []string{"b", "c", "ab"}},
		{[]string{"a", "c", "b", "f", "e"}, []string{"a", "b", "c", "e", "f"}},
	}

	for i, v := range testdata {
		v := v

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			set := NewStringSet()
			set.Add(v.items...)
			set.Sort()

			test.EXPECT_EQ(t, set.data, v.wanted, "")
		})
	}
}
