package handler

import (
	"context"
	"encoding/json"
	GETAREA "study/students_autoerrBooks/GetArea/proto/example"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"
	"net/http"
	"time"
)

/*func ExampleCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	exampleClient := example.NewExampleService("go.micro.srv.template", client.DefaultClient)
	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}*/
func GetArea(w http.ResponseWriter, r *http.Request,_ httprouter.Params) {
	// decode the incoming request as json
	//前端传递的数据，这里前端没有数据传递
	/*var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}*/

	// call the backend service
	// 创建新的grpc举兵
	service := grpc.NewService()
	//服务初始化
	service.Init()
	//创建获取地区的服务并返回句柄
	exampleClient := GETAREA.NewExampleService("go.micro.srv.GetArea", service.Client())
	rsp, err := exampleClient.Call(context.TODO(), &Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}