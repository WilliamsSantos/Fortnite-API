package main

import (
    "github.com/gorilla/mux"
    "fmt"
    "io/ioutil"    
    "net/http"
//     "github.com/urfave/negroni" // Talvez seja preciso instalar o negroni go get github.com/urfave/negroni
)

func APIMiddleware(w http.ResponseWriter, r *http.Request){
	
}

func getAllStoreItems(w http.ResponseWriter, r *http.Request)  {

	url := "https://fortnite-api.theapinetwork.com/store/get"
	method := "GET"

        //Allow CORS here By * or specific origin
        w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") 

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

	//Allow CORS here By * or specific origin
        w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") 

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
  
	//Allow CORS here By * or specific origin
        w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") 
	
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
  
	//Allow CORS here By * or specific origin
        w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") 

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

	api := router.PathPrefix("/api").Subrouter()
	
	api.HandleFunc("/store", getAllStoreItems)
	api.HandleFunc("/weapons", getAllWeapons)
	api.HandleFunc("/items-list", getAllItemsList)
	api.HandleFunc("/items-list", getAllItemsList)
	api.HandleFunc("/popular-items-list", getAllPopularItems)
	
// 	n := negroni.Classic()
// 	n.UseHandler(router)
	
	http.ListenAndServe(":8000", router)

}
