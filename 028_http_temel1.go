package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
<html>
<head>
<title>GOLANG </title>
</head>
<body>
Merhaba GOLANG WEB 17.10.2021
</body>
</html>`,
	)
}
func main() {
	http.HandleFunc("/golang", hello) // hello isimli fonksiyon çalıştırılır.
	http.ListenAndServe(":80", nil)
	//9000 portunda çalışır. 80 portu kullanılırsa browser da sadece ip yada localhost ile çalışır.
}
