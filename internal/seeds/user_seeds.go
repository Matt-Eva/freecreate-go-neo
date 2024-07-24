package seeds

import (
	"context"
	"fmt"
	"freecreate/internal/err"
	"freecreate/internal/models"
	"freecreate/internal/queries"
	"os"
	"strconv"

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

	masterUser, mErr := createMasterUser()
	if mErr.E != nil {
		return mErr
	}

	createQuery := `
		CREATE (u:User $userParams)
		SET u.masterUser = true
		RETURN u.username AS username
	`
	userParams := map[string]any{
		"userParams": queries.NeoParamsFromStruct(masterUser),
	}
	result, cErr := neo4j.ExecuteQuery(ctx, neo, createQuery, userParams, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
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

func createMasterUser() (models.User, err.Error) {
	p := models.PostedUser{
		DisplayName:          "Matt",
		Username:             "Matt",
		Email:                faker.Email(),
		Password:             os.Getenv("SEED_USER_PASSWORD"),
		PasswordConfirmation: os.Getenv("SEED_USER_PASSWORD"),
		BirthYear:            2000,
		BirthMonth:           1,
		BirthDay:             1,
	}

	u, vErr := models.MakeNewUser(p)
	if vErr.E != nil {
		return models.User{}, vErr
	}

	return u, err.Error{}
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
	user, sErr := makeSeedUser()
	if sErr.E != nil {
		return sErr
	}
	createQuery := `
		CREATE (u:User $userParams)
		SET u.seed = true
		RETURN u AS user
	`
	userParams := map[string]any{
		"userParams": queries.NeoParamsFromStruct(user),
	}
	result, qErr := neo4j.ExecuteQuery(ctx, neo, createQuery, userParams, neo4j.EagerResultTransformer, neo4j.ExecuteQueryWithDatabase("neo4j"))
	if qErr != nil {
		e := err.NewFromErr(qErr)
		return e
	}
	if len(result.Records) < 1 {
		return err.New("no record returned from database for seeded user")
	}

	return err.Error{}
}

func makeSeedUser() (models.User, err.Error) {
	password := faker.Password()
	birthYear, cErr := strconv.Atoi(faker.YearString())
	if cErr != nil {
		e := err.NewFromErr(cErr)
		return models.User{}, e
	}
	p := models.PostedUser{
		DisplayName:          faker.Name(),
		Username:             faker.Username(),
		Email:                faker.Email(),
		Password:             password,
		PasswordConfirmation: password,
		BirthYear:            birthYear,
		BirthMonth:           1,
		BirthDay:             1,
	}

	u, vErr := models.MakeNewUser(p)
	if vErr.E != nil {
		return u, vErr
	}

	return u, err.Error{}
}
