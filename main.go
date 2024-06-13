package main

import (
	"fmt"

	"taskday5.2/configs"
	"taskday5.2/internal/controllers"
	"taskday5.2/internal/models"
)

func main() {
	setup := configs.ImportSetting()
	connection, err := configs.ConnectDB(setup)
	if err != nil {
		fmt.Println("Stop program, masalah pada database:", err.Error())
		return
	}

	connection.AutoMigrate(&models.User{}, &models.Todo{})

	um := models.NewUserModel(connection)
	uc := controllers.NewUserController(um)

	tm := models.NewTodoModel(connection)
	tc := controllers.NewTodoController(tm)

	var inputMenu int

	for inputMenu != 9 {
		fmt.Println("Pilih menu")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("9. Keluar")
		fmt.Print("Masukkan input: ")
		fmt.Scanln(&inputMenu)

		switch inputMenu {
		case 1:
			data, err := uc.Login()
			if err != nil {
				fmt.Println("Terjadi error pada saat login, error:", err.Error())
				continue
			}
			handleTodoMenu(data.ID, tc)
		case 2:
			_, err := uc.Register()
			if err != nil {
				fmt.Println("Terjadi error pada saat registrasi, error:", err.Error())
			} else {
				fmt.Println("Registrasi berhasil")
			}
		case 9:
			fmt.Println("Terima kasih")
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func handleTodoMenu(userID uint, tc *controllers.TodoController) {
	var inputMenu int
	for inputMenu != 9 {
		fmt.Println("Selamat datang,")
		fmt.Println("Pilih menu")
		fmt.Println("1. Tambah Kegiatan")
		fmt.Println("2. Update Kegiatan")
		fmt.Println("3. Hapus Kegiatan")
		fmt.Println("4. Tampilkan daftar kegiatan")
		fmt.Println("9. Keluar")
		fmt.Print("Masukkan input: ")
		fmt.Scanln(&inputMenu)

		switch inputMenu {
		case 1:
			_, err := tc.AddTodo(userID)
			if err != nil {
				fmt.Println("Error ketika menambahkan aktivitas:", err)
			} else {
				fmt.Println("Berhasil menambahkan aktivitas")
			}
		case 2:
			_, err := tc.UpdateTodo()
			if err != nil {
				fmt.Println("Error ketika mengupdate aktivitas:", err)
			} else {
				fmt.Println("Berhasil mengupdate aktivitas")
			}
		case 3:
			_, err := tc.DeleteTodo()
			if err != nil {
				fmt.Println("Error ketika menghapus aktivitas:", err)
			} else {
				fmt.Println("Berhasil menghapus aktivitas")
			}
		case 4:
			_, err := tc.GetTodos(userID) // Memanggil GetTodos dengan ownerID
			if err != nil {
				fmt.Println("Error ketika menampilkan daftar aktivitas:", err)
			}
		case 9:
			fmt.Println("Keluar dari menu kegiatan")
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
