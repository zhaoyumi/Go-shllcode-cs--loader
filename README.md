# Go-shellcode(cs)-loader

<u>2022.8.26</u>

原项目：https://github.com/HZzz2/go-shellcode-loader

#### 使用

在cs中生成shellcode（C格式文件），将 \x 全部替换为空
`// \xfc\x48\x83\xe4\xf0...  替换为  fc4883e4f0...`

把替换后的shellcode进行一次base64加密
`//ZmM0ODgzZTRmMA...`

复制到 aes-encrypt.go 的44行，base64-payload处。

![image-20220826091109563](img\image-20220826091109563.png)

执行 aes-encrypt.go 生成 AES 加密的 payload

```
go run aes-encrypt.go
```

![image-20220826091613768](img\image-20220826091613768.png)

复制输出的值，替换 cs-loader.go 的58行 AES-payload

![image-20220826091845149](img\image-20220826091845149.png)

编译成exe文件即可（有黑框）~~（过 360、火绒）~~

```
go build -o shell.exe cs-loader.go
```

编译后执行无黑框 ~~（过火绒，不过360）~~

```
go build -ldflags="-w -s -H windowsgui" -o shell.exe cs-loader.go
```

注：

***直接使用 go 编译会有个黑框（cmd）一直开着，关闭后会直接kill掉进程，但是不会被 360 查杀，可以根据自己的情景进行编译***

也可以使用 garble 进行混淆编译，（免杀效果会更好点）：

```
//有黑框
garble -tiny -literals -seed=random build cs-loader.go
//无黑框
garble -tiny -literals -seed=random build -ldflags="-w -s -H windowsgui" cs-loader.go
```

#### 效果

![image-20220826103413391](img\image-20220826103413391.png)

![image-20220826103454683](img\image-20220826103454683.png)

![image-20220826103850958](img\image-20220826103850958.png)