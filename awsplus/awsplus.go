package main

import (
	"flag"
	"fmt"
	"github.com/humankeyboard/awsplus/functions"
)

func main() {

	flag.Usage = func() {
		fmt.Printf("Usage of awsplus:\n")
		fmt.Printf("    -help: displays help contents\n")
		fmt.Printf("    -configure: configures awsplus tool\n")
		fmt.Printf("    -create-users -file=users.txt: creates users from source file\n")
		fmt.Printf("      example users.txt is below, each line has a single user information\n")
		fmt.Printf("\n      users.txt\n")
		fmt.Printf("      ----------\n")
		fmt.Printf("      user1,password1,group1\n")
		fmt.Printf("      user2,password2,group2\n\n")
		fmt.Printf("    -create-users-undo -file=users.txt: deletes users created using source file\n")
		fmt.Printf("    -profile=default: used to select a profile created with -configure\n")
		fmt.Println()
		//flag.PrintDefaults()
	}

	var srcFile string
	var profile string
	var configure bool
	var createUsers bool
	var unCreateUsers bool

	flag.StringVar(&srcFile, "file", "users.txt", "file to read from")
	flag.StringVar(&profile, "profile", "default", "profile to use for credentials")
	flag.BoolVar(&configure, "configure", false, "configures awsplus tool")
	flag.BoolVar(&createUsers, "create-users", false, "creates users defined in a file")
	flag.BoolVar(&unCreateUsers, "create-users-undo", false, "deletes users created with create-users as defined in a file")
	flag.Parse()

	if configure {
		functions.Configure()
		return
	}

	if createUsers {
		functions.CreateUsers(profile, srcFile)
	}

	if unCreateUsers {
		functions.UnCreateUsers(profile, srcFile)
	}

}
