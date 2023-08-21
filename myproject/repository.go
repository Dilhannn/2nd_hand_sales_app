// repository.go

package main

import (
	"context"
	"log"
	"time"

	"example.com/m/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
}

func NewRepository() *Repository {
	uri := "mongodb+srv://DilhanKosekul:Dn.636900@clusterproject.db0lfnb.mongodb.net/"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return &Repository{client: client}
}

func (r *Repository) UserExists(username string) (bool, error) {
	collection := r.client.Database("database1").Collection("Users")

	filter := bson.M{"username": username}
	count, err := collection.CountDocuments(context.Background(), filter, nil)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *Repository) CreateUser(user models.User) error {
	collection := r.client.Database("database1").Collection("Users")

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUserByUsername(username string) (*models.User, error) {
	collection := r.client.Database("database1").Collection("Users")

	filter := bson.M{"username": username}
	var user models.User
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) CreatePhoto(photo models.PhotoCreate) error {
	photo.CreatedAt = time.Now()
	collection := r.client.Database("database1").Collection("Photos")

	_, err := collection.InsertOne(context.Background(), photo)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdatePhoto(photo models.Photo) error {
	photo.CreatedAt = time.Now()
	collection := r.client.Database("database1").Collection("Photos")
	filter := bson.M{"_id": photo.ID}
	update := bson.M{"$set": bson.M{"url": photo.URL, "userid": photo.UserID, "tags": photo.Tags, "description": photo.Description, "price": photo.Price, "createdat": photo.CreatedAt}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) ListPhoto(users models.IDUserName) (error, []models.Photo) {
	collection := r.client.Database("database1").Collection("Photos")
	photos := []models.Photo{}
	filter := bson.M{"userid": users.UserID}
	photoList, err := collection.Find(context.Background(), filter)
	if err != nil {
		return err, nil
	}

	for photoList.Next(context.Background()) {
		var result bson.M
		err := photoList.Decode(&result)
		log.Println(result)
		if err != nil {
			log.Fatal(err)
		} else {

			photo := models.Photo{
				ID:          result["_id"].(primitive.ObjectID),
				UserID:      result["userid"].(string),
				Tags:        result["tags"].(string),
				URL:         result["url"].(string),
				Price:       result["price"].(string),
				Description: result["description"].(string),
			}
			photos = append(photos, photo)
		}
	}

	if err := photoList.Err(); err != nil {
		log.Fatal(err)
	}

	photoList.Close(context.Background())

	return nil, photos
}

func (r *Repository) ListCart(users models.IDUserName) (error, []models.Cart) {
	collection := r.client.Database("database1").Collection("Cart")
	photos := []models.Cart{}
	filter := bson.M{"userid": users.UserID}
	photoList, err := collection.Find(context.Background(), filter)
	if err != nil {
		return err, nil
	}

	for photoList.Next(context.Background()) {
		var result bson.M
		err := photoList.Decode(&result)
		log.Println(result)
		if err != nil {
			log.Fatal(err)
		} else {

			photo := models.Cart{
				ID:     result["_id"].(primitive.ObjectID),
				UserID: result["userid"].(string),
				Photo:  ConvertToPhoto(result["photo"].(primitive.M)),
			}
			photos = append(photos, photo)
		}
	}

	if err := photoList.Err(); err != nil {
		log.Fatal(err)
	}

	photoList.Close(context.Background())

	return nil, photos
}

func ConvertToPhoto(m primitive.M) models.Photo {
	var photo models.Photo

	// primitive.M nesnesinden gerekli alanları çıkararak ve tür dönüşümleri yaparak
	// models.Photo nesnesini doldurun
	photo.ID = m["id"].(primitive.ObjectID)
	photo.UserID = m["userid"].(string)
	photo.URL = m["url"].(string)
	photo.Tags = m["tags"].(string)
	photo.Price = m["price"].(string)
	photo.Description = m["description"].(string)

	// Diğer alanları da dönüştürün

	return photo
}

func (r *Repository) ListAllPhotos() (error, []models.Photo) {
	collection := r.client.Database("database1").Collection("Photos")
	photos := []models.Photo{}
	filter := bson.M{} // Herhangi bir filtreleme yapmadan tüm fotoğrafları getirmek için boş bir filtre kullanılıyor
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return err, nil
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		} else {
			photo := models.Photo{
				ID:          result["_id"].(primitive.ObjectID),
				UserID:      result["userid"].(string),
				Tags:        result["tags"].(string),
				URL:         result["url"].(string),
				Price:       result["price"].(string),
				Description: result["description"].(string),
			}
			photos = append(photos, photo)
		}
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return nil, photos
}
func (r *Repository) SearchPhotosByTag(tags models.Search) ([]models.Photo, error) {
	collection := r.client.Database("database1").Collection("Photos")
	filter := bson.M{"tags": bson.M{"$in": tags.Field}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var photos []models.Photo

	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		} else {
			photo := models.Photo{
				ID:          result["_id"].(primitive.ObjectID),
				UserID:      result["userid"].(string),
				Tags:        result["tags"].(string),
				URL:         result["url"].(string),
				Price:       result["price"].(string),
				Description: result["description"].(string),
			}
			photos = append(photos, photo)
		}
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return photos, nil
}

func (r *Repository) DeletePhoto(PhotoID primitive.ObjectID) error {
	collection := r.client.Database("database1").Collection("Photos")

	filter := bson.M{"_id": PhotoID}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateComment(photo models.Photo) error {
	collection := r.client.Database("database1").Collection("Photos")

	_, err := collection.InsertOne(context.Background(), photo)
	if err != nil {
		return err
	}

	return nil
} //3

func (r *Repository) GetCommentsByPhotoID(photo models.Photo) error {
	collection := r.client.Database("database1").Collection("Photos")

	_, err := collection.InsertOne(context.Background(), photo)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetAllPhotos() ([]models.Photo, error) {
	collection := r.client.Database("database1").Collection("Photos")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var photos []models.Photo
	for cursor.Next(context.Background()) {
		var photo models.Photo
		if err := cursor.Decode(&photo); err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}

	return photos, nil
}

func (r *Repository) AddToCart(cartItem models.Cart) error {
	collection := r.client.Database("database1").Collection("Cart")

	_, err := collection.InsertOne(context.Background(), cartItem)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) RemoveFromCart(PhotoID primitive.ObjectID) error {
	collection := r.client.Database("database1").Collection("Cart")
	filter := bson.M{"_id": PhotoID}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) RemoveFromFavorite(PhotoID primitive.ObjectID) error {
	collection := r.client.Database("database1").Collection("Favori")
	filter := bson.M{"_id": PhotoID}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
func (r *Repository) AddToFavorites(cartItem models.Cart) error {
	collection := r.client.Database("database1").Collection("Favori")

	_, err := collection.InsertOne(context.Background(), cartItem)
	if err != nil {
		return err
	}

	return nil
}
func (r *Repository) ListFavorite(users models.IDUserName) (error, []models.Cart) {
	collection := r.client.Database("database1").Collection("Favori")
	photos := []models.Cart{}
	filter := bson.M{"userid": users.UserID}
	photoList, err := collection.Find(context.Background(), filter)
	if err != nil {
		return err, nil
	}

	for photoList.Next(context.Background()) {
		var result bson.M
		err := photoList.Decode(&result)
		log.Println(result)
		if err != nil {
			log.Fatal(err)
		} else {

			photo := models.Cart{
				ID:     result["_id"].(primitive.ObjectID),
				UserID: result["userid"].(string),
				Photo:  ConvertToPhoto(result["photo"].(primitive.M)),
			}
			photos = append(photos, photo)
		}
	}

	if err := photoList.Err(); err != nil {
		log.Fatal(err)
	}

	photoList.Close(context.Background())

	return nil, photos
}
