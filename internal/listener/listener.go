package listener

import (
	"fmt"
	"runtime"

	reuseport "github.com/AugustineAurelius/reverse-proxy/pkg/reuse_port"
	"github.com/alitto/pond/v2"
)

type listener struct {
	pool pond.Pool
}

func New() *listener {
	return &listener{
		pool: pond.NewPool(runtime.GOMAXPROCS(0)),
	}
}

func (l *listener) Do(network string) {

	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		l.pool.Submit(func() {
			lis, err := reuseport.Listen(network, "localhost:8080")
			if err != nil {
				panic(err)
			}
			for {
				func() {
					c, err := lis.Accept()
					if err != nil {
						fmt.Println(err)
						return
					}
					defer c.Close()
					fmt.Fprintf(c, "GET / HTTP/1.0\r\n\r\n")
				}()
			}
		})
	}

	l.pool.StopAndWait()
}
