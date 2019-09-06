package main

import (
	"github.com/gorilla/mux"
    "fmt"
    "io/ioutil"    
    "net/http"
)

func getAllStoreItems(w http.ResponseWriter, r *http.Request)  {

	url := "https://fortnite-api.theapinetwork.com/store/get"
	method := "GET"
  
	client := &http.Client {
	  CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	  },
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Authorization", "98bb52f315fa0dea192e7f1388731da7")
	
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	
	  w.Write(body)
}

func getAllWeapons(w http.ResponseWriter, r *http.Request)  {

	url := "https://fortnite-api.theapinetwork.com/weapons/get"
	method := "GET"
  
	client := &http.Client {
	  CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	  },
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Authorization", "98bb52f315fa0dea192e7f1388731da7")
	
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	
	  w.Write(body)
}

func getAllItemsList(w http.ResponseWriter, r *http.Request)  {

	url := "https://fortnite-api.theapinetwork.com/items/list"
	method := "GET"
  
	client := &http.Client {
	  CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	  },
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Authorization", "98bb52f315fa0dea192e7f1388731da7")
	
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	
	  w.Write(body)
}

func getAllPopularItems(w http.ResponseWriter, r *http.Request)  {

	url := "https://fortnite-api.theapinetwork.com/items/popular"
	method := "GET"
  
	client := &http.Client {
	  CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	  },
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Authorization", "98bb52f315fa0dea192e7f1388731da7")
	
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	
	  w.Write(body)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/store", getAllStoreItems)
	router.HandleFunc("/api/weapons", getAllWeapons)
	router.HandleFunc("/api/items-list", getAllItemsList)
	router.HandleFunc("/api/items-list", getAllItemsList)
	router.HandleFunc("/api/popular-items-list", getAllPopularItems)
	
	http.ListenAndServe(":8000", router)

}