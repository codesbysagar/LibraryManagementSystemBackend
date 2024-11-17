package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MemberStruct struct {
	FullName string `json:"fullname,omitempty"`
	Contact  int    `json:"contact,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type MemberStructDB struct {
	Id            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MemberId      int                `json:"memberId,omitempty" bson:"memberId,omitempty"`
	FullName      string             `json:"fullName,omitempty" bson:"fullName,omitempty"`
	Contact       int                `json:"contact,omitempty" bson:"contact,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty"`
	Password      string             `json:"password,omitempty" bson:"password,omitempty"`
	BorrowedBooks []int              `json:"borrowedBooks,omitempty" bson:"borrowedBooks,omitempty"`
	CreatedAt     time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

type BookStruct struct {
	Title    string `json:"title,omitempty"`
	Author   string `json:"author,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
	Genre    string `json:"genre,omitempty"`
}

type BookStructDB struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BookId   int                `json:"bookId,omitempty" bson:"bookId,omitempty"`
	Title    string             `json:"title,omitempty" bson:"title,omitempty"`
	Author   string             `json:"author,omitempty" bson:"author,omitempty"`
	Quantity int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Genre    string             `json:"genre,omitempty" bson:"genre,omitempty"`
	Count    int                `json:"count,omitempty" bson:"count,omitempty"`
}

type BorrowedBookRecord struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RecordId   int                `json:"recordId,omitempty" bson:"recordId,omitempty"`
	MemberId   int                `json:"memberId,omitempty" bson:"memberId,omitempty"`
	BookId     int                `json:"bookId,omitempty" bson:"bookId,omitempty"`
	IssueDate  time.Time          `json:"issueDate,omitempty" bson:"issueDate,omitempty"`
	DueDate    time.Time          `json:"dueDate,omitempty" bson:"dueDate,omitempty"`
	ReturnDate time.Time          `json:"returnDate,omitempty" bson:"returnDate,omitempty"`
	Fine       float32            `json:"fine,omitempty" bson:"fine,omitempty"`
	Status     bool               `json:"status,omitempty" bson:"status,omitempty"`
}

type NeedBook struct {
	MemberId int    `json:"memberId,omitempty" bson:"memberId,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	BookId   int    `json:"bookId,omitempty" bson:"bookId,omitempty"`
}

type Response struct {
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
