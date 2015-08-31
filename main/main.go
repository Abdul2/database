package main
import (

	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)


func main (){
	
	db, err := sql.Open("mysql", "user:password@(127.0.0.1:3306)/test")
	
	if err != nil {
		
		log.Panicln(err)
	}
	
	defer db.Close()
}