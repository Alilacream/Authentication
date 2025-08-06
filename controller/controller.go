package controller

import (
	"Auth/database"
	"Auth/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	godotenv.Load()
}

func Greetings(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Welcome to my server")
}

var SecretKey= os.Getenv("SECRET_KEY")

// Register function that indecates the user, by using the data variable
// we can Parse the values as a json form, then we input them into the Name email
// BUT... for the password we used an extended library called golang.org crypto (still don't understand it but)
// it uses the hashing method
func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	// hashing methode:
	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: string(password),
	}
	//Storing the new post we created into the table
	//No need to increment Id it is already a PRIMARYKEY
	database.DB.Create(&user)

	return c.Status(201).JSON(user)
}

func Login(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}
	//Compares the actual userpassword, with the given password in the json format, if not the same, the err will be printed else the err == nil
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString("Unknown password")
	}
	clams := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email":   user.Email,
		"Expires": time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := clams.SignedString([]byte(SecretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Error": "Internal Server Error"
		})
	}
	return c.Status(201).JSON(user)
}
