package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type MemberData struct {
	MemberDataCollection *mongo.Collection
}

func NewMemberData(collectionName string) MemberData{
	return MemberData{MemberDataCollection: GetCollection(collectionName)}
}

func CreateMemberService(member MemberStruct) (any, error) {

	resp := make(map[string]interface{})
	finalMember , err := SetMemberToStruct(member)
	if err!= nil {
		return nil, err
	}

	mCollection := NewMemberData("member")
	finalMember.CreatedAt = time.Now()
	received, err := mCollection.MemberDataCollection.InsertOne(context.Background(), finalMember)
	if err != nil {
		return nil, err
	}
	resp["_Id"] = received.InsertedID
	return resp, nil
}
