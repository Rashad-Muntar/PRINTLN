package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Rashad-Muntar/println/config"
	"github.com/Rashad-Muntar/println/graph/model"
	"github.com/Rashad-Muntar/println/models"
	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Signup is the resolver for the signup field.
func (r *mutationResolver) Signup(ctx context.Context, input model.NewUser) (*models.User, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(bcryptPassword),
	}

	newUser := config.DB.Create(&user)
	if newUser.Error != nil {
		return nil, newUser.Error
	}
	return &user, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginUser) (string, error) {
	var user models.User

	config.DB.First(&user, "email = ?", input.Email)
	if user.Id == 0 {
		return "User not found", nil
	}
	checkPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if checkPass != nil {
		return "Password does not match", nil
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "Ooops try again", err
	}

	return tokenString, nil
}

// CreateJob is the resolver for the createJob field.
func (r *mutationResolver) CreateJob(ctx context.Context, input model.NewJob) (*models.Job, error) {
	job := models.Job{
		Description: *input.Description,
		File:        input.File,
	}

	newJob := config.DB.Create(&job)
	if newJob.Error != nil {
		return nil, newJob.Error
	}
	return &job, nil
}

// DeleteJob is the resolver for the deleteJob field.
func (r *mutationResolver) DeleteJob(ctx context.Context, id string) (string, error) {
	config.DB.Where("name = ?", "jinzhu").Delete(id)
	return "Job deleted", nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	config.DB.Find(&users)
	return users, nil
}

// Jobs is the resolver for the jobs field.
func (r *queryResolver) Jobs(ctx context.Context) ([]*models.Job, error) {
	panic(fmt.Errorf("not implemented: Jobs - jobs"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
