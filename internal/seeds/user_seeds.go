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
	m, mErr := createMasterUser()
	if mErr != nil {
		return mErr
	}
	sErr := seedMasterUser(neo, ctx, m)
	if sErr != nil {
		return sErr
	}
	
	return nil
}

func seedMasterUser(neo neo4j.DriverWithContext, ctx context.Context, params map[string]any)(error){
	existenceQuery := "MATCH (u:User {masterUser: true}) RETURN u.masterUser as masterUser"
	result, eErr := neo4j.ExecuteQuery(ctx, neo, existenceQuery, nil, neo4j.EagerResultTransformer,neo4j.ExecuteQueryWithDatabase("neo4j"))
	if eErr != nil {
		return eErr
	}
	if len(result.Records) > 0 {
		fmt.Println("master user already exists")
		return nil
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
		return errors.New("master user record not return upon seeding")
	}

	username, _ := result.Records[0].Get("username")

	fmt.Println("master user created. username:", username)

	return nil
}

func createMasterUser()(map[string]any, error){
	p := models.PostedUser{
		DisplayName: "Matt",
		Username: "Matt",
		Email: faker.Email(),
		Password: os.Getenv("SEED_USER_PASSWORD"),
		PasswordConfirmation: os.Getenv("SEED_USER_PASSWORD"),
		BirthYear: "2000",
		BirthMonth: "1",
		BirthDay: "1",
	}

	masterUser, err := p.GenerateUser()
	if err != nil{
		return map[string]any{}, err
	}

	params := masterUser.NewUserParams()

	return params, err
}

func makeSeedUser()(map[string]any, error){
	password := faker.Password()
	p := models.PostedUser{
		DisplayName: faker.Name(),
		Username: faker.Username(),
		Email: faker.Email(),
		Password: password,
		BirthYear: faker.YearString(),
		BirthMonth: faker.MonthName(),
		BirthDay: faker.DayOfMonth(),
	}
	
	u, err := p.GenerateUser()
	if err != nil{
		return map[string]any{}, err
	}

	params := u.NewUserParams()
	params["seed"] = true

	return params, nil
}
