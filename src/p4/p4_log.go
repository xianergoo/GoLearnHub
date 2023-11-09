package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	//你可以使用 log.Fatal() 函数记录错误并结束程序
	//log.Fatal("hey, I'm a error log(fatal)!")
	// log.Panic("hey, this is a error log(panic)!") //get erro stack info

	/* log.SetPrefix("main(): ")
	fmt.Println("Can you see me ")
	log.Print("hey, I'm a log")
	log.Fatal("hey, I'm a error log(fatal)!") */

	/* file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	r := log.Writer()
	log.SetOutput(file)
	log.SetPrefix("file log: ")
	log.Print("Hey, I'm a log!")

	if r != nil {
		log.SetOutput(r)
		log.SetPrefix("origin log: ")
		log.Print("Hey, I'm a log!")
	} */

	// zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Debug().
		Int("EmployeeeID", 1001).
		Msg("Getting employee information")

	log.Debug().
		Str("Name", "John").
		Send()
	log.Print("hey!, I'm a log message!")
}
