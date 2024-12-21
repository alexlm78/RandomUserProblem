package main

import (
   "encoding/json"
   "fmt"
   "io/ioutil"
   "net/http"
   "sync"
)

type OUser struct {
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

type APIOResponse struct {
   Results []OUser `json:"results"`
}

func mainOne() {
   urls := []string{
      "https://randomuser.me/api/?results=5000",
      "https://randomuser.me/api/?results=5000",
      "https://randomuser.me/api/?results=5000",
   }
   
   var wg sync.WaitGroup
   ch := make(chan []OUser, len(urls))
   
   for _, url := range urls {
      wg.Add(1)
      go fetchUsers(url, ch, &wg)
   }
   
   wg.Wait()
   close(ch)
   
   var allUsers []OUser
   for users := range ch {
      allUsers = append(allUsers, users...)
   }
   
   for _, user := range allUsers {
      fmt.Println(user.Name.First, user.Name.Last)
   }
}

func fetchUsers(url string, ch chan<- []OUser, wg *sync.WaitGroup) {
   defer wg.Done()
   
   resp, err := http.Get(url)
   if err != nil {
      fmt.Println(err)
      return
   }
   
   defer resp.Body.Close()
   
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      fmt.Println(err)
      return
   }
   
   var response APIOResponse
   err = json.Unmarshal(body, &response)
   if err != nil {
      fmt.Println(err)
      return
   }
   
   ch <- response.Results
}
