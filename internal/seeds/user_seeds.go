package seeds

import (
	"context"
	"errors"
	"fmt"
	"freecreate/internal/models"
	"os"

	"github.com/go-faker/faker/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func SeedUsers(neo neo4j.DriverWithContext, ctx context.Context) error {
	sErr := seedMasterUser(neo, ctx)
	if sErr != nil {
		return sErr
	}

	return nil
}

func seedMasterUser(neo neo4j.DriverWithContext, ctx context.Context) error {

	existenceQuery := "MATCH (u:User {masterUser: true}) RETURN u.username AS username"

	result, eErr := neo4j.ExecuteQuery(ctx, neo, existenceQuery, nil, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if eErr != nil {
		return eErr
	}
	if len(result.Records) > 0 {
		username, ok := result.Records[0].Get("username")
		if ok {
			fmt.Println("master user already exists; name: ", username)
		} else {
			return errors.New("master user record does not have username field")
		}
		return nil
	}

	params, mErr := createMasterUser()
	if mErr != nil {
		return mErr
	}

	createQuery := `
		CREATE (u:User $userParams)
		SET u.masterUser = true
		RETURN u.username AS username
	`
	result, cErr := neo4j.ExecuteQuery(ctx, neo, createQuery, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if cErr != nil {
		return cErr
	}
	if len(result.Records) < 1 {
		return errors.New("master user record not returned upon seeding")
	}

	username, _ := result.Records[0].Get("username")
	fmt.Println("master user created. username:", username)

	return nil
}

func createMasterUser() (map[string]any, error) {
	p := models.PostedUser{
		DisplayName:          "Matt",
		Username:             "Matt",
		Email:                faker.Email(),
		Password:             os.Getenv("SEED_USER_PASSWORD"),
		PasswordConfirmation: os.Getenv("SEED_USER_PASSWORD"),
		BirthYear:            "2000",
		BirthMonth:           "1",
		BirthDay:             "1",
	}

	masterUser, err := p.GenerateUser()
	if err != nil {
		return map[string]any{}, err
	}

	params := masterUser.NewUserParams()

	return params, err
}

func seedUsers(ctx context.Context, neo neo4j.DriverWithContext) error {
	for i:= 0; i < 100; i++{
		err := seedUser(ctx, neo)
		if err != nil {
			return err
		}
	}
	return nil
}

func seedUser(ctx context.Context, neo neo4j.DriverWithContext) error {
	params, sErr := makeSeedUser()
	if sErr != nil {
		return sErr
	}
	createQuery := `
	CREATE (u:User $userParams)
	SET u.seed = true
	RETURN u.username AS username
	`
	result, qErr := neo4j.ExecuteQuery(ctx, neo, createQuery, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if qErr != nil {
		return qErr
	}
	if len(result.Records) < 1{
		return errors.New("now record returned from database for seeded user")
	}
	username, ok := result.Records[0].Get("username")
	if ok {
		fmt.Println("seed user created: name:", username)
	} else {
		return errors.New("seed user record does not have username property")
	}

	return nil
}

func makeSeedUser() (map[string]any, error) {
	password := faker.Password()
	p := models.PostedUser{
		DisplayName: faker.Name(),
		Username:    faker.Username(),
		Email:       faker.Email(),
		Password:    password,
		BirthYear:   faker.YearString(),
		BirthMonth:  faker.MonthName(),
		BirthDay:    faker.DayOfMonth(),
	}

	u, err := p.GenerateUser()
	if err != nil {
		return map[string]any{}, err
	}

	params := u.NewUserParams()

	return params, nil
}


