package main

import (
	"fmt"
	"os/exec"
	"runtime"

	"Github-Aiko/Auto-Install-Application/app/python"
)

func checkos() {
	os := runtime.GOOS
	if os == "darwin" {
		os = "MacOS"
	}
	fmt.Println("Hệ điều hành của bạn là:", os)
}

func main() {
	checkos()
	if runtime.GOOS == "windows" {
		comand := exec.Command("cmd.exe", "/c", "choco", "-v")

		output, err := comand.Output()
		if err != nil {
			fmt.Println("Đang kiểm tra ENV...")
			// Cài đặt Chocolatey
			comand = exec.Command("cmd.exe", "/c", "powershell", "-NoProfile", "-InputFormat", "None", "-ExecutionPolicy", "Bypass", "-Command", "[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))")
			err = comand.Run()
			if err != nil {
				fmt.Println("Lỗi khi cài đặt Chocolatey: ", err)
				return
			}
			fmt.Println("Cài đặt Chocolatey thành công")
		} else {
			fmt.Println("Chocolatey đã được cài đặt với phiên bản: ", string(output))
		}
	} else if runtime.GOOS == "linux" {
		fmt.Printf("Comming soon...")
	} else if runtime.GOOS == "darwin" {
		// Cài đặt Homebrew
		comand := exec.Command("cmd.exe", "/c", "brew", "-v")

		output, err := comand.Output()
		if err != nil {
			fmt.Println("Đang kiểm tra ENV...")
			comand = exec.Command("cmd.exe", "/c", "ruby", "-e", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)")
			err = comand.Run()
			if err != nil {
				fmt.Println("Lỗi khi cài đặt Homebrew: ", err)
				return
			}
			fmt.Println("Cài đặt Homebrew thành công")
		} else {
			fmt.Println("Homebrew đã được cài đặt với phiên bản: ", string(output))
		}
	}
	python.PythonPIP()
}
