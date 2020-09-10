package neivyfunctions

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Database struct {
	firebaseApp *App
}

func (d *Database) Open() error {
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	return err
}
