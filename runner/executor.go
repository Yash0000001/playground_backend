package runner

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/yash0000001/playgroundbackend/utils"
)

func CompileAndRun(code, language string) (string, string, error) {
	u := uuid.New()
	var fileName string
	if language == "cpp" {
		fileName = fmt.Sprintf("temp_%d_%d.cpp", u, time.Now().UnixNano())
	} else if language == "py" {
		fileName = fmt.Sprintf("temp_%d_%d.py", u, time.Now().UnixNano())
	} else if language == "go" {
		fileName = fmt.Sprintf("temp_%d_%d.go", u, time.Now().UnixNano())
	}

	fullPath, err := utils.CreateFile(fileName, code)
	if err != nil {
		fmt.Printf("Error in Creating file: %v", err)
	}
	defer os.Remove(fullPath)

	if language == "cpp" {
		execPath := strings.TrimSuffix(fullPath, ".cpp")
		_, stderr, err := utils.ExecuteCmd("g++", fullPath, "-o", execPath)
		if err != nil {
			fmt.Println("Compilation failed:", stderr)
			return "", "", err
		}
		stdout, stderr, err := utils.ExecuteCmd("./" + execPath)
		return stdout, stderr, err
	} else if language == "go" {
		stdout, stderr, err := utils.ExecuteCmd("go", "run", fullPath)
		return stdout, stderr, err
	} else {
		stdout, stderr, err := utils.ExecuteCmd("python", fullPath)
		return stdout, stderr, err
	}
}
