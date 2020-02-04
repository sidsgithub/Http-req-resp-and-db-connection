package main
 
  import (

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
  		"fmt"
       "io/ioutil"
		 "log"
		"encoding/json"
       "net/http"
   )
  func requestResponse(w http.ResponseWriter, r *http.Request){
	type User struct{
		Email string
		Password string
	}
      if r.URL.Path != "/" {
			 http.NotFound(w, r)  }
      switch r.Method {
      case "GET":
              for k, v := range r.URL.Query() {
                  fmt.Printf("%s: %s\n", k, v)
          }
              w.Write([]byte("Received a GET request\n"))
     case "POST":
              reqBody, err := ioutil.ReadAll(r.Body)
              if err != nil {
                       log.Fatal(err)
              }
			
			  fmt.Printf("%s\n", reqBody)
			  jsonBody := reqBody
			  var user User
				json.Unmarshal([]byte(jsonBody), &user)
				fmt.Printf("email :  %s password :  %s",user.Email,user.Email)
				db, err := sql.Open("mysql","root:password@tcp(localhost:3306)/customerOne")
				row, err := db.Query("INSERT INTO user(username,password) values (?,?);",user.Email,user.Password)
				log.Println(row,"inserted")
				defer db.Close()
			  w.Write([]byte("Received a POST request\n"))
     default:
              w.WriteHeader(http.StatusNotImplemented)
              w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	 }
  }

  func main() {

      http.HandleFunc("/", requestResponse)
	  log.Fatal(http.ListenAndServe(":8000", nil))
  }