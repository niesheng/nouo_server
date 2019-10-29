package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var FileType_ sync.Map

func init_file_type() {
	FileType_.Store("ffd8ffe", "jpg")               //JPEG (jpg)
	FileType_.Store("89504e470d0a1a0a0000", "png")  //PNG (png)
	FileType_.Store("474946383961", "gif")          //GIF (gif)
	FileType_.Store("49492a00227105008037", "tif")  //TIFF (tif)
	FileType_.Store("424d228c010000000000", "bmp")  //16色位图(bmp)
	FileType_.Store("424d8240090000000000", "bmp")  //24位位图(bmp)
	FileType_.Store("424d8e1b030000000000", "bmp")  //256色位图(bmp)
	FileType_.Store("41433130313500000000", "dwg")  //CAD (dwg)
	FileType_.Store("3c21444f435459504520", "html") //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	FileType_.Store("3c68746d6c3e0", "html")        //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	FileType_.Store("3c21646f637479706520", "htm")  //HTM (htm)
	FileType_.Store("48544d4c207b0d0a0942", "css")  //css
	FileType_.Store("696b2e71623d696b2e71", "js")   //js
	FileType_.Store("7b5c727466315c616e73", "rtf")  //Rich Text Format (rtf)
	FileType_.Store("38425053000100000000", "psd")  //Photoshop (psd)
	FileType_.Store("46726f6d3a203d3f6762", "eml")  //Email [Outlook Express 6] (eml)
	FileType_.Store("d0cf11e0a1b11ae10000", "doc")  //MS Excel 注意：word、msi 和 excel的文件头一样
	FileType_.Store("d0cf11e0a1b11ae10000", "vsd")  //Visio 绘图
	FileType_.Store("5374616E64617264204A", "mdb")  //MS Access (mdb)
	FileType_.Store("252150532D41646F6265", "ps")
	FileType_.Store("255044462d312e350d0a", "pdf")  //Adobe Acrobat (pdf)
	FileType_.Store("2e524d46000000120001", "rmvb") //rmvb/rm相同
	FileType_.Store("464c5601050000000900", "flv")  //flv与f4v相同
	FileType_.Store("00000020667479706d70", "mp4")
	FileType_.Store("49443303000000002176", "mp3")
	FileType_.Store("000001ba210001000180", "mpg") //
	FileType_.Store("3026b2758e66cf11a6d9", "wmv") //wmv与asf相同
	FileType_.Store("52494646e27807005741", "wav") //Wave (wav)
	FileType_.Store("52494646d07d60074156", "avi")
	FileType_.Store("4d546864000000060001", "mid") //MIDI (mid)
	FileType_.Store("504b0304140000000800", "zip")
	FileType_.Store("526172211a0700cf9073", "rar")
	FileType_.Store("235468697320636f6e66", "ini")
	FileType_.Store("504b03040a0000000000", "jar")
	FileType_.Store("4d5a9000030000000400", "exe")        //可执行文件
	FileType_.Store("3c25402070616765206c", "jsp")        //jsp文件
	FileType_.Store("4d616e69666573742d56", "mf")         //MF文件
	FileType_.Store("3c3f786d6c2076657273", "xml")        //xml文件
	FileType_.Store("494e5345525420494e54", "sql")        //xml文件
	FileType_.Store("7061636b616765207765", "java")       //java文件
	FileType_.Store("406563686f206f66660d", "bat")        //bat文件
	FileType_.Store("1f8b0800000000000000", "gz")         //gz文件
	FileType_.Store("6c6f67346a2e726f6f74", "properties") //bat文件
	FileType_.Store("cafebabe0000002e0041", "class")      //bat文件
	FileType_.Store("49545346030000006000", "chm")        //bat文件
	FileType_.Store("04000000010000001300", "mxp")        //bat文件
	FileType_.Store("504b0304140006000800", "docx")       //docx文件
	FileType_.Store("d0cf11e0a1b11ae10000", "wps")        //WPS文字wps、表格et、演示dps都是一样的
	FileType_.Store("6431303a637265617465", "torrent")
	FileType_.Store("6D6F6F76", "mov")         //Quicktime (mov)
	FileType_.Store("FF575043", "wpd")         //WordPerfect (wpd)
	FileType_.Store("CFAD12FEC5FD746F", "dbx") //Outlook Express (dbx)
	FileType_.Store("2142444E", "pst")         //Outlook (pst)
	FileType_.Store("AC9EBD8F", "qdf")         //Quicken (qdf)
	FileType_.Store("E3828596", "pwl")         //Windows Password (pwl)
	FileType_.Store("2E7261FD", "ram")         //Real Audio (ram)
}

// 获取前面结果字节的二进制
func bytesToHexString(src []byte) string {
	res := bytes.Buffer{}
	if src == nil || len(src) <= 0 {
		return ""
	}
	temp := make([]byte, 0)
	for _, v := range src {
		sub := v & 0xFF
		hv := hex.EncodeToString(append(temp, sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0), 10))
		}
		res.WriteString(hv)
	}
	return res.String()
}

func GetFileType(fSrc []byte) string {
	var fileType string
	fileCode := bytesToHexString(fSrc)
	fmt.Println(fileCode)
	FileType_.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(string)
		if strings.HasPrefix(fileCode, strings.ToLower(k)) ||
			strings.HasPrefix(k, strings.ToLower(fileCode)) {
			fileType = v
			return false
		}
		return true
	})
	return fileType
}
