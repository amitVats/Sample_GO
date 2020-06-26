package datasource

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"sync"
	//"sample.com/datasource"
)

type MarvelCharacterData struct{

	Name string `json:"name"`
	CharacterList []Character `json:"character"`
}


const size int= 3
var urls [size]string
var wg sync.WaitGroup


func init(){


	urls[0] = "http://localhost:10000/avenger"
	urls[1] = "http://localhost:10000/villian"
	urls[2] = "http://localhost:10000/mutant"

    UpdateMarvelData("Initial")
    
}

func fetchMdata(ch chan *MarvelCharacterData){

	for i:=0 ; i < size; i++{
		wg.Add(1)
		go consumeAPI(ch ,urls[i])
	}

    wg.Wait()

}

func consumeAPI(ch chan *MarvelCharacterData,url string) {

	defer wg.Done()
	response, err := http.Get(url)


	if err != nil {
		fmt.Printf("The http resquest has failed...")
	}else{
		data, _ := ioutil.ReadAll(response.Body)
		mdata,_ := mapToMarvelCharacterData([]byte(data))
		//dataQueque <- mdata
		ch <- mdata
	}
}


func mapToMarvelCharacterData(body []byte) (*MarvelCharacterData,error) {
 	var s = new(MarvelCharacterData)
    err := json.Unmarshal(body, &s)
    if(err != nil){
        fmt.Println("whoops:", err)
    }
    return s, err
}

func UpdateMarvelData(source string){

	var dataQueque = make(chan *MarvelCharacterData,5)
	fetchMdata(dataQueque) // fetching the latest data 
    close(dataQueque)
	for data := range dataQueque {

		for i:=0 ; i < len(data.CharacterList) ; i++{
			if(source == "Initial"){
			    character := data.CharacterList[i]
				character.AddNodeToSortedMData()

				}else{
					fmt.Println("calling update..")
					Put(data.CharacterList[i].Name, data.CharacterList[i].Max_power)
				}
			
		}

		
	}

	PrintData()
			
}
