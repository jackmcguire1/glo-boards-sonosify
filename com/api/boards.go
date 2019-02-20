package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jackmcguire1/glo-boards-sonosify/utils"
)

var boardFields = []string{
	"archived_columns",
	"archived_date",
	"columns",
	"created_by",
	"created_date",
	"invited_members",
	"labels",
	"members",
	"name",
}

type BoardsResp struct {
	HasMore bool
	Boards  []*Board
}

type BoardMember struct{
	ID       string
	Role     string
	Username string
}
type Board struct {
	ID      string
	Name    string
	Columns []*Column
	ArchivedColumns []*Column
	InvitedMembers []*BoardMember
	Members []*BoardMember
	ArchivedDate string
	CreatedDate  string
	CreatedBy    struct {
		ID string
	}
	Labels []struct {
		ID    string
		Name  string
		Color struct {
			R int64
			G int64
			B int64
			A int64
		}
		CreatedDate string
		CreatedBy   struct {
			ID string
		}
	}
}

func (a *API) GetBoards(
	page string,
	limit string,
	sortAsc bool,
	archived bool,
) (
	boardsResp *BoardsResp,
) {

	addr := fmt.Sprintf("%s/boards", a.BaseURI)

	q := url.Values{}
	q.Set("fields", utils.ToJSON(boardFields))
	q.Set("page", page)
	q.Set("per_page", limit)

	if !sortAsc {
		q.Set("sort", "desc")
	}
	if archived {
		q.Set("archived", "true")
	}

	boardsResp = &BoardsResp{}
	resp, headers, err := a.do(http.MethodGet, addr, nil, q)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &boardsResp.Boards)
	if err != nil {
		return
	}

	hasMore, err := strconv.ParseBool(headers.Get("has-more"))
	if err != nil {
		return
	}

	boardsResp.HasMore = hasMore

	return
}

func (a *API) GetBoard(ID string) (board *Board) {
	addr := fmt.Sprintf("%s/boards/%s", a.BaseURI, ID)

	q := url.Values{}
	q.Set("fields", utils.ToJSON(boardFields))

	board = &Board{}
	resp, _, err := a.do(http.MethodGet, addr, nil, q)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &board)

	return
}
