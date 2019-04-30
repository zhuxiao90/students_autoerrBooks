package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)
func getuser(w http.ResponseWriter,r *http.Request,ps httprouter.Params)  {
uid:=ps.ByName("uid")
	_, _ = fmt.Fprintf(w, "获取用户名 %s", uid)
mo:=r.Method
	_, _ = fmt.Fprintf(w, "获取请求方式 %s", mo)

}
func main()  {
router:=httprouter.New()
router.GET("/getuser/:uid",getuser)
http.Handle("/", router)
	_ = http.ListenAndServe(":8088", nil)
}
