package runner

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/yash0000001/playgroundbackend/utils"
)

func CompileAndRun(code, language string) (string, string, int64, error) {
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

	initialTime := time.Now()
	if language == "cpp" {
		execPath := strings.TrimSuffix(fullPath, ".cpp")
		_, stderr, err := utils.ExecuteCmd("g++", fullPath, "-o", execPath)
		defer os.Remove(execPath + ".exe")
		if err != nil {
			fmt.Println("Compilation failed:", stderr)
			return "", "", 0, err
		}
		stdout, stderr, err := utils.ExecuteCmd("./" + execPath)
		finalTime := time.Now()
		diff := finalTime.Sub(initialTime).Milliseconds()
		return stdout, stderr, diff, err
	} else if language == "go" {
		stdout, stderr, err := utils.ExecuteCmd("go", "run", fullPath)
		finalTime := time.Now()
		diff := finalTime.Sub(initialTime).Milliseconds()
		return stdout, stderr, diff, err
	} else {
		stdout, stderr, err := utils.ExecuteCmd("python", fullPath)
		finalTime := time.Now()
		diff := finalTime.Sub(initialTime).Milliseconds()
		return stdout, stderr, diff, err
	}
}
