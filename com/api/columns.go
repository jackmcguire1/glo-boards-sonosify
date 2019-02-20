package api

type Column struct {
	Id            string
	Name          string
	Position      int
	ArchivedDate string
	CreatedDate  string
	CreatedBy    struct {
		ID string
	}
}

func (a *API) GetColumns(boardID string) {

}

func (a *API) CreateColumn(boardID, columnID string) {

}

func (a *API) DeleteComment(boardID, cardID string) {

}
