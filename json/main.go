package json

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type User struct {
	Id      string
	Balance uint64
}

func main() {
	u := User{Id: "US008", Balance: 88}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	res, _ := http.Post("https://httpbin.org/post", "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, res.Body)
}
