package main

import (
	"fmt"
	"io"
	"net/http"
)
import "github.com/TwiN/go-color"

//%d number
//%s string
//%f float
func main(){
	var cmd int
	fmt.Println(color.InBlue("Получить данные 1 - да, 0 - нет"))
	fmt.Scan(&cmd)
	comand(cmd)
}
type Nums interface{
		int|int8|int16|int32|int64|
		uint|uint8|uint16|uint32|uint64|float32|float64
}
func comand(num int){
	switch num{
	case 0:
		break;
	case 1:
		req, _ := http.NewRequest("GET","https://jsonplaceholder.typicode.com/posts", nil)
        res, _ := http.DefaultClient.Do(req)
       		defer res.Body.Close()
      		 body, _ := io.ReadAll(res.Body)
        fmt.Println(res)
        fmt.Println(string(body))
	default:
		break;
	}
}