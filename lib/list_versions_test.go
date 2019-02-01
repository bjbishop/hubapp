package lib_test

import (
	"log"
	"testing"

	"github.com/warrensbox/github-appinstaller/lib"
)

const (
	gruntURL = "https://api.github.com/repos/gruntwork-io/terragrunt/releases"
)

// TestGetAppList : Get list from github
func TestGetAppList(t *testing.T) {

	// list, _ := lib.GetAppList(gruntURL)

	// val := "0.14.11"
	// var exists bool

	// switch reflect.TypeOf(list).Kind() {
	// case reflect.Slice:
	// 	s := reflect.ValueOf(list)

	// 	for i := 0; i < s.Len(); i++ {
	// 		if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
	// 			exists = true
	// 		}
	// 	}
	// }

	// if !exists {
	// 	log.Fatalf("Not able to find version: %s\n", val)
	// } else {
	// 	t.Log("Write versions exist (expected)")
	// }

}

//TestRemoveDuplicateVersions :  test to removed duplicate
func TestRemoveDuplicateVersions(t *testing.T) {

	test_array := []string{"0.0.1", "0.0.2", "0.0.3", "0.0.1"}

	list := lib.RemoveDuplicateVersions(test_array)

	if len(list) == len(test_array) {
		log.Fatalf("Not able to remove duplicate: %s\n", test_array)
	} else {
		t.Log("Write versions exist (expected)")
	}
}
