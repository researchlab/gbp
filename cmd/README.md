
Go - Combining cmd.StdoutPipe and Cmd.StderrPipe


solution#1 

合并后实时打印信息

Use the function io.MultiReader to combine multiple readers into one:

```
outReader, err := cmd.StdoutReader()
if err != nil {
    // handle err
}

errReader, err := cmd.StderrReader()
if err != nil {
    // handle err
}

cmdReader := io.MultiReader(outReader, errReader)
```


solution#2 

合并后实时打印信息

To combine stdout and stderr to a single reader, assign a single pipe to Command.Stdout and Command.Stderr:
```
cmdReader, err := cmd.StdoutPipe()
cmd.Stderr = cmd.Stdout
```

solution#3 

合并后 统一执行完成后再打印信息

```
	cmd = exec.Command("/bin/bash", "-c", "ls /")
	res, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(res))
		panic(err)
	}
	fmt.Println(string(res))

```
