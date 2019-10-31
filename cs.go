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

// csCmd represents the cs command
var csCmd = &cobra.Command{
	Use:   "cs",
	Short: "Run a concurrent web server.",
	Long: `cs - concurrent server
Run a server  on multiple thread 
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

		cServer(port)
	},
}

func init() {
	runCmd.AddCommand(csCmd)
}

func cServer(port int) {

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

		go handleClient(conn)
		
	}
}

func handleClient(conn net.Conn) {
	// set 2 minutes timeout
        conn.SetReadDeadline(time.Now().Add(1 * time.Minute))
        // set maximum request length to 128B to prevent flood based attacks
        request := make([]byte, 128) 
        // close connection before exit
        defer conn.Close()      
	for {
        read_len, err := conn.Read(request)

        if err != nil {
           // fmt.Println(err)
            break
        }

        if read_len == 0 {
            break // connection already closed by client
        } else {
            daytime := time.Now().String()
            conn.Write([]byte(daytime))
        }
    }
}
