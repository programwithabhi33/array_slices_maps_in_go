package maps

import "fmt"

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
