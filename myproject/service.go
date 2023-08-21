package main

import (
	"errors"
	"log"
	"strings"

	"example.com/m/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

var (
	UserAlreadyExistsError  = errors.New("User already exists")
	PasswordHashingError    = errors.New("Error while hashing password")
	InvalidCredentialsError = errors.New("Invalid credentials")
	PhotoNotFoundError      = errors.New("Photo not found")
	InvalidTagsError        = errors.New("Invalid tags")
	CommentCreationError    = errors.New("Failed to create comment")
)

func GenerateUUID(length int) string {
	uuid := uuid.New().String()

	uuid = strings.ReplaceAll(uuid, "-", "")

	if length < 1 {
		return uuid
	}
	if length > len(uuid) {
		length = len(uuid)
	}

	return uuid[0:length]
}

func (s *Service) CreateUser(user models.User) error {
	// Check if user already exists
	exists, err := s.Repository.UserExists(user.Username)
	if err != nil {
		return err
	}
	if exists {
		return UserAlreadyExistsError
	}

	// Save user to database
	err = s.Repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AuthenticateUser(username, password string) (*models.User, error) {
	// Get user from database by username
	user, err := s.Repository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	// Compare passwords
	match := s.CheckPasswordHash(password, user.Password)
	if !match {
		return nil, InvalidCredentialsError
	}

	return user, nil
}

func (s *Service) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(hashedPassword), nil
}

func (s *Service) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/*func (s *Service) GetPhotoByID(id string) (*models.Photo, error) {
	// Get photo from database by ID
	photo, err := s.Repository.GetPhotoByID(id)
	if err != nil {
		return nil, err
	}

	return photo, nil
}*/

func (s *Service) GetAllPhotos() ([]models.Photo, error) {
	// Get all photos from database
	photos, err := s.Repository.GetAllPhotos()
	if err != nil {
		return nil, err
	}

	return photos, nil
}

func (s *Service) CreatePhoto(photo models.PhotoCreate) error {
	// Validate tags
	/*
		if !s.ValidateTags(photo.Tags) {
			return InvalidTagsError
		}*/

	// Save photo to database
	err := s.Repository.CreatePhoto(photo)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdatePhoto(photo models.Photo) error {
	// Update photo in the database
	err := s.Repository.UpdatePhoto(photo)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ListPhoto(users models.IDUserName) (error, []models.Photo) {
	// Update photo in the database
	photos := []models.Photo{}
	err, photos := s.Repository.ListPhoto(users)
	if err != nil {
		return err, nil
	}

	return nil, photos
}

func (s *Service) ListCart(users models.IDUserName) (error, []models.Cart) {
	// Update photo in the database
	photos := []models.Cart{}
	err, photos := s.Repository.ListCart(users)
	if err != nil {
		return err, nil
	}

	return nil, photos
}

func (s *Service) ListFavorite(users models.IDUserName) (error, []models.Cart) {
	// Update photo in the database
	photos := []models.Cart{}
	err, photos := s.Repository.ListFavorite(users)
	if err != nil {
		return err, nil
	}

	return nil, photos
}

func (s *Service) ListAllPhotos() (error, []models.Photo) {
	photos := []models.Photo{}
	err, photos := s.Repository.ListAllPhotos()
	if err != nil {
		return err, nil
	}

	return nil, photos
}
func (s *Service) SearchPhotosByTag(tags models.Search) ([]models.Photo, error) {
	photos, err := s.Repository.SearchPhotosByTag(tags)
	if err != nil {
		return nil, err
	}

	return photos, nil
}

func (s *Service) DeletePhoto(PhotoID primitive.ObjectID) error {
	// Delete photo from the database
	err := s.Repository.DeletePhoto(PhotoID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AddToCart(cartItem models.Cart) error {

	err := s.Repository.AddToCart(cartItem)
	if err != nil {
		return err
	}

	return nil

}

func (s *Service) RemoveFromCart(PhotoID primitive.ObjectID) error {
	err := s.Repository.RemoveFromCart(PhotoID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoveFromFavorite(PhotoID primitive.ObjectID) error {
	err := s.Repository.RemoveFromCart(PhotoID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AddToFavorites(cartItem models.Cart) error {

	err := s.Repository.AddToFavorites(cartItem)
	if err != nil {
		return err
	}

	return nil

}
