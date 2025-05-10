package main

var allImg []img = []img{
	{1, "Starry night", "A starry night by Van Gough.", "https://media.timeout.com/images/106006274/image.jpg"},
	{2, "Mona Lisa", "Mona Lisa by Leonardo DaVinci.", "https://media.cnn.com/api/v1/images/stellar/prod/190430171751-mona-lisa.jpg?q=w_2000,c_fill"},
	{3, "Water lilies", "Water lilies by Claude Monet.", "https://collectionapi.metmuseum.org/api/collection/v1/iiif/438008/preview"},
}

var t = blockText{"", "", "", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi venenatis viverra varius. Duis suscipit fringilla arcu quis fringilla. Phasellus fringilla porta nisi auctor porta. Cras vel nulla porta, laoreet erat eu, pellentesque diam. Aliquam ut laoreet massa, at dignissim neque."}
var t2 = blockText{"", "", "Text block med titel", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi venenatis viverra varius. Duis suscipit fringilla arcu quis fringilla. Phasellus fringilla porta nisi auctor porta. Cras vel nulla porta, laoreet erat eu, pellentesque diam. Aliquam ut laoreet massa, at dignissim neque."}
var i = blockImage{"", "", allImg[1]}
var g = blockGallery{"", "", allImg}
var b = []block{
	{"b1", 1, 1, "", t, blockImage{}, blockGallery{}},
	{"b2", 2, 1, "", t2, blockImage{}, blockGallery{}},
	{"b3", 3, 2, "", blockText{}, i, blockGallery{}},
	{"b4", 4, 1, "", t, blockImage{}, blockGallery{}},
	{"b5", 5, 3, "", blockText{}, blockImage{}, g},
}
var a1 = article{
	"a1",
	"Exempel Artikel 1",
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi venenatis viverra varius. Duis suscipit fringilla arcu quis fringilla. Phasellus fringilla porta nisi auctor porta. Cras vel nulla porta, laoreet erat eu, pellentesque diam. Aliquam ut laoreet massa, at dignissim neque.",
	"https://media4.giphy.com/media/v1.Y2lkPTc5MGI3NjExYzgxaDQzNG1xdDgxNnB5ZHhtYnpvN3RkZm10bWp5cW5nYjMxanBoMCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/dvNwhVCMyvEwSLPRem/giphy.gif",
	b,
}

var a2 = article{
	"a2",
	"Exempel Artikel 2",
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi venenatis viverra varius. Duis suscipit fringilla arcu quis fringilla. Phasellus fringilla porta nisi auctor porta. Cras vel nulla porta, laoreet erat eu, pellentesque diam. Aliquam ut laoreet massa, at dignissim neque.",
	"https://media4.giphy.com/media/v1.Y2lkPTc5MGI3NjExYzgxaDQzNG1xdDgxNnB5ZHhtYnpvN3RkZm10bWp5cW5nYjMxanBoMCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/dvNwhVCMyvEwSLPRem/giphy.gif",
	[]block{},
}

var allArticles = []article{a1, a2}

func getArticles() []article {
	return allArticles
}

func getArticle(id string) article {
	var article article

	for _, a := range allArticles {
		if a.ID == id {
			article = a
		}
	}
	return article
}
