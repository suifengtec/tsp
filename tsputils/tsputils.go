/*
* @Author: suifengtec
* @Date:   2017-08-06 14:10:57
* @Last Modified by:   suifengtec
* @Last Modified time: 2017-08-06 15:05:33
 */

package tsputils2

import (
	"bytes"
	"fmt"
	. "github.com/Kenshin/cprint"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const VERSION string = "1.2.0"

func ShowVersion() {
	fmt.Println(VERSION)
}
func WatchProject(args []string) {

	var (
		dirName string = args[0]
	)
	if !IsPathExist(JoinPath(dirName)) {

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

func ServeProject(args []string) {

	var (
		projectDir string = args[0]
		port       string = args[1]
		dirName    string = filepath.Join(projectDir, "dist")
	)

	valid, _ := IsPortValid(port)

	if !valid {
		P(WARING, "Invalid port %v .\n", port)
	}

	if !IsPathExist(JoinPath(dirName)) {

		P(WARING, "dir %v not exist.\n", dirName)
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
	P(NOTICE, "服务器已启动")
	fmt.Print(string(cmdOutput.Bytes()))

}

func IsPortValid(port string) (bool, error) {

	host := ":" + port

	server, err := net.Listen("tcp", host)

	if err != nil {
		P(WARING, "Invalid Port: %v\n", port)
		return false, err
	}
	server.Close()
	return true, nil

}

func GenerateProject(args []string) {

	dirName, className, projectName := args[0], args[1], args[2]

	/*	fmt.Println(dirName, className, projectName)
	 */
	cPath := GetCurrentPath()

	newProjectDir := JoinPath(cPath, dirName)

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
		"mkdir src && cd src && touch index.html && echo '<!DOCTYPE html>\n<html lang=\"zh_CN\">\n \t<head>\n\t\t<meta charset=\"utf-8\">\n\t\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0, maximum-scale=1\">\n\t\t<link rel=\"icon\" href=\"data:;base64,iVBORw0KGgo=\" type=\"image/x-icon\">\n\t\t<title>" + projectName + "</title>\n\t\t</head>\n\t<body>\n\t\t<p id=\"sayhi-to-ts\">" + projectName + " tesing ...</p>\n\t\t<script src=\"bundle.js\"></script>\n\t</body>\n</html>' >> index.html && touch main.ts && echo '/*\n\tThis is the main file.\n*/\nclass " + className + "{ \n\t constructor(){}\n}' >> main.ts && cd ..",

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
func JoinPath(paths ...string) string {

	var (
		newPath string
		pLen    = len(paths)
	)
	if pLen == 1 {
		cPath := GetCurrentPath()

		newPath = filepath.Join(cPath, paths[0])
	} else {
		newPath = filepath.Join(paths[0], paths[1])
	}

	return newPath
}

func GetCurrentPath() string {
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
