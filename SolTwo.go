package main

import (
   "encoding/json"
   "fmt"
   "io/ioutil"
   "log"
   "net/http"
   "sync"
)

type TUser struct {
   Gender string `json:"gender"`
   Name   struct {
      Title string `json:"title"`
      First string `json:"first"`
      Last  string `json:"last"`
   } `json:"name"`
   Location struct {
      City    string `json:"city"`
      Country string `json:"country"`
   } `json:"location"`
   Email string `json:"email"`
   Login struct {
      Username string `json:"username"`
      UUID     string `json:"uuid"`
   } `json:"login"`
}

type APITResponse struct {
   Results []TUser `json:"results"`
}

func fetchUserData(url string, wg *sync.WaitGroup, ch chan<- []TUser) {
   defer wg.Done()
   response, err := http.Get(url)
   if err != nil {
      log.Println(err)
      return
   }
   defer response.Body.Close()
   
   responseData, err := ioutil.ReadAll(response.Body)
   if err != nil {
      log.Println(err)
      return
   }
   
   var apiResponse APITResponse
   err = json.Unmarshal(responseData, &apiResponse)
   if err != nil {
      log.Println(err)
      return
   }
   
   ch <- apiResponse.Results
}

func SolTwo() {
   const totalUsers = 15000
   const batchSize = 5000
   var allUsers []TUser
   var wg sync.WaitGroup
   ch := make(chan []TUser, totalUsers/batchSize)
   
   for i := 0; i < totalUsers/batchSize; i++ {
      wg.Add(1)
      go fetchUserData("https://randomuser.me/api/1.4/?results=5000", &wg, ch)
   }
   
   go func() {
      wg.Wait()
      close(ch)
   }()
   
   for users := range ch {
      allUsers = append(allUsers, users...)
   }
   
   for _, user := range allUsers {
      fmt.Printf("Gender: %s, Title: %s, First Name: %s, Last Name: %s, City: %s, Country: %s, Email: %s, Username: %s, UUID: %s\n",
         user.Gender, user.Name.Title, user.Name.First, user.Name.Last, user.Location.City, user.Location.Country, user.Email, user.Login.Username, user.Login.UUID)
   }
}
