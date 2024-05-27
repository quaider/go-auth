## cmd目录
存放应用的入口，可以把组件 main 函数所在的文件夹统一放在/cmd 目录下

每个组件的目录名应该跟可执行文件名是一致的。

这里要保证 /cmd/<组件名> 目录下不要存放太多的代码，如果你认为代码可以导入并在其他项目中使用，那么它应该位于 /pkg 目录中。

如果代码不是可重用的，或者你不希望其他人重用它，请将该代码放到 /internal 目录中。

**通常有一个小的 main 函数，从 /internal 和 /pkg 目录导入和调用代码，除此之外没有别的东西**。

如 https://github.com/vmware-tanzu/velero/blob/main/cmd/velero/velero.go

`cmd/servicename/servicename.go`，其中 servicename.go 就是 该服务的 main函数。