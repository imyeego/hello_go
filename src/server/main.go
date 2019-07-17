package main

import ("fmt"
		"net/http")


func main()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "<h1>Hello World! %s </h1>", request.FormValue("name"))
	})

	_ = http.ListenAndServe(":8888", nil)

}
