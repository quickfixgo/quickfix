package fix

type tagContainer struct {
	tag Tag
}

func (c tagContainer) Tag() Tag {
	return c.tag
}
