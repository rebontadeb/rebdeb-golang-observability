package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sample_webserver/controllers"
	"sample_webserver/selflogger"
)

var ctx = context.TODO()

type employee struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
	Role string `json:"Role"`
}

func CreateEmployeeData(w http.ResponseWriter, req *http.Request) {
	//params := mux.Vars(req)
	key := "employee"
	var data employee
	jsonEncodeError := json.NewDecoder(req.Body).Decode(&data)
	if jsonEncodeError != nil {
		http.Error(w, jsonEncodeError.Error(), http.StatusBadRequest)
		selflogger.ErrorLogger.Println("Employee Creating Data having issues => " + jsonEncodeError.Error())
		return
	} else {
		newData, jsonMarshalError := json.Marshal(data)
		//fmt.Println(reflect.TypeOf(newData))
		if jsonMarshalError != nil {
			fmt.Println(jsonMarshalError)
			selflogger.ErrorLogger.Println("Error in Marshalling JSON Data => " + jsonMarshalError.Error())
		} else {
			value := string(newData)
			insertErr := controllers.MyRedisClient().Set(ctx, key, value, 0).Err()
			if insertErr != nil {
				controllers.FailureResponseWriter(w, req, []byte("Error in inserting Data in employee Database => "+insertErr.Error()))
				selflogger.ErrorLogger.Println("Error in inserting Data in employee Database => " + insertErr.Error())
			} else {
				newval, _ := controllers.MyRedisClient().Get(ctx, key).Result()
				controllers.SuccessResponseWriter(w, req, newData)
				selflogger.InfoLogger.Println("Data Inserted in employee Database Successfully => " + newval)
			}
		}
	}

}
