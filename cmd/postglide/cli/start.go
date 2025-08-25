package cli

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"

	"postglide.io/postglide/internal/proxy"
)

const (
	DefaultTcpPort = ":5432"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the postglide server",
	Run: func(cmd *cobra.Command, args []string) {
		listener, err := net.Listen("tcp", DefaultTcpPort)
		if err != nil {
			log.Fatal("failed to listen to the tcp server")
		}
		fmt.Println("postglide running on port:", DefaultTcpPort)

		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal("failed to listen to the tcp server")
			}

			proxy := proxy.NewProxy(conn)
			proxy.Run()
		}

	},
}
