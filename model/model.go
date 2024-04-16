package model

type User struct{
   Id int `json:"id"`
   Name string `json:"name"`
   UserName string `json:"username"`
   Email string `json:"email"`
   Adres Address  `json:"address"`
}

type Address struct{
    Street string `json:"street"`
    Suite string `json:"suite"`
    City string `json:"city"`
    Zipcode string `json:"zipcode"`

}
type Geoo struct{
   Lar string `json:"lar"`
   Lng string `json:"lng"`
}

