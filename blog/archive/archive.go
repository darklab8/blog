package archive

import "fmt"

type Link struct {
	OrignUrl    string
	HtmlRelpath string
	PdfRelath   string
}

func NewLink(
	original_url string,
	archived_html_filename string,
	archived_pdf_filename string,
) *Link {
	return &Link{
		OrignUrl:    original_url,
		HtmlRelpath: fmt.Sprintf("archive/%s", archived_html_filename),
		PdfRelath:   fmt.Sprintf("archive/%s", archived_pdf_filename),
	}
}
