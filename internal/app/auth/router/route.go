package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shortenEmail/internal/app/auth"
	"shortenEmail/internal/util/context"
)

type Http struct {
	serv auth.Service
}

func Route(serv auth.Service) {
	h := &Http{
		serv: serv,
	}
	http.HandleFunc("/auth", h.ok)
	http.HandleFunc("/", h.handleCode)
	http.HandleFunc("/auth/getCode", h.handleGetCode)
}

func (h *Http) handleCode(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	email := r.URL.Query().Get("email")
	h.serv.HandleCode(code, email)
}

func (h *Http) handleGetCode(w http.ResponseWriter, r *http.Request) {
	keyVal := make(map[interface{}]interface{})
	keyVal["r"] = r
	keyVal["w"] = w
	fmt.Println("email: ", r.URL.Query().Get("email"))
	h.serv.HandleGetCode(context.ServiceContext(keyVal), r.URL.Query().Get("email"))
}

func (h *Http) ok(w http.ResponseWriter, r *http.Request) {
	keyVal := make(map[interface{}]interface{})
	keyVal["r"] = r
	keyVal["w"] = w

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	val := &auth.AuthRequest{}

	err = json.Unmarshal(bodyBytes, val)
	if err != nil {
		fmt.Println(err)
	}

	h.serv.HandleAuth(context.ServiceContext(keyVal), val)
}
