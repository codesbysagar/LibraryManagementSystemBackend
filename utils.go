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

// 1,00,000 --> 8,99,999 + 1,00,000 --> 18,99,999
// 10,00,000 99,99,999
// generate random integer ID of specified digits
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
	}, nil
}
