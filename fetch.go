// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {

			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)

		}
		fmt.Printf("%s", b)
		//var name string
		//var path = []string{
		//	"/Volumes/ITM-STICK2/Briefcase4NOV17/src/go/affordable-housing/",
		//}
		//name="count.sh"

		//pathflag := "-Xbootclasspath:" + strings.Join(path, ";")pathflag := "-Xbootclasspath:"

		//pathflag := "-Xbootclasspath:"
		//cmd := exec.Command(name, "-verbose", pathflag, "-cp slewis_lib\\*", "-jar slewis.jar")
		//err := cmd.Run()

		//if err != nil {
		//	fmt.Println("an error occurred.\n")
		//	log.Fatal(err)
		//os.StartProcess("count.sh",1,2,)

		//cmd := exec.Command("çcount.sh","5")
		//cmd := exec.Command("pwd","-verbose","\\Volumes\\ITM-STICK2\\Briefcase4NOV17\\src\\go\\affordable-housing")
		//exec.LookPath("/Volumes/ITM-STICK2/Briefcase4NOV17/src/go/affordable-housing")
		cmd := exec.Command("count.sh")

		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Waiting for command to finish...")
		err = cmd.Wait()
		log.Printf("Command finished with error: %v", err)
	}
}

//!-
