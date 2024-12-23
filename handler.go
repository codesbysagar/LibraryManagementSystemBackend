package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var memberData MemberStruct
	err := json.NewDecoder(r.Body).Decode(&memberData)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	err = MemberValidator(memberData)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	fmt.Println(memberData)

	resp, err := CreateMemberService(memberData)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	res := Response{Status: "OK", Message: "Member Added Successfully", Data: resp}
	w.WriteHeader(http.StatusOK)
	rd, _ := json.Marshal(res)
	w.Write(rd)
}

func AddNewBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newBook BookStruct
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	err = BookValidator(newBook)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	fmt.Println(newBook)

	resp, err := AddBookService(newBook)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	res := Response{Status: "OK", Message: "Member Added Successfully", Data: resp}
	w.WriteHeader(http.StatusOK)
	rd, _ := json.Marshal(res)
	w.Write(rd)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var issueBookRequest NeedBook
	err := json.NewDecoder(r.Body).Decode(&issueBookRequest)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	err = RequestValidator(issueBookRequest)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}
	resp, err := IssueBook(issueBookRequest)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	res := Response{Status: "OK", Message: "Member Added Successfully", Data: resp}
	w.WriteHeader(http.StatusOK)
	rd, _ := json.Marshal(res)
	w.Write(rd)

}

func BackBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var returnBook ReturnBook
	err := json.NewDecoder(r.Body).Decode(&returnBook)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	err = ReturnValidator(returnBook)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	received, err := ReturnBookService(returnBook)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	res := Response{Status: "OK", Message: "Book Return Successfully", Data: received}
	w.WriteHeader(http.StatusOK)
	rd, _ := json.Marshal(res)
	w.Write(rd)

}

func GetAllRecordMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var member MemberStructDB
	err := json.NewDecoder(r.Body).Decode(&member)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	err = GetAllRecordMemberReqValidator(member)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	received, err := GetAllRecordMemberService(member)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	res := Response{Status: "OK", Message: "Got All Borrowed Books", Data: received}
	w.WriteHeader(http.StatusOK)
	rd, _ := json.Marshal(res)
	w.Write(rd)

}

func BookIssueRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book BookStructDB
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	err = BookIdValidator(book.BookId)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	received, err := BookIssueRecordService(book)
	if err != nil {
		res := Response{Status: "NOT-OK", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	res := Response{Status: "OK", Message: "Got All Borrowed Books", Data: received}
	w.WriteHeader(http.StatusOK)
	rd, _ := json.Marshal(res)
	w.Write(rd)
}
