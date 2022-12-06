package types

type Community struct {
	Name        string
	Id          string
	Description string
	PreviewURI  string
	Creator     string
}

func NewCommunity(name, id, desc, previewURI, creator string) Community {
	return Community{
		Name:        name,
		Id:          id,
		Description: desc,
		PreviewURI:  previewURI,
		Creator:     creator,
	}
}
