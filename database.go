package neivyfunctions

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

const (
	url          = "https://neivy-4c6a2.firebaseio.com/"
	databasePath = "server/saving-data/neivy"
)

type Database struct {
	client      *db.Client
	firebaseApp *firebase.App
	ref         *db.Ref
}

func newDatabase() *Database {
	return &Database{}
}

func (d *Database) Open() error {
	opt := option.WithCredentialsFile("./.secrets/serviceAccountKey.json")
	ctx := context.Background()
	config := &firebase.Config{DatabaseURL: url}

	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return err
	}

	d.firebaseApp = app

	client, err := d.firebaseApp.Database(ctx)
	if err != nil {
		return err
	}

	d.client = client

	d.ref = d.client.NewRef(databasePath)
	if err := d.PushImage(newImage("test", "place", "comment")); err != nil {
		//	fmt.Errorf("database reference error: %v", err)
		return err
	}

	return nil
}

func (d *Database) PushImage(image *Image) error {
	ctx := context.Background()

	imagesRef := d.ref.Child("images")
	err := imagesRef.Set(ctx, map[string]*Image{
		image.FileName: image,
	})

	return err
}
