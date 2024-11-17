package main

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MemberData struct {
	MemberDataCollection *mongo.Collection
}

func FindColl(collectionName string) MemberData {
	return MemberData{MemberDataCollection: GetCollection(collectionName)}
}

func CreateMemberService(member MemberStruct) (any, error) {

	resp := make(map[string]interface{})
	finalMember, err := SetMemberToStruct(member)
	if err != nil {
		return nil, err
	}

	mCollection := FindColl("Member")
	finalMember.CreatedAt = time.Now()
	received, err := mCollection.MemberDataCollection.InsertOne(context.Background(), finalMember)
	if err != nil {
		return nil, err
	}
	resp["_Id"] = received.InsertedID
	return resp, nil
}

func AddBookService(newBook BookStruct) (any, error) {

	resp := make(map[string]interface{})
	finalBook, err := SetBookToStruct(newBook)
	if err != nil {
		return nil, err
	}

	mCollection := FindColl("Book")
	received, err := mCollection.MemberDataCollection.InsertOne(context.Background(), finalBook)
	if err != nil {
		return nil, err
	}
	resp["_Id"] = received.InsertedID
	return resp, nil
}

func IssueBook(issueBookRequest NeedBook) (any, error) {

	member, err := FindMember(issueBookRequest.MemberId)
	if err != nil {
		return nil, err
	}
	if member.Password != issueBookRequest.Password {
		return nil, errors.New("incorrect password")
	}

	book, err := FindBook(issueBookRequest.BookId)
	if err != nil {
		return nil, err
	}

	if book.Count <= 0 {
		return nil, errors.New("book is not available right now")
	}

	received, err := getBorrowId(member.MemberId, book.BookId)
	if err != nil {
		return nil, err
	}

	resp := make(map[string]interface{})
	borrowedBookCol := FindColl("BorrowedBook")
	result, err := borrowedBookCol.MemberDataCollection.InsertOne(context.Background(), received)
	if err != nil {
		return nil, err
	}

	memberCol := FindColl("Member")
	filter := bson.M{"memberId": issueBookRequest.MemberId}
	update := bson.M{"$push": bson.M{"borrowedBooks" : received.RecordId}}
	_, err = memberCol.MemberDataCollection.UpdateOne(context.Background(),filter, update)
	if err != nil {
		return nil, err
	}

	bookCol := FindColl("Book")
	filter = bson.M{"bookId": issueBookRequest.BookId}
	update = bson.M{"$inc": bson.M{"count" : -1}}
	_, err = bookCol.MemberDataCollection.UpdateOne(context.Background(),filter, update)
	if err != nil {
		return nil, err
	}

	resp["BorrowId"] = received.RecordId
	resp["_Id"] = result.InsertedID
	return resp, nil

}


