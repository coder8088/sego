package dict

import (
	"fmt"
	"testing"
)

func TestNewDict(t *testing.T) {
	d := NewDict()
	hit := d.Match("中华")
	fmt.Println(hit)
}

func TestDict_Remove(t *testing.T) {
	data := []string{"中华", "中华人民", "华中", "中国"}
	d := NewDictWithLoadFunc(func() []string {
		return data
	})

	hit := d.Match("中华")
	fmt.Println(hit)

	d.Remove("中华")
	hit = d.Match("中华")
	fmt.Println(hit)
}
