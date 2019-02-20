package api

type Card struct {
	id          string
	name        string
	position    int
	description struct {
		Text        string
		CreatedDate string
		UpdatedDate string
		CreatedBy   struct {
			ID string
		}
		UpdatedBy struct {
			ID string
		}
	}
	BoardID      string
	ColumnID     string
	CreatedDate  string
	UpdatedDate  string
	ArchivedDate string
	assignees    []struct {
		ID string
	}
	labels []struct {
		ID   string
		Name string
	}
	due_date           string
	CommentCount       int
	AttachmentCount    int
	CompletedTaskCount int
	TotalTaskCount     int
	CreatedBy          struct {
		ID string
	}
}

func (a *API) GetCards(boardID string) {

}

func (a *API) CreateCard(boardID, cardID string) {

}

func (a *API) GetCard(boardID, cardID string) {

}

func (a *API) DeleteCard(boardID, cardID string) {

}

func (a *API) EditCard(boardID, cardID string) {

}

func (a *API) CardsForColumn(boardID, columnID string) {

}
