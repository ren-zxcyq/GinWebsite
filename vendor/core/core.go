package core

import (
	"crypto/rand"
	"fmt"
)

//	Issue Cookie according to Gin Example using https://stackoverflow.com/a/38418781
func GenCookie(potentialCreds string) string {

	//	@TODO	Perform some operation to verify:	potentialCreds

	//	Generate Random 32 digit Value using crypto/rand
	var b []byte = make([]byte, 32)

	var newcval string //	To hold new Cookie Value

	rand.Read(b)
	if _, err := rand.Read(b); err != nil {
		//	Generate Error page

		//	panic(err)	@TODO	-	check documentation for		panic()
		fmt.Printf("[+] Cookie Generation Error\t-\t%s", err)

	}
	newcval = fmt.Sprintf("%x", b)

	return newcval
}
