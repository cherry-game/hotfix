# hotfix
基于yaegi + gomonkey技术，在运行时支持热刷go脚本，可动态替换函数、属性。

## 原理
- 使用[yaegi](https://github.com/traefik/yaegi)动态执行`go`脚本。
- 使用[gomonkey](https://github.com/agiledragon/gomonkey)进行函数打桩，完成函数替换。
- [monkey原理](https://bou.ke/blog/monkey-patching-in-go/)技术实现细节。

## 支持平台
- 架构
  - amd64
  - arm64
  - 386
  - loong64
- 操作系统
  - Linux
  - MAC OS X
  - Windows

## 测试
- 找到`test/foo_test.go`文件，运行`TestFixFooHelloFunc()`。