package main

import (
	"encoding/json"
	"github.com/antchfx/htmlquery"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type CrawlerOutput struct {
	Status        string   `json:"status,omitempty"`
	Message       interface{}  `json:"message,omitempty"`
}




type CrawlerInput struct {
	URL       string   `json:"url,omitempty"`
	MAP       map[string]string `json:"map"`
}
type ParserList []string


func ItsON(w http.ResponseWriter, r *http.Request) {
	var Response CrawlerOutput
	Response.Status = "It's On"

	_ = json.NewEncoder(w).Encode(Response)
}
func ParsePage(w http.ResponseWriter, r *http.Request) {
	var Response CrawlerOutput
	Response.Status = "Success"

	var responseList map[string]ParserList
	responseList = make(map[string]ParserList)

	var input CrawlerInput
	_ = json.NewDecoder(r.Body).Decode(&input)

	doc, err := htmlquery.LoadURL(input.URL)
	if err != nil {
		panic(err)
	}

	for reqIndex, reqElement := range input.MAP {
		list :=  htmlquery.Find(doc, reqElement)
		var pList ParserList
		for _, element := range list {
			pList = append(pList,  htmlquery.OutputHTML(element, false))

			// index is the index where we are
			// element is the element from someSlice for where we are

			// fmt.Printf("total count is %+v\n", htmlquery.OutputHTML(element, false))
			// fmt.Printf("total count is %+v\n", htmlquery.SelectAttr(element, "href"))
			// fmt.Printf("total count is %d\n", index)
		}

		responseList[reqIndex] = pList
	}

	Response.Message = responseList

	_ = json.NewEncoder(w).Encode(Response)
}

func parse(){

	//fmt.Printf("total count is %+v\n", nodes)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", ParsePage).Methods("POST")
	router.HandleFunc("/", ItsON).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}