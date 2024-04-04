package api

// go get -u golang.org/x/crypto/...
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"todoapp/db"
	"todoapp/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginForm struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var user models.User
	db.DB.First(&user, id)
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request){
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	user.Password = hashPassword([]byte(user.Password))
	createdUser := db.DB.Create(&user)
	if err := createdUser.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := r.PathValue("id")
	db.DB.Find(&user,id)
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	} else {
		db.DB.Unscoped().Delete(&user)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User "+user.FirstName+" has been deleted"))
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	var form LoginForm
	json.NewDecoder(r.Body).Decode(&form)
	db.DB.Where("email = ?", form.Email).First(&user)
	fmt.Printf("The user request value %v", user)
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	if verifyPassword(user.Password, []byte(form.Password)) {
		//fmt.Println(verifyPassword(user.Password, []byte(form.Password)))
		tokenString, err := createToken(user.Email)
		fmt.Println(err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w,"Invalid credentials")
		}
		w.WriteHeader(http.StatusOK)
    	fmt.Fprint(w, tokenString)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w,"Invalid credentials")
	}
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
  	tokenString := r.Header.Get("Authorization")
  	if tokenString == "" {
    	w.WriteHeader(http.StatusUnauthorized)
    	fmt.Fprint(w, "Missing authorization header")
    	return
  	}
  	tokenString = tokenString[len("Bearer "):]
  
	err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
  	}
  
  fmt.Fprint(w, "Welcome to the the protected area")
}




var secretKey = []byte("e1e88324dadf0f46aac25f42de2a2278f67f854878552425d456995e1b17fcec")
func createToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, error := token.SignedString(secretKey)
	if error != nil {
		return "", error
	}
	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		return secretKey, nil	
	})

	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid login")
	}
	return nil	
}


func hashPassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func verifyPassword(hashed string, plainPassword []byte) bool {
	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
