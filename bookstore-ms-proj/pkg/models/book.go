package models
 
import (
	"github.com/jinzhu/gorm"
	"github.com/amrmetwaly/bookstore-ms-proj/config"
)

var db *gorm.DB

type Book struct{
	gorm.model
	Name string `gorm:""json:"name"`
	Author string `json:author`
	Publication string `json:"publication"`
}

func init(){
	config.Conenct()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func CreateBook
func GetAllBooks
func GetBookById