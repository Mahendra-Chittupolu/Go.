package greetings

import "fmt"
import "math/rand"
import "errors"
func Greet (name string) (string, error){
	if name == "" {
		return name, errors.New("Give me a name")
	}
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil

}
func Greets(names []string) (map[string]string, error){
	messages := make(map[string]string)
	for _,name := range names{
		message  ,err :=  Greet(name)
		if err != nil{
			return nil,err
		}
		messages[name] = message
	}

	return messages , nil

}
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Aimimor, %v",
        "Eliot %v",
	}
	return formats[rand.Intn(len(formats))]
}
