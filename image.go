package neivyfunctions

type Image struct {
	FileName string
	Place    string
	Comment  string
}

func newImage(fileName string, place string, comment string) *Image {
	return &Image{
		FileName: fileName,
		Place:    place,
		Comment:  comment,
	}
}
