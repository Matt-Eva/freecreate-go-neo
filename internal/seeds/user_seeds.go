package seeds

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/models"
	"os"

	"github.com/go-faker/faker/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func SeedUsers(neo neo4j.DriverWithContext, ctx context.Context) err.Error {
	sErr := seedMasterUser(neo, ctx)
	if sErr.E != nil {
		return sErr
	}

	uErr := seedUsers(ctx, neo)
	if uErr.E != nil {
		return uErr
	}

	return err.Error{}
}

func seedMasterUser(neo neo4j.DriverWithContext, ctx context.Context) err.Error {

	existenceQuery := "MATCH (u:User {masterUser: true}) RETURN u.username AS username"

	result, eErr := neo4j.ExecuteQuery(ctx, neo, existenceQuery, nil, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if eErr != nil {
		e := err.NewFromErr(eErr)
		return e
	}
	if len(result.Records) > 0 {
		username, ok := result.Records[0].Get("username")
		if ok {
			fmt.Println("master user already exists; name: ", username)
		} else {
			return err.New("master user record does not have username field")
		}
		return err.Error{}
	}

	params, mErr := createMasterUser()
	if mErr.E != nil {
		return mErr
	}

	createQuery := `
		CREATE (u:User $userParams)
		SET u.masterUser = true
		RETURN u.username AS username
	`
	result, cErr := neo4j.ExecuteQuery(ctx, neo, createQuery, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if cErr != nil {
		e := err.NewFromErr(cErr)
		return e
	}
	if len(result.Records) < 1 {
		return err.New("master user record not returned upon seeding")
	}

	username, _ := result.Records[0].Get("username")
	fmt.Println("master user created. username:", username)

	return err.Error{}
}

func createMasterUser() (map[string]any, err.Error) {
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

	masterUser, gErr := p.GenerateUser()
	if gErr.E != nil {
		return map[string]any{}, gErr
	}

	params := masterUser.NewUserParams()

	return params, err.Error{}
}

func seedUsers(ctx context.Context, neo neo4j.DriverWithContext) err.Error {
	for i := 0; i < 100; i++ {
		sErr := seedUser(ctx, neo)
		if sErr.E != nil {
			return sErr
		}
	}
	return err.Error{}
}

func seedUser(ctx context.Context, neo neo4j.DriverWithContext) err.Error {
	params, sErr := makeSeedUser()
	if sErr.E != nil {
		return sErr
	}
	createQuery := `
		CREATE (u:User $userParams)
		SET u.seed = true
		RETURN u.username AS username
	`
	result, qErr := neo4j.ExecuteQuery(ctx, neo, createQuery, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if qErr != nil {
		e := err.NewFromErr(qErr)
		return e
	}
	if len(result.Records) < 1 {
		return err.New("now record returned from database for seeded user")
	}
	username, ok := result.Records[0].Get("username")
	if ok {
		fmt.Println("seed user created: name:", username)
	} else {
		return err.New("seed user record does not have username property")
	}

	return err.Error{}
}

func makeSeedUser() (map[string]any, err.Error) {
	password := faker.Password()
	p := models.PostedUser{
		DisplayName: faker.Name(),
		Username:    faker.Username(),
		Email:       faker.Email(),
		Password:    password,
		PasswordConfirmation: password,
		BirthYear:   faker.YearString(),
		BirthMonth:  faker.MonthName(),
		BirthDay:    faker.DayOfMonth(),
	}

	u, gErr := p.GenerateUser()
	if gErr.E != nil {
		return map[string]any{}, gErr
	}

	params := u.NewUserParams()

	return params, err.Error{}
}
