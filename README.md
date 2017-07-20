# Intro

this is a simple CLI tool for TypeScript Projects management.


# Dependencies

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

`port` should be a number  in range (80,65535).



# Sample:

```

tsp g test12 Test12Class Project12

```

it will generate a project in `test12` named `Project12` and  its main Class named `Test12Class`:
![a screenshot of a TypeScript Project](https://raw.githubusercontent.com/suifengtec/tsp/master/screenshot.png)


after the project generated,you can do something with files in `src`, then open two CLI windows, execute the following command in a CLI window:

```
tsp gulp test12

//or

tsp w test12

```
it will watch your files.

the another CLI window will be used to run a simple http serve for your project,such as :

```
tsp s test12 6666
```

tsp will run a http server for your TypeScript project in `test12` and open it in your default browser.

# Homepage

[here](http://coolwp.com/cli-tool-typescript-projects.html)