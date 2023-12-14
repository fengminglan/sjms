package chouxiang

// Fb 具体产品
type Fb struct {
}

func (f *Fb) Push() Push {
	return Push{
		msg:   "fb",
		error: "",
	}
}
