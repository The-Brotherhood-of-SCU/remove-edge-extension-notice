package runner

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Runner struct{}

func (r *Runner) Backup() bool {
	path := filepath.Join(getEdgeConfigPath(), "Preferences")
	err := copyfile(path, "./Preferences")
	if err != nil {
		fmt.Println("backup fail")
		return false
	}
	fmt.Println("backup success")
	return true
}

func (r *Runner) Recovery() bool {
	path := filepath.Join(getEdgeConfigPath(), "Preferences")
	err := copyfile("./Preferences", path)
	if err != nil {
		fmt.Println("recovery fail")
		return false
	}
	fmt.Println("recovery success")
	return true
}

func (r *Runner) WriteConfig() {
	check, config := checkConfig()
	if !check {
		return
	}
	if !closeEdge() {
		fmt.Println("can't close edge")
		return
	}
	config["extensions"].(map[string]interface{})["ui"].(map[string]interface{})["dev_mode_warning_snooze_end_time"] = "9223372036854775807"
	bts, err := json.Marshal(config)
	if err != nil {
		fmt.Println("write fail")
		return
	}
	err = os.WriteFile(filepath.Join(getEdgeConfigPath(), "Preferences"), bts, os.ModePerm)
	if err != nil {
		fmt.Println("write fail: ", err.Error())
		return
	}
	fmt.Println("write success")
}
