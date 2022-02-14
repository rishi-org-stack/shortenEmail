package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("ok"))
	// })
	// http.ListenAndServe(":8080", nil)
	// encodedValue := "4%2F0AX4XfWhFHnjaBZRnh218FoQr2QaaR85ZkeCRmw5f2jMK_zBr3PpXCvdZLawvOUjSsfYN5g"
	// decodedValue, err := url.QueryUnescape(encodedValue)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Println(decodedValue)
	res := []byte{}
	_, err := base64.URLEncoding.Decode(res, []byte("ya29.A0ARrdaM_gZhoH1PTfqE_NL7Ck_PGqqVtOCT-i2vc45lK6yXoBU7scLkTr9MceqUN35I3Wh73VDfsqRZ_pee6WkA25wRIaz214FtqT9QKGdMNt8J78e6jSA5LwxUyHcvZDOhDfxhIE3RtwJg6p4bgPJlL5x7WW"))
	if err != nil {
		fmt.Println("err ", err)

	}
	fmt.Println(string(res))
}

// package main

// import (
// 	"fmt"
// 	"os"
// 	"time"
// )

// func main() {
// 	const layout = "Jan 2, 2006 at 3:04pm (MST)"

// 	// Calling Parse() method with its parameters
// 	// “Mon, 02 Jan 2006 15:04:05 MST”
// 	tm, e := time.Parse(time.RFC1123, "Mon, 24 Jan 2022 09:22:55 IST")
// 	if e != nil {
// 		panic("Can't parse time format")
// 	}
// 	// tm = time.Now()
// 	epoch := tm.Unix()

// 	fmt.Fprintf(os.Stdout, "Epoch: %d\n", epoch)
// }
