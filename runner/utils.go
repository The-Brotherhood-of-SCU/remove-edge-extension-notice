package runner

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

func copyfile(src, dst string) (err error) {
	source, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer source.Close()
	destination, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func getEdgeConfigPath() string {
	u, err := user.Current()
	if err != nil {
		panic("error when get system user info")
	}
	return filepath.Join(u.HomeDir, "AppData/Local/Microsoft/Edge/User Data/Default")
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println("Error:", err)
	return false
}

func checkConfig() (success bool, cfg map[string]interface{}) {
	path := filepath.Join(getEdgeConfigPath(), "Preferences")
	if !fileExists(path) {
		fmt.Println("file not exists")
		return false, nil
	}
	c, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("error when read config file")
		return false, nil
	}
	var config map[string]interface{}
	err = json.Unmarshal(c, &config)
	if err != nil {
		fmt.Println("error when parse config file")
		return false, nil
	}
	_, hasValue := config["extensions"].(map[string]interface{})["ui"].(map[string]interface{})["dev_mode_warning_snooze_end_time"].(string)
	if !hasValue {
		fmt.Println("config value not exists")
		return false, nil
	}
	return true, config
}

func closeEdge() bool {
	cmd := exec.Command("taskkill", "/f", "/t", "/im", "msedge.exe")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
