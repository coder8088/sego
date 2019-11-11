package dict

type node struct {
	next   *SyncMap
	isWord bool
}

func newNode() *node {
	return &node{
		next:   NewSyncMap(),
		isWord: false,
	}
}

func (n *node) match(chars []rune, begin, length int, hit *Hit) *Hit {
	if hit == nil {
		hit = &Hit{begin: begin, state: mismatched}
	} else {
		hit.SetMismatched()
	}
	hit.end = begin
	keyChar := chars[begin]
	v := n.next.Get(keyChar)

	if v != nil {
		nt := v.(*node)
		if length == 1 {
			if nt.isWord {
				hit.SetMatched()
			}
			if nt.next.Size() > 0 {
				hit.SetIsPrefix()
				hit.matchedNode = nt
			}
			return hit
		} else if length > 1 {
			return nt.match(chars, begin+1, length-1, hit)
		}
	}
	// 没有匹配到
	return hit
}
