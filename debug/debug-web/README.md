
go mod 


go mod init xx 这里的xx 是提交到github 或者gitlab的路径， 不要乱填， 

如果实在不知道怎么填， 

只需要在项目根目录下执行 go mod init 就会自动创建一个路径， 如
➜  debug-web git:(master) ✗ go mod init
go: creating new go.mod: module github.com/researchlab/gbp/debug/debug-web
go: to add module requirements and sums:
        go mod tidy

上面创建之后， vscode 打开引入的包就不会显示红色了， 如果还有就执行上面的go mod tidy 即可


viscode 中的文件是 .vscode的配置内容， 仅用于vscode  
