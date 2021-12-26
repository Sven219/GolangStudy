package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Username string `ini:"username"`
}
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

// Config
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0. 参数的校验
	// 0.1 传进来的 data 参数必须是指针类型（因为要在函数中对其赋值）
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer") //格式化输出之后返回一个 error 类型
		return
	}
	// 0.2 传进来的 data 参数必须是结构体类型指针
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param be a struct pointer") //格式化输出之后返回一个 error 类型
		return
	}
	//1. 读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	// 将字节类型的文件内容转换成字符串
	lineSlice := strings.Split(string(b), "\n")
	fmt.Printf("%#v\n", lineSlice)
	//2. 一行一行的读数据
	var structName string
	for idx, line := range lineSlice {
		// 去掉字符串收尾的空格
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		// 2.1 如果是注释就忽略
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2 如果[]开头表示是节（section）
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' && line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 把这一行首位的[] 去掉，取到中间的内容把首位的空格去掉拿到内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 根据字符串sectionName去 data 里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找到了对于的嵌套结构体
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 2.3 如果不是[]开头就是=分割的键值对
			// 1. 以=分割，左边为 key 右边为 value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			// 2. 根据 structName 去 data 里面把对应的嵌套结构体给取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) // 拿到嵌套结构体的值信息
			sType := sValue.Type()                     // 拿到嵌套结构体的类型信息
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data 中的%s 字段应该是一个结构体\n", structName)
				return
			}
			var fieldName string
			var fileType reflect.StructField
			//3 .遍历嵌套结构体的每一个字段判断 tag 是不是等于 key
			for i := 0; i < sValue.NumField(); i++ {
				filed := sType.Field(i)
				fileType = filed
				if filed.Tag.Get("ini") == key {
					// 找到对应的字段
					fieldName = filed.Name
					break
				}
			}
			// 4. 如果 key=tag ,给这个字段赋值
			// 4.1 根据 fieldName 取出这个字段
			if len(fieldName) == 0 {
				// 在结构体中找不到对应的字段
				continue
			}
			fileObj := sValue.FieldByName(fieldName)
			// 4.2 赋值
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int64, reflect.Int32:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return
				}
				fileObj.SetInt(valueInt)
			}

		}

	}

	return
}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)

	if err != nil {
		fmt.Printf("Load ini failed ,err:%v\n", err)
		return
	}
	fmt.Println(cfg)
}
