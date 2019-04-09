package main

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdd(t *testing.T){
	total:= Add(1,2)
	assert.NotNil(t, total,"The total should not be null")
	assert.Equal(t,3,total,"Expecting 3")
}

func TestSubtract(t *testing.T){
	total:= Subtract(4,2)
	assert.NotNil(t, total,"The total should not be null")
	assert.Equal(t,2,total,"Expecting 2")
}

func Router() *mux.Router{
	router:=mux.NewRouter()
	router.HandleFunc("/",RootEndPoint).Methods("GET")
	return router
}

func TestRootEndPoint(t *testing.T) {
	request,_:= http.NewRequest("GET","/", nil)
	response:= httptest.NewRecorder()
	Router().ServeHTTP(response,request)
	assert.Equal(t,200,response.Code,"OK response expected")
	assert.Equal(t,"Hello world",response.Body.String(),"Incorrect Body found")
}
