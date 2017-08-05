// Copyright © 2017 suifengtec <suifengtec@qq.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// go run main.go
/*

go build -o a.exe main.go && a A

a aa bb cc


可能有用
Run a command when files change


*/
package main

import (
	"bytes"
	"fmt"
	. "github.com/Kenshin/cprint"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"tsp2/cmd"
)

const (
	VERSION      = "1.2"
	CMD_HELP     = "h"
	CMD_VERSION  = "v"
	CMD_GENERATE = "g"
	CMD_WATCH    = "w"
	CMD_SERVE    = "s"
)

func gatherArgs() {

	args := os.Args[1:]

	if len(args) < 1 {
		IntroMe()
	} else {

		argsLen := len(args)

		switch {

		case argsLen == 1:
			Arg1(args)

		case argsLen == 2:
			Arg2(args)
		case argsLen == 3:
			Arg3(args)
		case argsLen == 4:

			Arg4(args)

		}

	}

}

func Arg1(args []string) {

	switch args[0] {

	case CMD_HELP:
		IntroMe()

	case CMD_VERSION:
		P(NOTICE, "%v", VERSION)
	}
}

func Arg2(args []string) {

	switch args[0] {
	case CMD_WATCH:
		watchProject(args)
	}

}
func Arg3(args []string) {

	switch args[0] {

	case CMD_SERVE:
		serveProject(args)
	}
}

/*
tsp g pSlug dName pName
*/
func Arg4(args []string) {

	switch args[0] {

	case CMD_GENERATE:
		generateProject(args)

	}
}

/*
go build -o a.exe main.go && a s aa 6611

*/

func serveProject(args []string) {
	var (
		dirName string = filepath.Join(args[1], "dist")
		port    string = args[2]
	)

	valid, _ := isPortValid(port)

	if !valid {
		P(WARING, "Invalid port %v .", port)
	}

	if !IsPathExist(joinPath(dirName)) {

		P(WARING, "dir %v not exist.", dirName)
		return
	}
	if err := os.Chdir(dirName); err != nil {

		P(WARING, "切换工作目录时出错")
		return
	}

	cmd := exec.Command("http-server", "-p", port, "-a", "127.0.0.1", "--cors", "-o")

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Print("服务器已启动")
	fmt.Print(string(cmdOutput.Bytes()))

}

/*

go build -o a.exe main.go && a w aa


golang command line Long running processes


*/
func watchProject(args []string) {

	var (
		dirName string = args[1]
	)
	if !IsPathExist(joinPath(dirName)) {

		P(WARING, "dir %v not exist.", dirName)
		return

	}

	if err := os.Chdir(dirName); err != nil {
		P(WARING, "dir %v can NOT be read.", dirName)
		return
	}

	cmd := exec.Command("gulp")

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	fmt.Print("项目已打包")
	fmt.Print(string(cmdOutput.Bytes()))

}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

