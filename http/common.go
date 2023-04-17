// Package http declare something
// MarsDong 2023/3/31
package http

type Common struct {
	Action    Action   `json:"Action"`
	RequestID string   `json:"RequestId"`
	Language  Language `json:"Language"`
	UserID    string   `json:"UserID"`
}

type ListReq struct {
	Offset int    `json:"Offset"`
	Limit  int    `json:"Limit"`
	Order  string `json:"Order"`
	Sort   string `json:"Sort"`
	// ignoring the paging parameter returns all data, allowing only internal calls.
	getAll bool
	Common
}

func (l *ListReq) SetAll() {
	l.getAll = true
}

func (l *ListReq) EnableGetAll() bool {
	return l.getAll
}

type ReqWithIdentify struct {
	ID   int    `json:"Id"`
	Name string `json:"Name"`
	Uuid string `json:"Uuid"`
}

type RspWithError struct {
}

type EmptyRsp struct {
}

type CreateRsp struct {
	ID   int    `json:"Id,omitempty"`
	Uuid string `json:"Uuid,omitempty"`
}

type DeleteRsp struct {
	AffectedRow int64 `json:"AffectedRow"`
}

type UpdateRsp struct {
	AffectedRow int64 `json:"AffectedRow"`
}
