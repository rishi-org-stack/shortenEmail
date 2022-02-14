package mail

import "net/http"

func Route() {
	http.HandleFunc("/mail", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("mail"))
	})
}
