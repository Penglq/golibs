package golibs

import (
	"archive/zip"
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// WriteAll 快写文件
func WriteAll(filename string) error {
	err := os.WriteFile("asong.txt", []byte("Hi asong\n"), 0666)
	if err != nil {
		return err
	}
	return nil
}

// WriteLine 按行写文件
// 直接操作IO
func WriteLine(filename string) error {
	data := []string{
		"asong",
		"test",
		"123",
	}
	f, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	for _, line := range data {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	f.Close()
	return nil
}

// WriteLine2 使用缓存区写入
func WriteLine2(filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(file)

	for i := 0; i < 2; i++ {
		// 写字符串到buffer
		bytesWritten, err := bufferedWriter.WriteString(
			"asong真帅\n",
		)
		if err != nil {
			return err
		}
		log.Printf("Bytes written: %d\n", bytesWritten)
	}
	// 写内存buffer到硬盘
	err = bufferedWriter.Flush()
	if err != nil {
		return err
	}

	file.Close()
	return nil
}

// WriteAt 偏移量写入
func WriteAt(filename string) error {
	data := []byte{
		0x41, // A
		0x73, // s
		0x20, // space
		0x20, // space
		0x67, // g
	}
	f, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}

	replaceSplace := []byte{
		0x6F, // o
		0x6E, // n
	}
	_, err = f.WriteAt(replaceSplace, 2)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}

// WriteBuffer 缓存区写入
func WriteBuffer(filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(file)

	// 写字符串到buffer
	bytesWritten, err := bufferedWriter.WriteString(
		"asong真帅\n",
	)
	if err != nil {
		return err
	}
	log.Printf("Bytes written: %d\n", bytesWritten)

	// 检查缓存中的字节数
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	// 还有多少字节可用（未使用的缓存大小）
	bytesAvailable := bufferedWriter.Available()
	if err != nil {
		return err
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)
	// 写内存buffer到硬盘
	err = bufferedWriter.Flush()
	if err != nil {
		return err
	}

	file.Close()
	return nil
}

// ReadAll 读文件
// 读取全文件
func ReadAll(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	log.Printf("read %s content is %s", filename, data)
	return nil
}

func ReadAll2(filename string) error {
	file, err := os.Open("asong.txt")
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(file)
	log.Printf("read %s content is %s\n", filename, content)

	file.Close()
	return nil
}

// ReadLine 逐行读取
func ReadLine(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	bufferedReader := bufio.NewReader(file)
	for {
		// ReadLine is a low-level line-reading primitive. Most callers should use
		// ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
		lineBytes, err := bufferedReader.ReadBytes('\n')
		bufferedReader.ReadLine()
		line := strings.TrimSpace(string(lineBytes))
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			break
		}
		log.Printf("readline %s every line data is %s\n", filename, line)
	}
	file.Close()
	return nil
}

// ReadByte 按块读取文件
// use bufio.NewReader
func ReadByte(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	// 创建 Reader
	r := bufio.NewReader(file)

	// 每次读取 2 个字节
	buf := make([]byte, 2)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}
		log.Printf("writeByte %s every read 2 byte is %s\n", filename, string(buf[:n]))
	}
	file.Close()
	return nil
}

// ReadByte2 use os
func ReadByte2(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	// 每次读取 2 个字节
	buf := make([]byte, 2)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}
		log.Printf("writeByte %s every read 2 byte is %s\n", filename, string(buf[:n]))
	}
	file.Close()
	return nil
}

// ReadByte3 use os and io.ReadAtLeast
func ReadByte3(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	// 每次读取 2 个字节
	buf := make([]byte, 2)
	for {
		n, err := io.ReadAtLeast(file, buf, 0)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}
		log.Printf("writeByte %s every read 2 byte is %s\n", filename, string(buf[:n]))
	}
	file.Close()
	return nil
}

// ReadScanner 分隔符读取
func ReadScanner(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	// 可以定制Split函数做分隔函数
	// ScanWords 是scanner自带的分隔函数用来找空格分隔的文本字
	scanner.Split(bufio.ScanWords)
	for {
		success := scanner.Scan()
		if success == false {
			// 出现错误或者EOF是返回Error
			err = scanner.Err()
			if err == nil {
				log.Println("Scan completed and reached EOF")
				break
			} else {
				return err
			}
		}
		// 得到数据，Bytes() 或者 Text()
		log.Printf("readScanner get data is %s", scanner.Text())
	}
	file.Close()
	return nil
}

// WriterZip 打包/解包
func WriterZip() {
	// Create archive
	zipPath := "out.zip"
	zipFile, err := os.Create(zipPath)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new zip archive.
	w := zip.NewWriter(zipFile)
	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"asong.txt", "This archive contains some text files."},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}
	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}
}
