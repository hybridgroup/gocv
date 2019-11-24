package contrib

func (c BlockMeanHashMode) String() string {
	switch c {
	case BlockMeanHashMode0:
		return "block-mean-hash-mode0"
	case BlockMeanHashMode1:
		return "block-mean-hash-mode1"
	}
	return ""
}
