# GoBuilder

* [Q&amp;A](#qa)
    * [GoBuilder 是什么？](#gobuilder-是什么)
    * [为什么需要构建工具？](#为什么需要构建工具)
    * [为什么选择 GoBuilder？](#为什么选择-gobuilder)
    * [GoBuilder 与其他类似的工具有什么不同？](#gobuilder-与其他类似的工具有什么不同)
* [快速上手](#快速上手)
* [原理简述](#原理简述)
* [示例](#示例)
* [背景](#背景)
* [更新计划](#更新计划)
* [致谢](#致谢)
* [项目结构](#项目结构)
* [注意事项](#注意事项)

注：该目录使用 [gh-md-toc](https://github.com/ekalinin/github-markdown-toc.go) 创建。

## Q&A

### GoBuilder 是什么？

GoBuilder一个无侵入性的的Go构建工具，易于学习，支持不同平台的构建，支持自定义构建任务。

### 为什么需要构建工具？

构建工具可以更方便的将源代码转换为可执行文件。假设一个场景，你有一个多人项目，需要在`go build`项目前编译静态资源，你没法保证你自己或者其他人不会忘记去编译静态资源，使用构建工具后就可以保证这一点。

### 为什么选择 GoBuilder？

GoBuilder兼容任何Go项目，无论之前使用什么构建工具，是否使用go mod，并且GoBuilder是跨平台的软件。

### GoBuilder 与其他同类工具有什么不同？

| |跨平台|易于学习|与go tool chain兼容|
|:---:|:---:|:---:|:---:|
|GoBuilder|Y|Y|Y|
|Makefile|N|N|Y|
|Gogradle|Y|Y|N|

（如有错误请联系修改）

## 快速上手

本章将会带你从无到有完成一个可以输出`hello, name`的task，`name`可通过命令行指定，默认为`world`。 以此来让你初步了解GoBuilder是什么以及它的用法。

*注：为了更好地区分GoBuilder的构建功能与自带`go build`命令，该章节中将`go build`命令所执行的操作称为编译。*

**预备知识：**

构建（build）：将项目从源代码编程可执行文件的过程。    
任务（task）：构建过程中的原子操作。  
命令（command）：一个或多个任务的有序集合。

1. 下载并安装GoBuilder
   ```text
   > go get gitee.com/KongchengPro/GoBuilder
   ```

2. 检测是否安装成功

   ```text
   > GoBuilder info
   [INFO] GoBuilder info { version: 0.0.1, authors: [kongchengpro <kongchengpro@163.com>] }
   ```

3. 初始化GoBuilder项目，使用-p或-path可以指定项目路径（两者等效）。
   ```text
   > GoBuilder init (-p|path) (路径)
   [INFO] init successfully { nil }
   ```
   *注：init 命令是在指定目录下初始化，不会在指定目录下新建目录再去新建目录里初始化，这点与其他构建工具不同。*

4. 开始编写task，进入`gobuilder/tasks`目录，创建`SayHello`目录，在里面创建`main.go`。
   `SayHello`这个目录名就是task名称，可以自己改，后面的代码中也要对应改掉。`main.go`的文件名可以随便改。
   ```text
   > cd gobuilder/tasks
   > mkdir SayHello
   > cd SayHello
   > touch main.go
   ```
   `main.go`的文件内容请 [点击此处](https://gitee.com/KongchengPro/GoBuilder/blob/main/example/gobuilder/tasks/SayHello/main.go)
   查看

5. 添加并运行task，先回到项目目录
   ```text
   > cd ../../
   > GoBuilder addTask SayHello
   > GoBuilder runTask SayHello
   hello, world
   ```
   `addTask`命令将SayHello添加到了项目的task列表中
   `runTask`命令运行已添加的task

   *注：拆分成两个命令是因为addTask需要花时间去编译，拆分出来可以达到一次编译多次运行的效果，如果需要编译后运行，可以使用`GoBuilder addTaskAndRun`命令。*

   可以通过命令行改变第二个单词
   ```text
   > GoBuilder runTask SayHello GoBuilder
   hello, GoBuilder
   ```

## 原理简述

这里仅简单说明原理，帮助理解GoBuilder的运作，具体实现可自行阅读源码。

addTask - 编译`gobuilder/tasks/taskName/`，输出到`gobuilder/.executable/`目录

runTask - 运行`gobuilder/.executable/taskName(.exe)`并传入json序列化的调用信息`tdk.TaskCaller`

## 示例

[实现批量交叉编译的task](https://gitee.com/KongchengPro/GoBuilder/blob/main/example/gobuilder/tasks/build-x/main.go)

## 背景

**以下以项目发起者第一人称叙述**

我在编写OPQ-OneBot项目时，有如下需求：将项目批量交叉编译为不同平台的可执行文件，并且在编译前自动执行编译静态资源的命令。  
结果搞了一下午，试了好几个工具，都没倒腾出来【捂脸】。  
于是乎……emm……就一言不合造轮子了

## 更新计划

- [ ] 增加命令系统，命令是一个或多个任务的有序集合。
- [ ] 增加内置task。
- [ ] 将task中常用的代码封装成函数。

## 致谢

GoBuilder 的诞生离不开以下开源项目：

<https://github.com/sirupsen/logrus>

<https://github.com/urfave/cli/v2>

## 项目结构

```text
GoBuilder: 项目根目录
    internal: 不供其他应用程序或库使用的代码
        app: 应用程序代码
        pkg: 可供该应用程序的其他包使用的代码（不供其他应用程序和库使用）
    pkg: 可供其他应用程序或库使用的代码
        tdk: Task Development Kit（任务开发工具包）的缩写，供 task 使用
    test: 测试（阅读源码时可以忽略该文件夹）
        *_test.go: 测试代码
    example: 示例项目
        gobuilder: GoBuilder相关的文件
            .executable: 执行 GoBuilder addTask 后编译好的 task 可执行文件
            tasks: GoBuilder 的 task
            cmd.gb: GoBuilder 中定义命令的文件（命令系统暂未开发，目前无用，但该文件不可删去）
        main.go: 程序主文件
``` 

## 注意事项

GoBuilder在1.0.0版本前的api并不稳定，请勿用于生产项目！