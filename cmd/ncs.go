package cmd

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// ncsCmd represents the ncs command
var ncsCmd = &cobra.Command{
	Use:   "ncs",
	Short: "Run a non-concurrent web server.",
	Long: `ncs - Non-concurrent server

Run a server only on a single thread without creating multiple threads
for multiple requests. Server returns current time`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := runCmd.PersistentFlags().GetInt("port")

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)

		go func() {
			<-sigChan
			fmt.Println("You pressed ctrl + C. Server terminated!")
			os.Exit(0)
		}()

		ncServer(port)
	},
}

func init() {
	runCmd.AddCommand(ncsCmd)
}

func ncServer(port int) {

	service := ":" + strconv.Itoa(port)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// Listens on given port number
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()

		fmt.Println("Client", conn.RemoteAddr(), "connected")

		if err != nil {
			continue
		}

		daytime := time.Now().String()
		myTime := daytime + "\n"

		conn.Write([]byte(myTime))
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stdout, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
