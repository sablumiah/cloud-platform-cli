package environment

import (
	"os"
	"testing"
)

func TestCreatesEcrTfFile(t *testing.T) {
	filename := "resources/ecr.tf"
	os.Mkdir("resources", 0755)

	err := createEcrTfFile()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	moduleName := "github.com/ministryofjustice/cloud-platform-terraform-ecr-credentials"
	fileContainsString(t, filename, moduleName)

	os.Remove(filename)
	os.Remove("resources")
}
