package main

import (
	"fmt"
	"net/http"
)

func main() {

	myServer()

	fmt.Printf("ssss")
	serveErr := http.ListenAndServe(":8888", nil)
	if serveErr != nil {
		return
	}
}

func myServer() {

	http.HandleFunc("/", Middlewares(myServer1))

	//fs := http.FileServer(http.Dir("healthz/"))
	http.HandleFunc("/healthz/", Middlewares(myServer2))
	//dealHttp(x, r)

}

func Middlewares(mySer func(x http.ResponseWriter, r *http.Request)) func(x http.ResponseWriter, r *http.Request) {

	return func(x http.ResponseWriter, r *http.Request) {

		dealHttpStart(x, r)
		mySer(x, r)
		dealHttperf(x, r)

	}

}

func dealHttpStart(x http.ResponseWriter, r *http.Request) {
	//fmt.Printf(r.Host)
	for k, v := range r.Header {
		s := fmt.Sprintf(
			"%s", v)
		//fmt.Printf(s)

		x.Header().Set(k, s)
		//fmt.Printf("%s--%s\n", k, v)
	}
}
func dealHttperf(x http.ResponseWriter, r *http.Request) {
	_ = x
	fmt.Printf("ip:%s statuscode:%b\n", r.Host, http.StatusOK)

}
func myServer1(x http.ResponseWriter, r *http.Request) {
	_ = r
	_, err := fmt.Fprintf(x, "my_server")
	if err != nil {
		return
	}
}
func myServer2(x http.ResponseWriter, r *http.Request) {
	_ = r
	//for k, _ := range r.Header {
	//
	//	x.Header().Set(k, "v")
	//}
	_, err := fmt.Fprintf(x, "200")
	//fmt.Printf("SSSS S %s", r.Header)
	if err != nil {

		return
	}
}
