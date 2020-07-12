package types

import (
	"sort"

	"github.com/pkg/errors"
)

type SetString struct {
	data []string
}

func NewSetString() *SetString {
	return &SetString{
		data: make([]string, 0),
	}
}

func (this *SetString) Len() int {
	return len(this.data)
}

func (this *SetString) Add(items ...string) {
	for _, v := range items {
		if !this.Find(v) {
			this.data = append(this.data, v)
		}
	}
}

func (this *SetString) Find(item string) bool {
	for _, v := range this.data {
		if item == v {
			return true
		}
	}
	return false
}

func (this *SetString) Foreach(f func(index int, item *string) (stop bool, err error)) error {
	if f == nil {
		return errors.Errorf("f is nil")
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

func (this *SetString) Sort() {
	sort.Slice(this.data, func(i, j int) bool {
		if len(this.data[i]) < len(this.data[j]) {
			return true
		}
		return this.data[i] < this.data[j]
	})
}
