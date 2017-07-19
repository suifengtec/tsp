/*
* @Author: Administrator
* @Date:   2017-07-19 09:22:49
* @Last Modified by:   Administrator
* @Last Modified time: 2017-07-19 10:27:07
*/

package main

/*
go build && tsp 
 */

import (
	"fmt"
    "os"
    "os/exec"
    "strings"
    "path/filepath"
)


var UsageTip = func(){

	fmt.Println("使用方法: tsp [项目目录名] [主类名] [项目名称] ")

}
func checkError( e error){
    if e != nil {
        panic(e)
        /*fmt.Println(e)*/
    }
}

func exe_cmd(cmds []string, dirName string ) {

/*	var(

			output_path = filepath.Join("./"+dirName)
			bash_script = filepath.Join( "_do.sh" )
		)  */

	var output_path = filepath.Join("./"+dirName)
	var bash_script = filepath.Join( "_script.sh" )
    

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
    fmt.Println(string(out))
}

func doAction (args []string){

		var(
			dirName = args[0]
			className = args[1]
			projectName = args[2]
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
	    "mkdir src && cd src && touch index.html && echo '<!DOCTYPE html>\n<html lang=\"zh_CN\">\n \t<head>\n\t\t<meta charset=\"utf-8\">\n\t\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0, maximum-scale=1\">\n\t\t<title>"+projectName+"</title>\n\t</head>\n\t<body>\n\t\t<p id=\"sayhi-to-ts\">测试 ...</p>\n\t\t<script src=\"bundle.js\"></script>\n\t</body>\n</html>' >> index.html && touch main.ts && echo '/*\n\t这是TS的类\n*/\nclass "+className+"{ \n\t constructor(){}\n}' >> main.ts && cd ..",

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

   exe_cmd(commands,dirName)

}


func main() {

	args := os.Args[1:]

	if args == nil || len(args)<3{

		UsageTip()

	}else{
		fmt.Println("****************开始******************")
		doAction (args)
		fmt.Println("****************完成******************")
		fmt.Println("项目"+ args[0]+"已生成!!! ")
	}

}

/*

go build && tsp  test11 Test11Class 测试项目11

*/