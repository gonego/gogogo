package main
import "fmt"
import "golang.org/x/crypto/bcrypt"

/*const (
    MinCost     int = 4  // the minimum allowable cost to GenerateFromPassword
    MaxCost     int = 31 // the maximum allowable cost to GenerateFromPassword
    DefaultCost int = 10 // if a cost below MinCost is passed into GenerateFromPassword
)*/

func main() {
	pswd := "gonego"
	
	//func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	//GenerateFromPassword returns the bcrypt hash of the password at the given cost
	bs, err := bcrypt.GenerateFromPassword([]byte(pswd), 8)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)

	//func Cost(hashedPassword []byte) (int, error)
	//Cost returns the hashing cost used to create the given hashed password. 
	//When, in the future, the hashing cost of a password system needs to be 
	//increased in order to adjust for greater computational power, this function 
	//allows one to establish which passwords need to be updated.
	cost,err:=bcrypt.Cost([]byte(bs))
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("Cost of hashing:",cost)


	pswd2:="gonegoing"
    
    //func CompareHashAndPassword(hashedPassword, password []byte) error
	//CompareHashAndPassword compares a bcrypt hashed password with its 
	//possible plaintext equivalent. Returns nil on success, or an error on failure.
	err=bcrypt.CompareHashAndPassword(bs, []byte(pswd2))
	if err!=nil{
		fmt.Println("Wrong password")
		return
	}
	fmt.Println("You're logged in.")
}
