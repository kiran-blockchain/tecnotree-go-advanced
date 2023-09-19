package lists

type Rectangle struct{
	Length int
	Width int
}

func (r Rectangle) Area () int{
	return r.Length*r.Width
}