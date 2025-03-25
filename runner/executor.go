package runner

import (
	"fmt"
	"os"
	"path/filepath"
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
	} else if language == "js" {
		fileName = fmt.Sprintf("temp_%d_%d.js", u, time.Now().UnixNano())
	} else if language == "c" {
		fileName = fmt.Sprintf("temp_%d_%d.c", u, time.Now().UnixNano())
	} else if language == "java" {
		fileName = fmt.Sprintf("temp_%d_%d.java", u, time.Now().UnixNano())
	} else if language == "rust" {
		fileName = fmt.Sprintf("temp_%d_%d.rs", u, time.Now().UnixNano())
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
	} else if language == "java" {
		// Compile Java code
		className := strings.TrimSuffix(filepath.Base(fullPath), ".java")
		_, stderr, err := utils.ExecuteCmd("javac", fullPath)
		if err != nil {
			fmt.Println("Compilation failed:", stderr)
			return "", "", 0, err
		}
		defer os.Remove(className + ".class")
		// Run Java code
		stdout, stderr, err := utils.ExecuteCmd("java", "-cp", filepath.Dir(fullPath), className)
		finalTime := time.Now()
		diff := finalTime.Sub(initialTime).Milliseconds()
		return stdout, stderr, diff, err
	} else if language == "rust" {
		// Compile Rust code
		execPath := strings.TrimSuffix(fullPath, ".rs")
		_, stderr, err := utils.ExecuteCmd("rustc", fullPath, "-o", execPath)
		defer os.Remove(execPath)
		if err != nil {
			fmt.Println("Compilation failed:", stderr)
			return "", "", 0, err
		}
		// Run Rust code
		stdout, stderr, err := utils.ExecuteCmd("./" + execPath)
		finalTime := time.Now()
		diff := finalTime.Sub(initialTime).Milliseconds()
		return stdout, stderr, diff, err
	} else if language == "js" {
		// Run JavaScript code using Node.js
		stdout, stderr, err := utils.ExecuteCmd("node", fullPath)
		finalTime := time.Now()
		diff := finalTime.Sub(initialTime).Milliseconds()
		return stdout, stderr, diff, err
	} else if language == "c" {
		// Compile C code
		execPath := strings.TrimSuffix(fullPath, ".c")
		_, stderr, err := utils.ExecuteCmd("gcc", fullPath, "-o", execPath)
		defer os.Remove(execPath)
		if err != nil {
			fmt.Println("Compilation failed:", stderr)
			return "", "", 0, err
		}
		// Run C code
		stdout, stderr, err := utils.ExecuteCmd("./" + execPath)
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
