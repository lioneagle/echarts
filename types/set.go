package types

import (
	"sort"

	"github.com/pkg/errors"
)

type StringSet struct {
	data []string
}

func NewStringSet() *StringSet {
	return &StringSet{
		data: make([]string, 0),
	}
}

func (this *StringSet) Len() int {
	return len(this.data)
}

func (this *StringSet) Add(items ...string) {
	for _, v := range items {
		if !this.Find(v) {
			this.data = append(this.data, v)
		}
	}
}

func (this *StringSet) Find(item string) bool {
	for _, v := range this.data {
		if item == v {
			return true
		}
	}
	return false
}

func (this *StringSet) Foreach(f func(index int, item *string) (stop bool, err error)) error {
	if f == nil {
		return errors.Errorf("func f is nil")
	}

	for i := 0; i < len(this.data); i++ {
		stop, err := f(i, &this.data[i])
		if err != nil {
			return err
		}
		if stop {
			return nil
		}
	}
	return nil
}

func (this *StringSet) Sort() {
	sort.Slice(this.data, func(i, j int) bool {
		if len(this.data[i]) < len(this.data[j]) {
			return true
		}
		return this.data[i] < this.data[j]
	})
}
