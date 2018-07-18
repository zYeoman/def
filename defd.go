package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/rongyi/stardict"
	"log"
	"net"
	"os"
)

var (
	h    bool
	u    string
	dict *stardict.Dictionary
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&u, "u", "", "specify the dictionary path")
	flag.Usage = usage
}

func main() {
	flag.Parse()
	n := flag.NArg()
	n = 1
	if h || n != 1 {
		flag.Usage()
		return
	} else {
		if u == "" {
			u = "dic/langdao/"
		}
	}
	a, err := Asset("langdao-ec-gb.ifo")
	b, err := Asset("langdao-ec-gb.idx")
	c, err := Asset("langdao-ec-gb.dict.dz")

	f1 := bytes.NewReader(a)
	f2 := bytes.NewReader(b)
	f3 := bytes.NewReader(c)

	dict, err = stardict.NewDictionary(f1, f2, f3)
	if err != nil {
		log.Fatal(err)
	}
	// s := flag.Arg(n - 1)
	netListen, err := net.Listen("tcp", "localhost:3154")
	CheckError(err)
	defer netListen.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		Log(conn.RemoteAddr().String(), " tcp connect success")
		handleConnection(conn)
	}
}

//处理连接
func handleConnection(conn net.Conn) {

	size := make([]byte, 1)

	n, err := conn.Read(size)
	buffer := make([]byte, int(size[0]))
	n, err = conn.Read(buffer)

	if err != nil {
		Log(conn.RemoteAddr().String(), " connection error: ", err)
		return
	}

	Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
	vs := dict.GetFormatedMeaning(string(buffer[:n]))
	for _, v := range vs {
		conn.Write([]byte(v))
		fmt.Println(v)
	}
	conn.Close()
}
func Log(v ...interface{}) {
	log.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `def version: 1.0.0
Usage: def [-h] [-u dictpath] word

Options:
`)
	flag.PrintDefaults()

}
