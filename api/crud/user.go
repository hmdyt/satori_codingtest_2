package crud

import (
	"context"

	"github.com/hmdyt/satori_codingtest-2/ent"
)

func CreateUser(client *ent.Client, name, email string) (*ent.User, error) {
	ctx := context.Background()
	user, err := client.User.Create().
		SetName(name).
		SetEmail(email).
		Save(ctx)
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func GetUser(client *ent.Client, user_id int) *ent.User {
	ctx := context.Background()
	user, _ := client.User.Get(ctx, user_id)
	return user
}
