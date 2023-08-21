// main.go

package main

import (
	"log"
	"net/http"

	"example.com/m/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	repository := NewRepository()
	service := NewService(repository)
	api := NewApi(service)
	app := fiber.New()

	SetupApp(app, api)

	app.Listen(":3001")
}

func SetupApp(app *fiber.App, api *Api) {
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))
	app.Post("/signup", api.SignUp)
	app.Post("/signin", api.SignIn)
	app.Post("/CreatePhotoModel", api.CreatePhotoModel)
	app.Post("/DeletePhotoModel", api.DeletePhotoModel)
	app.Post("/UpdatePhotoModel", api.UpdatePhotoModel)
	app.Post("/ListPhotoModel", api.ListPhotoModel)
	app.Post("/ListAllPhotosModel", api.ListAllPhotosModel)
	app.Post("/SearchPhotosByTag", api.SearchPhotosByTag)
	app.Post("/AddToCartModel", api.AddToCartModel)
	app.Post("/RemoveFromCartModel", api.RemoveFromCartModel)
	app.Post("/AddToFavoritesModel", api.AddToFavoritesModel)
	app.Post("/ListCartModel", api.ListCartModel)
	app.Post("/ListFavoriteModel", api.ListFavoriteModel)
	app.Post("/RemoveFromFavoriteModel", api.RemoveFromFavoriteModel)

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("The server is up and running.")
	})
}

func (api *Api) SignUp(c *fiber.Ctx) error {
	var user models.User
	log.Println(user)
	if err := c.BodyParser(&user); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Check if user already exists
	exists, err := api.Service.Repository.UserExists(user.Username)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to check user existence",
		})
	}
	if exists {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "User already exists",
			"status":  http.StatusConflict,
		})
	}

	// Hash password
	hashedPassword, err := api.Service.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	user.Password = hashedPassword

	// Save user to database
	err = api.Service.CreateUser(user)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
		"status":  http.StatusOK,
	})
}

func (api *Api) SignIn(c *fiber.Ctx) error {
	var credentials models.User
	log.Println(credentials)
	if err := c.BodyParser(&credentials); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Authenticate user
	user, err := api.Service.AuthenticateUser(credentials.Username, credentials.Password)
	log.Println(credentials.Username)
	log.Println(credentials.Password)
	if err != nil {
		log.Println(err)
		log.Println(user)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
			"status":  http.StatusUnauthorized,
		})
	}

	authUser := *user // models.User türünden bir değişkene atama yapılıyor

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"status":  http.StatusOK,
		"user":    authUser.ID, // authUser kullanılıyor
	})

}

func (api *Api) CreatePhotoModel(c *fiber.Ctx) error {
	var credentials2 models.PhotoCreate
	log.Println(credentials2)
	if err := c.BodyParser(&credentials2); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	err := api.Service.CreatePhoto(credentials2)
	if err != nil {
		log.Println(err)

		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
			"status":  http.StatusUnauthorized,
		})
	}

	photo := &credentials2 // models.User türünden bir değişkene atama yapılıyor

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Create Photo successful",
		"status":  http.StatusOK,
		"photo":   photo,
	})
}

func (api *Api) DeletePhotoModel(c *fiber.Ctx) error {
	var credentials2 models.IDName
	log.Println(credentials2)
	if err := c.BodyParser(&credentials2); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	err := api.Service.DeletePhoto(credentials2.ID)
	if err != nil {
		log.Println(err)

		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials2",
			"status":  http.StatusUnauthorized,
		})
	}

	photo := &credentials2

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Delete successful",
		"status":  http.StatusOK,
		"photo":   photo,
	})
}
func (api *Api) UpdatePhotoModel(c *fiber.Ctx) error {
	var credentials2 models.Photo
	log.Println(credentials2)
	if err := c.BodyParser(&credentials2); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	err := api.Service.UpdatePhoto(credentials2)
	if err != nil {
		log.Println(err)

		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Update",
			"status":  http.StatusUnauthorized,
		})
	}

	photo := &credentials2 // models.User türünden bir değişkene atama yapılıyor

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Update Photo successful",
		"status":  http.StatusOK,
		"photo":   photo,
	})
}

