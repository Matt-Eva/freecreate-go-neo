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
	fmt.Println("deleting all nodes")
	dErr := DeleteUserSeeds(ctx, neo)
	if dErr.E != nil {
		return dErr
	}

	fmt.Println("seeding users")
	sErr := seedMasterUser(neo, ctx)
	if sErr.E != nil {
		return sErr
	}

	uErr := seedUsers(ctx, neo)
	if uErr.E != nil {
		return uErr
	}

	fmt.Println("users seeded")
	return err.Error{}
}

func DeleteUserSeeds(ctx context.Context, neo neo4j.DriverWithContext) err.Error {
	deleteQuery := `
		MATCH(u:User)
		WHERE u.seed = true
		DETACH DELETE u
	`

	_, eErr := neo4j.ExecuteQuery(ctx, neo, deleteQuery, nil, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if eErr != nil {
		e := err.NewFromErr(eErr)
		return e
	}

	fmt.Println("seed nodes deleted")

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
	fmt.Println("user nodes seeded")
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
		RETURN u AS user
	`
	result, qErr := neo4j.ExecuteQuery(ctx, neo, createQuery, params, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if qErr != nil {
		e := err.NewFromErr(qErr)
		return e
	}
	if len(result.Records) < 1 {
		return err.New("no record returned from database for seeded user")
	}

	// user := result.Records[0].AsMap()
	// fmt.Println("seed user created", user)

	return err.Error{}
}

func makeSeedUser() (map[string]any, err.Error) {
	password := faker.Password()
	p := models.PostedUser{
		DisplayName:          faker.Name(),
		Username:             faker.Username(),
		Email:                faker.Email(),
		Password:             password,
		PasswordConfirmation: password,
		BirthYear:            faker.YearString(),
		BirthMonth:           "1",
		BirthDay:             "1",
	}

	u, gErr := p.GenerateUser()
	if gErr.E != nil {
		return map[string]any{}, gErr
	}

	params := u.NewUserParams()

	return params, err.Error{}
}
