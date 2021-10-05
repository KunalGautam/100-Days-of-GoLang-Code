package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func init() {

	envFlag := flag.String("env", "prod", "Load Production or Development environmental file. Use -env=prod for Production or -env=dev for Development")
	flag.Parse()
	var err error
	var filename string
	switch *envFlag {
	case "dev":
		filename = ".env.dev"
		err = godotenv.Load(filename)
	case "prod":
		filename = ".env"
		err = godotenv.Load(filename)
	default:
		fmt.Println("Invalid Option Provided")
		os.Exit(0)
	}
	// If not having enough permission to read the file
	if os.IsPermission(err) {
		fmt.Println(err)
		os.Exit(0)
	}
	// If file not exists.
	if os.IsNotExist(err) {
		fmt.Println(err)
		createFile(filename)
	}
}

func main() {
	fmt.Println(os.Getenv("DB_HOST"))
	fmt.Println(os.Getenv("DB_PORT"))
	fmt.Println(os.Getenv("DB_NAME"))
	fmt.Println(os.Getenv("DB_USER"))
	fmt.Println(os.Getenv("DB_PASS"))
}

func createFile(filename string) {
	var dbloc, dbname, dbuser, dbpass string
	var dbport int

	fmt.Println("Enter Database Location")
	fmt.Scanln(&dbloc)
	fmt.Println("Enter Database Port")
	fmt.Scanln(&dbport)
	fmt.Println("Enter Database Name")
	fmt.Scanln(&dbname)
	fmt.Println("Enter Database User")
	fmt.Scanln(&dbuser)
	fmt.Println("Enter Database Password")
	fmt.Scanln(&dbpass)

	data := fmt.Sprintf("DB_HOST:%s\nDB_PORT:%d\nDB_NAME:%s\nDB_USER:%s\nDB_PASS:%s", dbloc, dbport, dbname, dbuser, dbpass)

	f, err := os.Create(filename)
	defer f.Close()
	err = os.WriteFile(filename, []byte(data), 644)
	if err != nil {
		panic(err)
	}

	fmt.Println(filename, "Created")
	godotenv.Load(filename)

}
