package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	fmt.Println(received)

	resp := make(map[string]interface{})
	borrowedBookCol := FindColl("BorrowedBook")
	result, err := borrowedBookCol.MemberDataCollection.InsertOne(context.Background(), received)
	if err != nil {
		return nil, err
	}

	memberCol := FindColl("Member")
	filter := bson.M{"memberId": issueBookRequest.MemberId}
	update := bson.M{"$push": bson.M{"borrowedBooks": received.RecordId}}
	_, err = memberCol.MemberDataCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	bookCol := FindColl("Book")
	filter = bson.M{"bookId": issueBookRequest.BookId}
	update = bson.M{"$inc": bson.M{"count": -1}}
	_, err = bookCol.MemberDataCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	resp["BorrowId"] = received.RecordId
	resp["_Id"] = result.InsertedID
	return resp, nil

}

func ReturnBookService(returnBook ReturnBook) (any, error) {

	member, err := FindMember(returnBook.MemberId)
	if err != nil {
		return nil, err
	}
	if member.Password != returnBook.Password {
		return nil, errors.New("incorrect password")
	}

	Borrow, err := FindRecord(returnBook.RecordId)
	if err != nil {
		return nil, err
	}

	if Borrow.Status {
		return nil, errors.New("book is not already returned")
	}

	Book, err := FindBook(Borrow.BookId)
	if err != nil {
		return nil, err
	}

	var index int
	for i, value := range member.BorrowedBooks {
		if value == returnBook.RecordId {
			index = i
		}
	}

	member.BorrowedBooks = append(member.BorrowedBooks[:index], member.BorrowedBooks[index+1:]...)

	memberCol := FindColl("Member")
	filter := bson.M{"memberId": member.MemberId}
	update := bson.M{"$set": bson.M{"borrowedBooks": member.BorrowedBooks}}
	_, err = memberCol.MemberDataCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	bookCol := FindColl("Book")
	filter = bson.M{"bookId": Book.BookId}
	update = bson.M{"$inc": bson.M{"count": +1}}
	_, err = bookCol.MemberDataCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	borrowCol := FindColl("BorrowedBook")
	filter = bson.M{"recordId": Borrow.RecordId}
	update = bson.M{
		"$set": bson.M{
			"status":     true,
			"fine":       CalculateFine(Borrow.IssueDate),
			"returnDate": time.Now().Truncate(24 * time.Hour),
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetProjection(bson.M{"issueDate": 1, "dueDate": 1, "returnDate": 1})

	var updateDoc bson.M
	err = borrowCol.MemberDataCollection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&updateDoc)
	if err != nil {
		return nil, err
	}

	return updateDoc, nil
}
