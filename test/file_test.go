package test

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"testing"
)

// 分片大小
// const chunkSize = 100 * 1024 * 1024 // 100MB
const chunkSize = 1024 * 1024 // 1MB
// 文件分片\
func TestGenerateChunkFile(t *testing.T) {
	//fileInfo, err := os.Stat("img/R-C.jpg")
	fileInfo, err := os.Stat("img/a.jpg")
	if err != nil {
		t.Fatal(err)
	}
	//分片个数
	chunkNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	myFile, err := os.OpenFile("img/a.jpg", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b := make([]byte, chunkSize)
	for i := 0; i < int(chunkNum); i++ {
		//指定读取文件的起始位置.
		myFile.Seek(int64(i*chunkSize), 0)
		if chunkSize > fileInfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}
		myFile.Read(b)
		f, err := os.OpenFile("img/a"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		f.Write(b)
		f.Close()
	}
	myFile.Close()
}

// 分片文件的合并
func TestMergeChunkFile(t *testing.T) {
	myFile, err := os.OpenFile("test2.mp4", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	fileInfo, err := os.Stat("简单调度器.mp4")
	if err != nil {
		t.Fatal(err)
	}
	//分片的个数
	chunkNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	for i := 0; i < int(chunkNum); i++ {
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		b, err := io.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		myFile.Write(b)
		f.Close()
	}
	myFile.Close()
}

// 文件一致性校验
func TestCheckFile(t *testing.T) {
	//获取文件1的信息
	file1, err := os.OpenFile("简单调度器.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b1, err := io.ReadAll(file1)
	if err != nil {
		t.Fatal(err)
	}
	//获取文件2的信息
	file2, err := os.OpenFile("test2.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := io.ReadAll(file2)
	if err != nil {
		t.Fatal(err)
	}
	s1 := fmt.Sprintf("%x", md5.Sum(b1))
	s2 := fmt.Sprintf("%x", md5.Sum(b2))
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == s2)
}
