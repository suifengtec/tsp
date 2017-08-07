# Intro

this is a simple CLI tool for TypeScript Projects management.


# Usage Dependencies

```

Sublime Text
Node.js/ npm
TypeScript
Git

```


# Usage

to generate a TypeScript project:

```

tsp g [projectDirName] [mainClassName] [ProjectName]

```

to run a simple http server for your project:

```

tsp s [projectDirName] [port]

```

to run gulp for your project:
```

tsp w [projectDirName]

```

`port` should be a number  in range (80,65535).



# Sample:

```

tsp g test12 Test12Class Project12

```

it will generate a project in `test12` named `Project12` and  its main Class named `Test12Class`:
![a screenshot of a TypeScript Project](https://raw.githubusercontent.com/suifengtec/tsp/master/screenshot-1.png)
![a screenshot of a TypeScript Project](https://raw.githubusercontent.com/suifengtec/tsp/master/screenshot-2.png)
![a screenshot of a TypeScript Project](https://raw.githubusercontent.com/suifengtec/tsp/master/screenshot-3.png)


after the project generated,you can do something with files in `src`, then open two CLI windows, execute the following command in a CLI window:

```

tsp w test12

```
it will watch your files.

> Note: you can also change CLI to the project base dir,and execute `gulp` to do so.

the another CLI window will be used to run a simple http serve for your project,such as :

```
tsp s test12 6666
```

tsp will run a http server for your TypeScript project in `test12` and open it in your default browser.

![a screenshot of a TypeScript Project](https://raw.githubusercontent.com/suifengtec/tsp/master/screenshot-4.png)


# Dev Dependencies

```

go get -u github.com/Kenshin/cprint

go get -u github.com/spf13/cobra

```


# Homepage

[here](http://coolwp.com/cli-tool-typescript-projects.html)