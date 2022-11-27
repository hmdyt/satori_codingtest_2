package crud

import (
	"context"

	"github.com/hmdyt/satori_codingtest-2/ent"
	"github.com/hmdyt/satori_codingtest-2/ent/mesuringpoint"
)

func CreateMesuringPoint(client *ent.Client, user_id int, body_mass float64) (*ent.MesuringPoint, error) {
	ctx := context.Background()
	point, err := client.MesuringPoint.Create().
		SetUserID(user_id).
		SetBodyMass(body_mass).
		Save(ctx)
	return point, err
}

func GetMesuringPoints(client *ent.Client, user_id int) ([]*ent.MesuringPoint, error) {
	ctx := context.Background()
	mesuring_points, err := client.MesuringPoint.
		Query().
		Where(mesuringpoint.UserIDEQ(user_id)).
		All(ctx)
	return mesuring_points, err
}

func DeleteMesuringPoint(client *ent.Client, mesuring_point_id int) error {
	ctx := context.Background()
	err := client.MesuringPoint.DeleteOneID(mesuring_point_id).Exec(ctx)
	return err
}
