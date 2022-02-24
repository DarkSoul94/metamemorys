package http

import "time"

type newMember struct {
	Name string `json:"name"`
}

type outMember struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type inpFile struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type inpFiles struct {
	MemberID uint64    `json:"member_id"`
	Files    []inpFile `json:"files"`
}

type outFile struct {
	ID    uint64    `json:"id"`
	Name  string    `json:"name"`
	Value string    `json:"value"`
	Date  time.Time `json:"date"`
}
