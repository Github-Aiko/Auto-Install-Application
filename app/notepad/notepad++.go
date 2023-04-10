package notepad

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func Notepadpp() {
	// Tải xuống gói cài đặt của Notepad++
	response, err := http.Get("https://github.com/notepad-plus-plus/notepad-plus-plus/releases/download/v8.5.2/npp.8.5.2.Installer.x64.exe")
	if err != nil {
		fmt.Println("Lỗi khi tải xuống gói cài đặt của Notepad++: ", err)
		return
	}
	defer response.Body.Close()

	// Lưu gói cài đặt vào một file tạm
	file, err := os.Create("npp.8.5.2.Installer.x64.exe")
	if err != nil {
		fmt.Println("Lỗi khi tạo file tạm: ", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Lỗi khi lưu gói cài đặt vào file tạm: ", err)
		return
	}

	// Lệnh cài đặt Notepad++
	command := exec.Command("cmd.exe", "/c", "npp.8.5.2.Installer.x64.exe", "/S")

	// Thực thi lệnh cài đặt
	output, err := command.Output()
	if err != nil {
		fmt.Println("Lỗi khi cài đặt Notepad++: ", err)
		fmt.Println("Lỗi :", string(output))
		return
	}

	fmt.Println("Cài đặt Notepad++ thành công")
}
