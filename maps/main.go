package maps

import "fmt"
type StringMap map[string]string
func (m StringMap) outputMap(){
  fmt.Println(m)
}

func PlayWithMaps() {
	websiteWithCompanyNames := map[string]string{
		"Google": "https://google.com",
		"Meta":   "https://facebook.com",
	}
	fmt.Println(websiteWithCompanyNames)

	// Add or modify exisiting keys in map
	websiteWithCompanyNames["LinkedIn"] = "https://linkedin.com"
	websiteWithCompanyNames["Google"] = "https://newgoogle.com"
	fmt.Println(websiteWithCompanyNames)

	// Delete an existing key in the map using delte built function
	delete(websiteWithCompanyNames, "Google")
	fmt.Println(websiteWithCompanyNames)

}
func MakeFnOptimizationWithArraySlicesAndMaps() {
	userNames := make([]string, 2, 3) // The make function is used to allocate memory space for how length and capacity specified as the second and third argument, so that you don't need to append you can overwrite those [LENGTH] elements using their indexes and capacity for to efficient memory management when adding elements to that array

	userNames[0] = "Abhishek"
	userNames[1] = "Niranjan"

	userNames = append(userNames, "Rohan")
	userNames = append(userNames, "Swapnil")

	//fmt.Println("User names are: ", userNames)
  for index, value := range userNames {
    fmt.Println("Index: ", index)
    fmt.Println("Value is", value)
  }
  
  //make function for maps
  websiteWithCompanyNames := make(StringMap, 3) // The second argument for make function is the capacity of the map
  websiteWithCompanyNames["Google"] = "https://google.com"
  websiteWithCompanyNames["Meta"] = "https://facebook.com"
  websiteWithCompanyNames["LinkedIn"] = "https://linkedin.com"
  //websiteWithCompanyNames.outputMap();

  for key, value := range websiteWithCompanyNames {
    fmt.Println("Key: ", key, "Value: ", value)
  }
}
