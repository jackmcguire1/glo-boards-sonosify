package api

type Attachment struct {
	ID          string
	Filename    string
	MimeType    string
	CreatedDate string
	CreatedBy   struct {
		ID string
	}
}

func (a *API) GetAttachments(boardID string, cardID string) {

}
