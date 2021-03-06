package main

import (

   //under score means that the import is not
   //included in the scope
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"fmt"
)


func main (){
	
	//always check for errors
	
	db, err := sql.Open("mysql", "user:password@(172.0.0.1:3306)/hello")
	
	if err != nil {
		
		log.Panicln(err)
	}
	
	defer db.Close()
	
		
	//create a mem
	//on-shot query
	var str string
	
	err = db.QueryRow("select * from hello.world where id=1").Scan(&str)
	
	if err != nil && err != sql.ErrNoRows{
		
		log.Panicln(err)
	}

    log.Panicln(str)
    fmt.Println(str)

}