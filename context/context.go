package seg

import "github.com/tedux/sego/lexeme"

const (
	BUFFER_SIZE             = 4096 // 默认缓冲区大小
	BUFFER_EXHAUST_CRITICAL = 100  // 缓冲区耗尽的临界值
)

type Context struct {
	buffer        []rune
	charTypes     []int
	lexemePathMap map[int]*lexeme.Path
	result        []*lexeme.Word
}
