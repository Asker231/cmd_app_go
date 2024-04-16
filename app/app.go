package app

import (
	"encoding/json"
	"fmt"
	"io"
	"learn/model"
	"net/http"
)



 func app(){
  
	getUserById("3")
 
}

func getUserById(id string){
   urlPath  := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%s",id)

   var user = model.User{}

   res,err := http.Get(urlPath)
   if err != nil{
		fmt.Print(err)
   }

  result ,err := io.ReadAll(res.Body)
  if err != nil{
	fmt.Println(err)
  }

  err = json.Unmarshal(result,user)

 fmt.Println(user)
}