func (api *Api) ListPhotoModel(c *fiber.Ctx) error {
	var users models.IDUserName

	if err := c.BodyParser(&users); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	log.Println(&users)
	photos := []models.Photo{}
	err, photos := api.Service.ListPhoto(users)
	if err != nil {
		log.Println(err)

		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Update",
			"status":  http.StatusUnauthorized,
		})
	}

	list := &photos // models.User türünden bir değişkene atama yapılıyor

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Update Photo successful",
		"status":  http.StatusOK,
		"list":    list,
	})
}

func (api *Api) ListAllPhotosModel(c *fiber.Ctx) error {
	photos := []models.Photo{}
	err, photos := api.Service.ListAllPhotos()
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"status":  http.StatusInternalServerError,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "List all photos successful",
		"status":  http.StatusOK,
		"list":    photos,
	})
}

func (api *Api) AddToCartModel(c *fiber.Ctx) error {
	var cartItem models.Cart

	if err := c.BodyParser(&cartItem); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	err := api.Service.AddToCart(cartItem)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"status":  http.StatusInternalServerError,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Item added to cart successfully",
		"status":  http.StatusOK,
		"item":    cartItem,
	})
}

func (api *Api) AddToFavoritesModel(c *fiber.Ctx) error {
	var cartItem models.Cart

	if err := c.BodyParser(&cartItem); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	err := api.Service.AddToFavorites(cartItem)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"status":  http.StatusInternalServerError,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Item added to cart successfully",
		"status":  http.StatusOK,
		"item":    cartItem,
	})
}

func (api *Api) ListCartModel(c *fiber.Ctx) error {
	var users models.IDUserName

	if err := c.BodyParser(&users); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	log.Println(&users)
	photos := []models.Cart{}
	err, photos := api.Service.ListCart(users)
	if err != nil {
		log.Println(err)

		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Update",
			"status":  http.StatusUnauthorized,
		})
	}

	list := &photos // models.User türünden bir değişkene atama yapılıyor

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Update Photo successful",
		"status":  http.StatusOK,
		"list":    list,
	})
}

func (api *Api) ListFavoriteModel(c *fiber.Ctx) error {
	var users models.IDUserName

	if err := c.BodyParser(&users); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	log.Println(&users)
	photos := []models.Cart{}
	err, photos := api.Service.ListCart(users)
	if err != nil {
		log.Println(err)

		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Update",
			"status":  http.StatusUnauthorized,
		})
	}

	list := &photos // models.User türünden bir değişkene atama yapılıyor

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Listed Photo successful",
		"status":  http.StatusOK,
		"list":    list,
	})
}

func (api *Api) RemoveFromCartModel(c *fiber.Ctx) error {
	var cartItem models.IDName

	if err := c.BodyParser(&cartItem); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	log.Println(cartItem)

	err := api.Service.RemoveFromCart(cartItem.ID)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"status":  http.StatusInternalServerError,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Item removed from cart successfully",
		"status":  http.StatusOK,
		"item":    cartItem,
	})
}

func (api *Api) RemoveFromFavoriteModel(c *fiber.Ctx) error {
	var cartItem models.IDName

	if err := c.BodyParser(&cartItem); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	log.Println(cartItem)

	err := api.Service.RemoveFromCart(cartItem.ID)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"status":  http.StatusInternalServerError,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Item removed from cart successfully",
		"status":  http.StatusOK,
		"item":    cartItem,
	})
}

func (api *Api) SearchPhotosByTag(c *fiber.Ctx) error {
	var tags models.Search

	if err := c.BodyParser(&tags); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	log.Println(tags)

	photos, err := api.Service.SearchPhotosByTag(tags)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"status":  http.StatusInternalServerError,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Search successful",
		"status":  http.StatusOK,
		"photos":  photos,
	})
}
