package environment

import (
	"os"
	"testing"
)

func TestCreatesS3TfFile(t *testing.T) {
	filename := "resources/s3.tf"
	os.Mkdir("resources", 0755)

	err := createS3TfFile()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	moduleName := "github.com/ministryofjustice/cloud-platform-terraform-s3-bucket"
	fileContainsString(t, filename, moduleName)

	os.Remove(filename)
	os.Remove("resources")
}
