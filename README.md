# Intro

this is a simple CLI tool for TypeScript Projects management.


# Usage

```

tsp [projectDictionory] [mainClassName] [ProjectName]

```

# Sample:

```

tsp  test12 Test12Class Project12

```

it will generate a project in `test12` named `Project12` and  its main Class named `Test12Class`:
![a screenshot of a TypeScript Project](https://raw.githubusercontent.com/suifengtec/tsp/master/screenshot.png)


after the project generated,you can do somethig with files in `src`, then open two CLI windows:

```
gulp
```
it will watch your files.

the another CLI window will be used to run a simple http serve for your project,such as :
```
http-server  -p 6666 -a 127.0.0.1 --cors -o 
```

