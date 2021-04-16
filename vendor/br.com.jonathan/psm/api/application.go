package application

import (
    "strconv"
    "net/http"
    "github.com/gorilla/mux"
)

func Init(port int) {
    endpoints := getEndpoints()
    router := mux.NewRouter()
    for index := range endpoints {
        endpoint := endpoints[index]
        router.HandleFunc(endpoint.path, endpoint.handler).Methods(endpoint.method)
    }
    http.Handle("/", router)
    http.ListenAndServe(":" + strconv.Itoa(port), nil)
}
