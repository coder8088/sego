package dict

const (
	mismatched = 0x00000000
	matched    = 0x00000001
	prefix     = 0x00000010
)

type Hit struct {
	state       int
	begin, end  int
	matchedNode *node
}

func (h *Hit) Matched() bool {
	return (h.state & matched) > 0
}

func (h *Hit) SetMatched() {
	h.state = h.state | matched
}

func (h *Hit) IsPrefix() bool {
	return (h.state & prefix) > 0
}

func (h *Hit) SetIsPrefix() {
	h.state = h.state | prefix
}

func (h *Hit) Mismatched() bool {
	return h.state == mismatched
}

func (h *Hit) SetMismatched() {
	h.state = mismatched
}

func (h *Hit) MatchedNode() *node {
	return h.matchedNode
}

func (h *Hit) Begin() int {
	return h.begin
}

func (h *Hit) End() int {
	return h.end
}
