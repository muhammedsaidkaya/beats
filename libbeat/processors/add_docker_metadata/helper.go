package add_docker_metadata

import (
 "bytes"
 "encoding/json"
 "fmt"
 "os/exec"
)

const ShellToUse = "bash"

func executeScript(command string) (error, string, string) {
 var stdout bytes.Buffer
 var stderr bytes.Buffer
 cmd := exec.Command(ShellToUse, "-c", command)
 cmd.Stdout = &stdout
 cmd.Stderr = &stderr
 err := cmd.Run()
 return err, stdout.String(), stderr.String()
}

func jsonToMap(jsonStr string) map[string]interface{} {
 result := make(map[string]interface{})
 jsonStr = fmt.Sprintf(`%v`, jsonStr)
 json.Unmarshal([]byte(jsonStr), &result)
 return result
}
