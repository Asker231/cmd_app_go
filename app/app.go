package app

import (
	"encoding/json"
	"fmt"
	"github.com/TwiN/go-color"
	"io"
	"learn/model"
	"net/http"
)

 func App(){
	var id string
	fmt.Println("Введите id")
	fmt.Scan(&id)
	getUserById(id)
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
  _ = json.Unmarshal(result,&user)
  fmt.Println(string(color.InRedOverRed(user.Email)+ color.Bold))
  fmt.Println(string(color.InWhiteOverBlack(user.Name)+ color.Bold))
  fmt.Println(string(color.InGreenOverGreen(user.UserName)+ color.Bold))
  fmt.Println(string(color.InCyanOverCyan(user.Adres.City)+ color.Bold))
}