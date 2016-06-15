# go package详解

GO语言中package用于封装一个相对独立的功能提供给外部使用

1. package会对应一个目录 这点与Java类似

2. package中的源代码存放目录是package最后一个/结束的准，如package math/rand,那么所源代码都在rand目录下

3. main是一个特殊的package名字，类似Java的main函数，GO的可执行程序必须在main package下

4. 以大写字母开头的函数才会被导出（外部访问）。这点类似Java和访问权限控制，只是太隐晦了
