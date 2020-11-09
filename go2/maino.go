package main

import (
    "fmt"
    "database/sql"
    _ "github.com/godror/godror"
)

func main() {
    username := "egibide";
    password := "12345Abcde";
    host := "localhost";
    //database := "egibide";
    fmt.Println("... Setting up Database Connection") 
    //egibide@localhost/egibide
    //db, err := sql.Open("goracle", username+"/"+password+"@"+host+"/"+database)
	//db, err := sql.Open("godror", "<your username>/<your password>@service_name")
	db, err := sql.Open("godror", username+"/"+password+"@"+host)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()
      
      
    rows,err := db.Query("select sysdate from dual")
    if err != nil {
        fmt.Println("Error running query")
        fmt.Println(err)
        return
    }
    defer rows.Close()
  
    var thedate string
    for rows.Next() {
  
        rows.Scan(&thedate)
    }
    fmt.Printf("The date is: %s\n", thedate)

}