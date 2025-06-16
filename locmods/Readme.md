Decide on a project name to give in this case its called "locmods"

Create a package called main in the project home directory
	$ go mod init locmods/main

Write the main program (Refer ./main.go)

Now we write a package in the subdirectory 
Create a subdirectory ./utils

We are now creating a package called utils
Change to ./utils and initialize the utils package
	$ go mod init locmods/utils
Remember the namespace "locmods" this is the theme
Write the util program refer to "goutil1.go" program
Remember only function starting with upper case is exported
Functions starting with lower case cannot be called outside
Functions starting with lower case can be called within the package

Calling util program from main.go
You need to add this import first in "main.go" in the main dir 
	"locmods/utils"
One more important step on the main package "go.mod" file
This step updates the requirement that you will need the locmods/utils package
	$ go mod edit -replace locmods/utils=./utils
	$ go get locmods/utils


In the main package you can call this anywhere this way
	<package>.Function()
	Example: utils.Utilfunct1()
