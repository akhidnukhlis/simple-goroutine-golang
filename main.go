package main

import (
	"errors"
	"fmt"
	"time"
)

// DataSaver adalah interface untuk menyimpan data
type DataSaver interface {
	Save(data string) (string, error)
}

// DatabaseSaver adalah implementasi DataSaver yang menyimpan data ke database
type DatabaseSaver struct {
	Host     string
	Username string
	Password string
	Port     int
	Options  map[string]interface{}
}

func (d *DatabaseSaver) Save(data string) (string, error) {
	// Simulasikan proses penyimpanan data ke database
	time.Sleep(2 * time.Second)

	// Simulasikan terjadi error
	// Misalnya, jika data kosong
	if data == "" {
		return "", errors.New("Data kosong")
	}

	// Simulasikan hasil penyimpanan
	result := fmt.Sprintf("Data '%s' berhasil disimpan di database", data)

	return result, nil
}

// FileSaver adalah implementasi DataSaver yang menyimpan data ke file
type FileSaver struct {
	Path    string
	Format  string
	Options map[string]interface{}
}

func (f *FileSaver) Save(data string) (string, error) {
	// Simulasikan proses penyimpanan data ke file
	time.Sleep(2 * time.Second)

	// Simulasikan terjadi error
	// Misalnya, jika data kosong
	if data == "" {
		return "", errors.New("Data kosong")
	}

	// Simulasikan hasil penyimpanan
	result := fmt.Sprintf("Data '%s' berhasil disimpan di file", data)

	return result, nil
}

func saveData(data string, saver DataSaver, resultChan chan<- string, errorChan chan<- error) {
	result, err := saver.Save(data)
	if err != nil {
		errorChan <- err
		return
	}

	resultChan <- result
}

func main() {
	data := "Hi, World!"

	// Buat channel untuk menerima hasil dan error
	resultChan := make(chan string)
	errorChan := make(chan error)

	// Jalankan goroutine untuk menyimpan data ke database
	go saveData(data, &DatabaseSaver{}, resultChan, errorChan)

	// Menangani hasil dan error
	select {
	case result := <-resultChan:
		// Cetak hasil jika sukses
		fmt.Println(result)
	case err := <-errorChan:
		// Tangani error jika terjadi
		fmt.Println("Error:", err)
	}

	// Jalankan goroutine untuk menyimpan data ke file
	go saveData(data, &FileSaver{}, resultChan, errorChan)

	// Menangani hasil dan error
	select {
	case result := <-resultChan:
		// Cetak hasil jika sukses
		fmt.Println(result)
	case err := <-errorChan:
		// Tangani error jika terjadi
		fmt.Println("Error:", err)
	}
}
