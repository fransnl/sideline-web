package main

type article struct {
	ID          string
	Title       string
	Description string
	ImgURL      string
	Blocks      []block
}

type tempArticle struct {
	ID          int
	ArticleID   string
	Title       string
	Description string
	ImgURL      string
	Blocks      []block
}

type img struct {
	ID          int
	Title       string
	Description string
	ImgURL      string
}

type block struct {
	ID           string
	BlockOrder   int
	BlockType    int
	ArticleID    string
	TextBlock    blockText
	ImgBlock     blockImage
	GalleryBlock blockGallery
}

// type 1
type blockText struct {
	ID        string
	BlockID   string
	Title     string
	Paragraph string
}

// type 2
type blockImage struct {
	ID      string
	BlockID string
	Img     img
}

// type 3
type blockGallery struct {
	ID      string
	BlockID string
	Images  []img
}
