package chouxiang

// Xhs 具体产品
type Xhs struct {
}

func (x *Xhs) Push() Push {
	return Push{
		msg:   "xhs",
		error: "",
	}
}
