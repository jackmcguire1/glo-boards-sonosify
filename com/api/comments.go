package api

type Comment struct{
	ID	string
	CardID	string
	BoardID	string
	CreatedDate	string
	UpdatedDate	string
	createdBy struct{
		ID string
	}
	UpdatedBy struct{
		ID string
	}
	Text	string
}

func (a *API) GetComments(boardID string, cardID string){

}

func (a *API) CreateComment(boardID, cardID string){

}

func (a *API) EditComment(boardID, cardID string){

}

func (a *API) DeleteComment(boardID, cardID string){

}
