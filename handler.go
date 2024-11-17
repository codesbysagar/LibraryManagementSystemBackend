package main

import (
	"encoding/json"
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
