package haproxycmd

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func Command(sock string, args []string, out io.Writer) (oout int64) {
	conn, err := net.Dial("unix", sock)

	if err != nil {
		panic(err)
	}
	//defer conn.Close()
	cmd := strings.Join(args, " ")
	fmt.Fprintln(conn, cmd)

	//var out bytes.Buffer

	oout, err = io.Copy(out, conn)
	if err != nil {
		panic(err)
	}

	return oout

}
