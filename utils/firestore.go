package utils

import (
	"context"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// > This function initializes a Firestore client using the credentials in the file
// `./gcp/firebase_write.json`
//
// Args:
//   ctx: The context.Context object that will be used to control the lifetime of the Firestore client.
func initFirestore(ctx context.Context) *firestore.Client {

	sa := option.WithCredentialsFile("./gcp/firebase_write.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}

// It takes a House struct, converts it to a Firestore document, and writes it to the Firestore
// database
//
// Args:
//   houseInfo (House): The struct that contains the data to be written to Firestore.
func WriteFirestore(houseInfo House) {

	// Get a Firestore client.
	ctx := context.Background()
	client := initFirestore(ctx)

	docName := strings.ReplaceAll(houseInfo.Name, "/", " ")
	_, err := client.Collection("housePrice").Doc(docName).Set(ctx, houseInfo)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	defer client.Close()
}

// It gets a list of all the documents in the collection "housePrice" and returns a list of their
// names
//
// Returns:
//   A slice of strings
func GetDocsNameFirestore() []string {

	housesNames := []string{}

	// Get a Firestore client.
	ctx := context.Background()
	client := initFirestore(ctx)
	iter := client.Collection("housePrice").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil
		}
		housesNames = append(housesNames, doc.Ref.ID)
	}
	defer client.Close()

	return housesNames
}