/*

a g aa bb cc

go build -o a.exe main.go && a g aa bb cc
*/
func generateProject(args []string) {

	dirName, className, projectName := args[1], args[2], args[3]

	/*	fmt.Println(dirName, className, projectName)
	 */
	cPath := getCurrentPath()

	newProjectDir := joinPath(cPath, dirName)

	if IsPathExist(newProjectDir) == true {

		P(WARING, "dir '"+newProjectDir+"' exists, I will try to delete it")

		os.RemoveAll(newProjectDir)

		/*return*/
	}
	err := os.MkdirAll(newProjectDir, os.ModePerm|os.ModeDir)
	CheckError(err)
	fmt.Println("...")
	P(NOTICE, "dir '"+newProjectDir+"'  has be recreated.")

	var (
		bashScript string = filepath.Join("_init.sh")
		shFilePath string = filepath.Join(newProjectDir, bashScript)
	)

	file, err := os.Create(shFilePath)
	CheckError(err)
	defer file.Close()

	file.WriteString("#!/bin/sh\n")

	commands := []string{
		/*
		 tsconfig.json
		*/
		"echo '{ \n\t\"files\": [ \"src/main.ts\"], \n\t\"compilerOptions\": { \n\t\t\"noImplicitAny\": true,\n\t\t\"target\": \"es6\" \n\t}\n}' >> tsconfig.json",
		/*
			gulpfile.js
		*/
		"echo '\nvar gulp=require(\"gulp\");\nvar browserify=require(\"browserify\");\nvar source=require(\"vinyl-source-stream\");\nvar watchify=require(\"watchify\");\nvar tsify=require(\"tsify\");\nvar gutil=require(\"gulp-util\");\nvar paths={pages:[\"src/*.html\"]};\nvar watchedBrowserify=watchify(browserify({basedir:\".\",debug:true,entries:[\"src/main.ts\"],cache:{},packageCache:{}}).plugin(tsify));\ngulp.task(\"copy-html\",function(){\n\treturn gulp.src(paths.pages).pipe(gulp.dest(\"dist\"))\n});\nfunction bundle(){\n\treturn watchedBrowserify.bundle().pipe(source(\"bundle.js\")).pipe(gulp.dest(\"dist\"))\n}\ngulp.task(\"default\",[\"copy-html\"],bundle);\nwatchedBrowserify.on(\"update\",bundle);\nwatchedBrowserify.on(\"log\",gutil.log);' >> gulpfile.js",

		/*
			/src
		*/
		"mkdir src && cd src && touch index.html && echo '<!DOCTYPE html>\n<html lang=\"zh_CN\">\n \t<head>\n\t\t<meta charset=\"utf-8\">\n\t\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0, maximum-scale=1\">\n\t\t<title>" + projectName + "</title>\n\t</head>\n\t<body>\n\t\t<p id=\"sayhi-to-ts\">" + projectName + " tesing ...</p>\n\t\t<script src=\"bundle.js\"></script>\n\t</body>\n</html>' >> index.html && touch main.ts && echo '/*\n\tThis is the main file.\n*/\nclass " + className + "{ \n\t constructor(){}\n}' >> main.ts && cd ..",

		"mkdir dist",
		"npm init --y",

		/*npm
		gulp-cli
		browserify tsify vinyl-source-stream
		watchify gulp-util
		typescript gulp gulp-typescript
		http-server

		comment it when dev.
		*/

		"npm install  --save-dev gulp && npm install -g gulp-cli && npm install --save-dev browserify tsify vinyl-source-stream  && npm install --save-dev watchify gulp-util && npm i --save-dev typescript gulp gulp-typescript && npm install http-server --save-dev",

		"echo '/node_modules/*' >>.gitignore",
		"git init && git config --local core.autocrlf false && git add .  && git commit -m \"init\" ",

		/*
			从用户输入中获取项目名称
		*/
		"echo '#  " + projectName + "' >> README.md",
	}

	/*	file.WriteString( "http-server -p " + port + "  -a 127.0.0.1 --cors -o\n"))
	 */
	file.WriteString(strings.Join(commands, "\n"))
	err = os.Chdir(newProjectDir)
	CheckError(err)

	out, err := exec.Command("sh", bashScript).Output()
	CheckError(err)
	fmt.Println(string(out))
}

/*=======================================*/
// Y:\DevSpace\Go\src\tsp
func getCurrentPath() string {
	path, err := os.Getwd()
	if err != nil {

		defer func() {

			if er := recover(); er != nil {

				P(ERROR, err.Error())
				os.Exit(0)
			}

		}()
		panic("get current path Error: " + err.Error())
	}
	return path
}

/*
 判断目录或文件是否存在

 Param:
    - paths: 切片形式

 Return:
    - true : exist
    - false: no exit
*/
func IsPathExist(paths ...string) bool {
	path := filepath.Join(paths...)
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

/*
Join Path( folder )
 Param:
    - paths: one or two valid path e.g. "someProjectDirName" OR "Y:\\DevSpace\\Go\\src\\tsp" , "someProjectDirName"

 Return:
    - true : exist
    - false: no exit
	path1 := joinPath("cmd")
	path2 := joinPath(cPath, "cmd1")
*/
func joinPath(paths ...string) string {

	var (
		newPath string
		pLen    = len(paths)
	)
	if pLen == 1 {
		cPath := getCurrentPath()

		newPath = filepath.Join(cPath, paths[0])
	} else {
		newPath = filepath.Join(paths[0], paths[1])
	}

	return newPath
}

/*
错误检查
*/
func CheckError(e error) {

	if e != nil {
		defer func() {

			if err := recover(); err != nil {
				P(ERROR, e.Error())
				os.Exit(0)
			}

		}()
		panic(e)

	}
}

/*
恢复错误

*/
func mayRecover(args string) {

	if len(args) < 1 {
		args = "WARN: something wrong, pls restart this CLI window ."
	}

	defer func() {

		if err := recover(); err != nil {

			P(ERROR, args)
			os.Exit(0)
		}

	}()

}

func isPortValid(port string) (bool, error) {

	host := ":" + port

	server, err := net.Listen("tcp", host)

	if err != nil {
		P(WARING, "Invalid Port: %v\n", port)
		return false, err
	}
	server.Close()
	return true, nil

}

/*=======================================*/
func IntroMe() {

	P(NOTICE, "项目首页\t%v  \n", "https://github.com/suifengtec/tsp")
	P(NOTICE, "说明\n")
	cmd.Execute()

}

func main() {
	gatherArgs()
}
