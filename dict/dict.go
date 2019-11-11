package dict

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type LoadFunc func() []string

type Dict interface {
	Insert(string)
	Remove(string)
	Match(string) *Hit
	MatchWithHit(word string, currentIndex int, hit *Hit) *Hit
}

type dict struct {
	root *node
}

var defaultLoadFunc = func() (words []string) {
	filepath := "./dictionary.txt"
	dictFile, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Fail to load dictionary file: %s err: %s\n", filepath, err.Error())
	}
	defer func() { _ = dictFile.Close() }()

	log.Printf("Load dict file: %s ...\n", filepath)

	reader := bufio.NewReader(dictFile)
	var word string
	for {
		size, err := fmt.Fscanln(reader, &word)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("invalid line, err: %s", err.Error())
			break
		}
		if size == 0 {
			// 文件结束
			break
		}
		words = append(words, word)
	}

	log.Printf("Load dict file: %s compelete.\n", filepath)

	return
}

func NewDict() *dict {
	d := &dict{root: newNode()}
	words := defaultLoadFunc()
	for _, w := range words {
		d.Insert(w)
	}
	return d
}

func NewDictWithLoadFunc(loadFunc LoadFunc) *dict {
	d := &dict{root: newNode()}
	words := loadFunc()
	for _, w := range words {
		d.Insert(w)
	}
	return d
}

func (d *dict) Insert(word string) {
	cur := d.root
	chars := []rune(word)
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		if !cur.next.Contains(c) {
			cur.next.Store(c, newNode())
		}
		cur = cur.next.Get(c).(*node)
	}
	if !cur.isWord {
		cur.isWord = true
	}
}

func (d *dict) Remove(word string) {
	remove(d.root, []rune(word))
}

func remove(n *node, chars []rune) {
	if len(chars) == 1 {
		if v, ok := n.next.Load(chars[0]); ok {
			nt := v.(*node)
			if nt.isWord {
				nt.isWord = false
			}
			if nt.next.Size() == 0 {
				// 删除
				n.next.Delete(chars[0])
			}
		} else {
			return
		}
	} else {
		c := chars[0]
		if v, ok := n.next.Load(c); ok {
			nt := v.(*node)
			remove(nt, chars[1:])
		} else {
			return
		}
	}

}

func (d *dict) Match(word string) *Hit {
	chars := []rune(word)
	return d.root.match(chars, 0, len(chars), nil)
}

func (d *dict) MatchWithHit(word string, currentIndex int, hit *Hit) *Hit {
	n := hit.MatchedNode()
	if n != nil {
		return n.match([]rune(word), currentIndex, 1, hit)
	}
	return &Hit{state: mismatched}
}
