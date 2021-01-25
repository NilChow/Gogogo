package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Zip Test Start...")

	GetBasicInfo()
}

func GetBasicInfo() {
	reader, err := zip.OpenReader("20200808 103136.zip")
	if err != nil {
		fmt.Printf("[OpenReader Error]: [%s]\n", err)
		return
	}

	defer reader.Close()

	for _, file := range reader.File {

		if file.FileInfo().Name() != "UserLastConfig.ini" {
			continue
		}

		//f, err := os.OpenFile(file.Name, os.O_RDONLY, os.ModePerm)
		//if err != nil {
		//	fmt.Printf("[OS OpenFile Error]: [%s]\n", err)
		//	return
		//}
		//var p []byte
		//lens, err := f.Read(p)
		//if err != nil {
		//	fmt.Printf("[Read Error]: [%s]\n", err)
		//	return
		//} else {
		//	fmt.Printf("[Read Success, total bytes: %d]\n", lens)
		//	fmt.Printf("[Data]: [%s]\n", p)
		//}



		rc, err := file.Open()
		if err != nil {
			fmt.Printf("[File Open Error]: [%s]\n", err)
			return
		}

		//w := &bufio.Writer{}
		//buf := &bytes.Buffer{}

		var symbol, period string
		var indexSize, objSize int
		s := bufio.NewScanner(rc)

		for s.Scan() {
			ss := strings.Split(s.Text(), ",")
			if len(ss) != 4 {
				fmt.Printf("[Info Error]: [Info: %s]\n", s.Text())
				continue
			}
			symbol = ss[0]
			period = ss[1]
			indexSize, err = strconv.Atoi(ss[2])
			objSize, err = strconv.Atoi(ss[3])
			fmt.Printf("[%s %s %d %d]\n", symbol, period, indexSize, objSize)
		}

		rc.Close()
		//lens2, err := io.Copy(buf, rc)
		//if err != nil {
		//	fmt.Printf("[Copy Error]: [%s]\n", err)
		//	return
		//} else {
		//	//fmt.Printf("[Read Success, total bytes: %d]\n", lens)
		//	//fmt.Printf("[Data]: [%s]\n", p)
		//	fmt.Println(lens2)
		//	// fmt.Printf("[Data]:\n%s", buf.Bytes())
		//}
		//
		//rc.Close()
		//
		//info := string(buf.Bytes())
		//infos := strings.Split(info, "\n")
		//var symbol, period string
		//var indexSize, objSize int
		//for _, v := range infos {
		//	ss := strings.Split(v, ",")
		//	if len(ss) != 4 {
		//		fmt.Printf("[Info Error]: [Info: %s]", v)
		//		continue
		//	}
		//	symbol = ss[0]
		//	period = ss[1]
		//	indexSize, err = strconv.Atoi(ss[2])
		//	objSize, err = strconv.Atoi(ss[3])
		//	fmt.Printf("[%s %s %d %d]\n", symbol, period, indexSize, objSize)
		//}
	}
}

//func Compress(files []*os.File, dest string) error {
//	df, _ := os.Create(dest)
//	defer df.Close()
//
//	w := zip.NewWriter(df)
//	defer w.Close()
//
//	for _, f := range files {
//		err := compress(f, "", w)
//		if err != nil {
//			panic("[compress error]")
//		}
//	}
//	return nil
//}
//
//func compress(file *os.File, prefix string, w *zip.Writer) error {
//	//info, err := file.Stat()
//	//if err != nil {
//	//	fmt.Printf("[compress error]: [%s]\n", err)
//	//	return err
//	//}
//	//
//	//if info.IsDir() {
//	//	prefix = prefix + "/" + info.Name()
//	//	fileInfos, err := file.Readdir()
//	//}
//}

//files 文件数组，可以是不同dir下的文件或者文件夹
//dest 压缩文件存放地址
func Compress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//解压
func DeCompress(zipFile, dest string) (err error) {
	//目标文件夹不存在则创建
	if _, err = os.Stat(dest); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(dest, 0755)
		}
	}

	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}

	defer reader.Close()

	for _, file := range reader.File {
		//    log.Println(file.Name)

		if file.FileInfo().IsDir() {

			err := os.MkdirAll(dest+"/"+file.Name, 0755)
			if err != nil {
				log.Println(err)
			}
			continue
		} else {

			err = os.MkdirAll(getDir(dest+"/"+file.Name), 0755)
			if err != nil {
				return err
			}
		}

		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		filename := dest + "/" + file.Name
		//err = os.MkdirAll(getDir(filename), 0755)
		//if err != nil {
		//    return err
		//}

		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()

		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		//w.Close()
		//rc.Close()
	}
	return
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

func CompressZip(src string, dest string) (err error) {

	f, err := ioutil.ReadDir(src)
	if err != nil {
		log.Println(err)
	}

	fzip, _ := os.Create(dest)
	w := zip.NewWriter(fzip)

	defer fzip.Close()
	defer w.Close()
	for _, file := range f {
		fw, _ := w.Create(file.Name())
		filecontent, err := ioutil.ReadFile(src + file.Name())
		if err != nil {
			log.Println(err)
		}
		_, err = fw.Write(filecontent)

		if err != nil {
			log.Println(err)
		}
		//    log.Println(n)
	}
	return
}