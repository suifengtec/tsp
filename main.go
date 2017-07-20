/*
* @Author: Administrator
* @Date:   2017-07-19 09:22:49
* @Last Modified by:   Administrator
* @Last Modified time: 2017-07-20 16:42:03
*/

package main

/*
go build && tsp 
 */

import (
	"fmt"
	/*"io"*/
    "os"
    "os/exec"
    "strings"
    "path/filepath"
    "strconv"
/*    _"tsp/utils"*/
)


func checkError( e error){
    if e != nil {
        panic(e)
        /*fmt.Println(e)*/
    }
}


func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func printMe(str string){

	fmt.Println(str)

}

func deleteFile(path string) {

	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		return;
	}
	err = os.Remove(path)
	if isError(err) { 
		return 
	}
}



func usageTip(){

	printMe("to generate a TypeScript project:\n******\ntsp [projectDirName] [ProjectMainClassName] [ProjectName] \n******\nto run s simple http server for your project: \n******\ntsp s [port]\n******\nto run gulp for your project:\n******\ntsp w [projectDirName]\n******\n");

}

/*
tsp g projectDir MainClassName ProjectName


 */
func exeCmd(cmds []string, dirName string ) {


			var(
				output_path = filepath.Join("./"+ dirName)
				bash_script = filepath.Join( "_do.sh" )
			)  
			os.RemoveAll(output_path)
			err := os.MkdirAll( output_path, os.ModePerm|os.ModeDir )
			checkError(err)
			file, err := os.Create( filepath.Join(output_path, bash_script))
			checkError(err)
			defer file.Close()
			file.WriteString("#!/bin/sh\n")
			file.WriteString( strings.Join(cmds, "\n"))
			err = os.Chdir(output_path)
			checkError(err)
			out, err := exec.Command("sh", bash_script).Output()
			checkError(err)

			printMe(string(out))


}


func runGulp(dirName string){

			var(
				output_path = filepath.Join("./"+ dirName)
				bash_script = filepath.Join( "_gulp.sh" )
			  
				cmds = [] string{
					"gulp",
				}
				shFilePath = filepath.Join(output_path, bash_script)
			)

			deleteFile(shFilePath)

			err := os.MkdirAll( output_path, os.ModePerm|os.ModeDir )
			checkError(err)
			file, err := os.Create( shFilePath )
			checkError(err)
			defer file.Close()
			file.WriteString("#!/bin/sh\n")
			file.WriteString( strings.Join(cmds, "\n"))
			err = os.Chdir(output_path)
			checkError(err)
			out, err := exec.Command("sh", bash_script).Output()
			checkError(err)

			printMe(string(out))


}


func simpleHttpServer(dirName string,  port string){

			var(
				output_path = filepath.Join("./"+ dirName+"/dist")
				bash_script = filepath.Join( "_s.sh" )
			  
				cmds = [] string{
					"http-server -p "+port+"  -a 127.0.0.1 --cors -o",
				}
				shFilePath = filepath.Join(output_path, bash_script)
			)

			deleteFile(shFilePath)

			err := os.MkdirAll( output_path, os.ModePerm|os.ModeDir )
			checkError(err)
			file, err := os.Create( shFilePath )
			checkError(err)
			defer file.Close()
			file.WriteString("#!/bin/sh\n")
			file.WriteString( strings.Join(cmds, "\n"))
			err = os.Chdir(output_path)
			checkError(err)
			out, err := exec.Command("sh", bash_script).Output()
			checkError(err)

			printMe(string(out))





}



func doAction (args []string){

	if nil == args{
		usageTip()
		return;
	}

	var mainCmd = args[0]

	switch(mainCmd){

		case "g":

			printMe("****************start: maybe need some time to get npm packages******************")
			var(

				dirName = args[1]
				className = args[2]
				projectName = args[3]
			) 


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
				"mkdir src && cd src && touch index.html && echo '<!DOCTYPE html>\n<html lang=\"zh_CN\">\n \t<head>\n\t\t<meta charset=\"utf-8\">\n\t\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0, maximum-scale=1\">\n\t\t<title>"+projectName+"</title>\n\t</head>\n\t<body>\n\t\t<p id=\"sayhi-to-ts\">测试 ...</p>\n\t\t<script src=\"bundle.js\"></script>\n\t</body>\n</html>' >> index.html && touch main.ts && echo '/*\n\tThis is the main file.\n*/\nclass "+className+"{ \n\t constructor(){}\n}' >> main.ts && cd ..",

				"mkdir dist",
				"npm init --y",

				/*npm
					gulp-cli
					browserify tsify vinyl-source-stream
					watchify gulp-util
					typescript gulp gulp-typescript 
					http-server


				*/

				"npm install -g gulp-cli && npm install --save-dev browserify tsify vinyl-source-stream  && npm install --save-dev watchify gulp-util && npm i --save-dev typescript gulp gulp-typescript && npm install http-server --save-dev",


				"echo '/node_modules/*' >>.gitignore",
				"git init && git config --local core.autocrlf false && git add .  && git commit -m \"init\" ",

				/*
				从用户输入中获取项目名称
				 */
				"echo '#  "+projectName+"' >> README.md",

			}

			exeCmd(commands,dirName)

			printMe("****************done******************")
			printMe("project in dir "+ dirName +".\n you can run it using:\ntsp s "+ dirName +" 6666")

		case "s":
			port,error := strconv.Atoi(args[2])

			if error != nil{

				printMe("The third parameter must be a number in range: 1024-65535")

			}else{

				if(port<80 || port>65535){
					printMe("The third parameter must be a number in : 1024-65535")
				}else{
						printMe("Press Ctrl+C to close the server running by me.")
						simpleHttpServer(args[1],  string(args[2]))
					}
				

			}

		case "w":
		case "gulp":

			runGulp(string(args[1]));

		case "h":
		case "help":
		default:
			usageTip()
	}




}




func main() {

	args := os.Args[1:]

	if args == nil || len(args)<3{


		if args== nil {
			usageTip()
			return;
		}

		if args[0]=="v"{

			printMe("v1.1.0")
			
		}else{
			usageTip()
		}
		


	}else{
		
		doAction (args)

	}

}

/*

go build && tsp  test11 Test11Class 测试项目11

*/