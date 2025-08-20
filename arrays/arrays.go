package main

import "fmt"

func main() {
	//var DeploymentOptions = [4]string{"R-pi", "AWS", "GCP", "Azure"}
	DeploymentOptions := [...]string{"R-pi", "AWS", "GCP", "Azure"}

	//loop
	/*for i := 0; i < len(DeploymentOptions); i++ {
		option := DeploymentOptions[i]
		fmt.Println(i, option)
	}*/
	for index, option := range DeploymentOptions {
		fmt.Println(index, option)
	}

}
