package gocv

func (c SeamlessCloneFlags) String() string {
	switch c {
	case NormalClone:
		return "normal-clone"
	case MixedClone:
		return "mixed-clone"
	case MonochromeTransfer:
		return "monochrome-transfer"
	}
	return ""
}
