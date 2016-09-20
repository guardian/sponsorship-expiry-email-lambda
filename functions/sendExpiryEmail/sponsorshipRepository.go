package main

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"time"
	"strconv"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"sort"
)


type Sponsorship struct {
	Id int
	ValidFrom int64
	ValidTo int64
	Status string
	SponsorshipType string
	SponsorName string
	SponsorLogo Image
}

type Image struct {
	ImageId string
	Assets []ImageAsset
}

type ImageAsset struct {
	ImageUrl string
	Width int64
	Height int64
	MimeType string
}

type ByExpiry []Sponsorship

func (a ByExpiry) Len() int           { return len(a) }
func (a ByExpiry) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByExpiry) Less(i, j int) bool { return a[i].ValidTo < a[j].ValidTo }

func asDynamoDate(t time.Time) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{N: aws.String(strconv.FormatInt(t.Unix() * 1000, 10)) }
}

func expiringSoonQuery() *dynamodb.ScanInput {
	now := time.Now()
	plus7 := time.Duration(24 * 7) * time.Hour
	nowPlus7 := time.Now().Add(plus7)

	return &dynamodb.ScanInput{
		TableName: aws.String("tag-manager-sponsorships-CODE"),
		FilterExpression: aws.String("validTo BETWEEN :from AND :to AND #stat = :stat"),
		ExpressionAttributeNames: map[string]*string{"#stat" : aws.String("status")},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":from" : asDynamoDate(now),
			":to" : asDynamoDate(nowPlus7),
			":stat" : {S: aws.String("active") },
		},
	}
}

func expiredRecentlyQuery() *dynamodb.ScanInput {
	now := time.Now()
	minus7 := time.Duration(24 * 7) * -time.Hour
	nowMinus7 := time.Now().Add(minus7)

	return &dynamodb.ScanInput{
		TableName: aws.String("tag-manager-sponsorships-CODE"),
		FilterExpression: aws.String("validTo BETWEEN :from AND :to AND #stat = :stat"),
		ExpressionAttributeNames: map[string]*string{"#stat" : aws.String("status")},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":from" : asDynamoDate(nowMinus7),
			":to" : asDynamoDate(now),
			":stat" : {S: aws.String("expired") },
		},
	}
}

func loadSponsorships(query *dynamodb.ScanInput, dynamo *dynamodb.DynamoDB) ([]Sponsorship, error) {
	var sponsorships []Sponsorship
	var marshalError error
	err := dynamo.ScanPages(query, func(page *dynamodb.ScanOutput, lastPage bool) bool {
		for _, item := range page.Items {
			sponsorship := Sponsorship{}
			err := dynamodbattribute.UnmarshalMap(item, &sponsorship)

			if(err != nil){
				marshalError = err
				return false
			}
			sponsorships = append(sponsorships, sponsorship)
		}
		return !lastPage
	})

	if marshalError != nil {
		return nil, marshalError
	}

	if err != nil {
		return nil, err
	}


	return sponsorships, nil
}

func LoadExpiringSoon(dynamo *dynamodb.DynamoDB) ([]Sponsorship, error) {
	spons, err := loadSponsorships(expiringSoonQuery(), dynamo)

	if (err != nil) {
		return nil, err
	}
	sort.Sort(ByExpiry(spons))

	return spons, nil
}

func LoadExpiredRecently(dynamo *dynamodb.DynamoDB) ([]Sponsorship, error) {
	spons, err := loadSponsorships(expiredRecentlyQuery(), dynamo)

	if (err != nil) {
		return nil, err
	}
	sort.Sort(sort.Reverse(ByExpiry(spons)))

	return spons, nil
}