package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type ConfigMain struct {
	MongoUri string
}

var Config ConfigMain

// var GenreSlice = []string{"Horror", "Romance", "Fiction", "Self-Help", "Novel", "Fantasy"}

// loading mongoURI
func LoadConfig() {
	mongoUri := os.Getenv("MONGODB")
	if mongoUri == "" {
		log.Fatal("DB URI not found")
	}
	Config.MongoUri = mongoUri
	fmt.Println(Config.MongoUri)

}

var Client *mongo.Client

func ConnectDB() {
	//connecting to MongoDB server
	var err error
	Client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(Config.MongoUri))
	if err != nil {
		log.Fatal("MONGO Connection Failed", err)
	}

	//ping to MongoDB server
	err = Client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Ping Failed", err)
	}

	if Client == nil {
		log.Fatal("Client is nil")

	}
	fmt.Println("Connected to MongoDB!")
}

// getting collection
func GetCollection(name string) *mongo.Collection {
	if Client == nil {
		ConnectDB()
	}
	return Client.Database("LMS").Collection(name)
}

func IdGenerator(digits int) (int, error) {

	if digits <= 0 {
		return -1, errors.New("digits limit is not valid. It should be greater than 0")
	}
	lowerBound := int(math.Pow10(digits - 1))
	upperBound := int(math.Pow10(digits))

	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(upperBound-lowerBound) + lowerBound, nil
}

func MemberValidator(input MemberStruct) error {
	if input.FullName == "" {
		return errors.New("fullName is missing")
	}
	if input.Contact > 9999999999 || input.Contact < 1000000000 {
		return errors.New("contact number should be of 10 digits only")
	}
	if input.Email == "" || !strings.Contains(input.Email, "@") {
		return errors.New("invalid email address")
	}
	if input.Password == "" || len(input.Password) < 8 {
		return errors.New("password cannot be empty or less than 8 digits")
	}
	return nil
}

func BookValidator(input BookStruct) error {
	if input.Title == "" {
		return errors.New("book Title missing")
	}
	if input.Author == "" {
		return errors.New("book Author missing")
	}
	if input.Genre == "" {
		return errors.New("book Genre missing")
	}
	if input.Quantity <= 0 {
		return errors.New("book Quantity cannot but equal or less than 0")
	}
	return nil
}

func RequestValidator(input NeedBook) error {
	if input.MemberId > 999999 || input.MemberId < 100000 {
		return errors.New("invalid MemberID - It must be of 6 digits")
	}
	if len(input.Password) < 8 {
		return errors.New("invalid Password - It must be greater or equal to 8 letter")
	}
	if input.BookId > 99999 || input.BookId < 10000 {
		return errors.New("invalid BookId - It must be of 5 digits")
	}
	return nil
}

func ReturnValidator(input ReturnBook) error {
	if input.MemberId > 999999 || input.MemberId < 100000 {
		return errors.New("invalid MemberID - It must be of 6 digits")
	}
	if len(input.Password) < 8 {
		return errors.New("invalid Password - It must be greater or equal to 8 letter")
	}
	if input.RecordId > 9999999 || input.RecordId < 1000000 {
		return errors.New("invalid RecordId - It must be of 7 digits")
	}
	return nil
}

func SetMemberToStruct(member MemberStruct) (MemberStructDB, error) {
	Id, err := IdGenerator(6)
	if err != nil {
		return MemberStructDB{}, err
	}

	// tempCol := NewMemberData("Member")
	// var result bson.M
	// err = tempCol.MemberDataCollection.FindOne(context.Background(),bson.M{"memberId": Id}).Decode(&result)
	// if err!= nil {
	// 	if err == mongo.ErrNoDocuments {
	// 		fmt.Println("Document")
	// 	}
	// }

	return MemberStructDB{
		FullName:      member.FullName,
		Contact:       member.Contact,
		Email:         member.Email,
		Password:      member.Password,
		BorrowedBooks: nil,
		MemberId:      Id,
	}, nil
}

func SetBookToStruct(newBook BookStruct) (BookStructDB, error) {
	Id, err := IdGenerator(5)
	if err != nil {
		return BookStructDB{}, err
	}
	return BookStructDB{
		BookId:   Id,
		Title:    newBook.Title,
		Author:   newBook.Author,
		Genre:    newBook.Genre,
		Quantity: newBook.Quantity,
		Count:    newBook.Quantity,
	}, nil
}

func FindMember(Id int) (MemberStructDB, error) {
	memberCol := FindColl("Member")
	var member MemberStructDB
	err := memberCol.MemberDataCollection.FindOne(context.Background(), bson.M{"memberId": Id}).Decode(&member)
	if err != nil {
		return MemberStructDB{}, errors.New("member not registered")
	}
	return member, nil
}

func FindBook(Id int) (BookStructDB, error) {
	bookCol := FindColl("Book")
	var book BookStructDB
	err := bookCol.MemberDataCollection.FindOne(context.Background(), bson.M{"bookId": Id}).Decode(&book)
	if err != nil {
		return BookStructDB{}, errors.New("book not found")
	}
	return book, nil
}

func FindRecord(Id int) (BorrowedBookRecord, error){
	borrowedBookCol := FindColl("BorrowedBook")
	var BorrowedBook BorrowedBookRecord
	err := borrowedBookCol.MemberDataCollection.FindOne(context.Background(), bson.M{"recordId": Id}).Decode(&BorrowedBook)
	if err != nil {
		return BorrowedBookRecord{}, errors.New("RecordId not found")
	}
	return BorrowedBook, nil
}

func getBorrowId(memberId int, bookId int) (BorrowedBookRecord, error) {

	Id, err := IdGenerator(7)
	if err != nil {
		return BorrowedBookRecord{}, err
	}
	BorrowedBook := BorrowedBookRecord{
		RecordId:  Id,
		BookId:    bookId,
		MemberId:  memberId,
		IssueDate: time.Now().Truncate(24 * time.Hour),
		DueDate:   time.Now().Truncate(24 * time.Hour).Add(7 * 24 * time.Hour),
		Fine:      0.0,
		Status:    false,
	}

	return BorrowedBook, nil
}


