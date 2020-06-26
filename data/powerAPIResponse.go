package data

import(
	 "encoding/json"
	 "io"	 
)

type PowerResponse struct{
	Name string `json:"name"`
	Max_power int `json:"max_power"`
}

func (pr * PowerResponse) ToJSON(w io.Writer) error{
	e := json.NewEncoder(w)
	return e.Encode(pr)
}


func FmtPowerRes(name string, max_power int) *PowerResponse {
	var samplePowerResponse *PowerResponse = &PowerResponse { Name : name, Max_power : max_power }
	return samplePowerResponse
}