You may have to ruhn the following command:
  go get -u github.com/manifoldco/promptui

Then to run the program you can do it o ne of 2 ways:
1. ~ go run generate.go --> follow the CLI
2. Or you can go into generate.go, make the useCLI variable equal to false:
  In the file define the struct / interface name ("function"), receiver name ("fn"), and package name ("funcmodels")
  Then fill in the field name and types right below
  Then run: ~ go run generate.go

Capitalization should not matter for the entry in generate.go
