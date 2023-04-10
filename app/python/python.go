package python

import (
	"fmt"
	"os/exec"
	"runtime"
)

func PythonPIP() {
	if runtime.GOOS == "windows" {
		// Cài đặt Python
		comand := exec.Command("cmd.exe", "/c", "python", "-V")
		output, err := comand.Output()
		if err != nil {
			fmt.Println("Đang kiểm tra ENV...")
			// Cài đặt Python
			comand = exec.Command("cmd.exe", "/c", "choco", "install", "python", "-y")
			err = comand.Run()
			if err != nil {
				fmt.Println("Lỗi khi cài đặt Python: ", err)
				return
			}
			fmt.Println("Cài đặt Python thành công")
		} else {
			fmt.Println("Python đã được cài đặt với phiên bản: ", string(output))
		}
		// cài đặt PIP
		comand = exec.Command("cmd.exe", "/c", "pip", "-V")
		output, err = comand.Output()
		if err != nil {
			fmt.Println("Đang kiểm tra ENV...")
			// Cài đặt PIP
			comand = exec.Command("cmd.exe", "/c", "python", "-m", "ensurepip", "--default-pip")
			err = comand.Run()
			if err != nil {
				fmt.Println("Lỗi khi cài đặt PIP: ", err)
				return
			}
			fmt.Println("Cài đặt PIP thành công")
		} else {
			fmt.Println("PIP đã được cài đặt với phiên bản: ", string(output))
		}
	} else if runtime.GOOS == "linux" {
		// Cài đặt Python apt
		comand := exec.Command("python3", "-V")
		output, err := comand.Output()
		if err != nil {
			fmt.Println("Đang kiểm tra ENV...")
			// Cài đặt Python
			comand = exec.Command("sudo", "apt", "install", "python3", "-y")
			err = comand.Run()
			if err != nil {
				fmt.Println("Lỗi khi cài đặt Python: ", err)
				return
			}
			fmt.Println("Cài đặt Python thành công")
		} else {
			fmt.Println("Python đã được cài đặt với phiên bản: ", string(output))
		}
		// Cài đặt PIP
		comand = exec.Command("pip3", "-V")
		output, err = comand.Output()
		if err != nil {
			fmt.Println("Đang kiểm tra ENV...")
			// Cài đặt PIP
			comand = exec.Command("sudo", "apt", "install", "python3-pip", "-y")
			err = comand.Run()
			if err != nil {
				fmt.Println("Lỗi khi cài đặt PIP: ", err)
				return
			}
			fmt.Println("Cài đặt PIP thành công")
		} else {
			fmt.Println("PIP đã được cài đặt với phiên bản: ", string(output))
		}
	} else if runtime.GOOS == "darwin" {
		// Cài đặt Python
		comand := exec.Command("python3", "-V")
		output, err := comand.Output()
		if err != nil {
			fmt.Println("Đang kiểm tra ENV...")
			// Cài đặt Python
			comand = exec.Command("brew", "install", "python3")
			err = comand.Run()
			if err != nil {
				fmt.Println("Lỗi khi cài đặt Python: ", err)
				return
			}
			fmt.Println("Cài đặt Python thành công")
		} else {
			fmt.Println("Python đã được cài đặt với phiên bản: ", string(output))
		}
		// Cài đặt PIP bằng brew
		comand = exec.Command("pip3", "-V")
		output, err = comand.Output()
		if err != nil {
			fmt.Println("Đang kiểm tra ENV...")
			// Cài đặt PIP
			comand = exec.Command("brew", "install", "pip3")
			err = comand.Run()
			if err != nil {
				fmt.Println("Lỗi khi cài đặt PIP: ", err)
				return
			}
			fmt.Println("Cài đặt PIP thành công")
		} else {
			fmt.Println("PIP đã được cài đặt với phiên bản: ", string(output))
		}
		// cài đặt python-tk
		comand = exec.Command("python3", "-m", "tkinter")
		output, err = comand.Output()
		if err != nil {
			fmt.Println("Đang kiểm tra ENV...")
			// Cài đặt PIP
			comand = exec.Command("brew", "install", "python-tk")
			err = comand.Run()
			if err != nil {
				fmt.Println("Lỗi khi cài đặt python-tk: ", err)
				return
			}
			fmt.Println("Cài đặt python-tk thành công")
		} else {
			fmt.Println("python-tk đã được cài đặt với phiên bản: ", string(output))
		}
	}
}
