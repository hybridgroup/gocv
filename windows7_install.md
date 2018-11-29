参考 ：https://gocv.io

  WINDOWS 7  Installing

This page has information on how to install and use GoCV on Microsoft Windows 7, 64-bit.

#  Installing 安装 
Install the GoCV package: 
安装GoCV包：

    `go get -u -d gocv.io/x/gocv`

**In order to use GoCV on Windows you must build and install OpenCV  4.0.0. First download and install MinGW-W64 and CMake, as follows.
要在Windows上使用GoCV，您必须构建并安装OpenCV 4.0.0。首先下载并安装MinGW-W64和CMake，如下所示。**

#### MinGW-W64

Download and run the MinGW-W64 compiler installer from [https://sourceforge.net/projects/mingw-w64/?source=typ_redirect](https://sourceforge.net/projects/mingw-w64/?source=typ_redirect).
从[https://sourceforge.net/projects/mingw-w64/?source=typ_redirect](https://sourceforge.net/projects/mingw-w64/?source=typ_redirect)下载并运行MinGW-W64编译器安装程序。

**The first step is to choose Files.第一步，选Files**
![image.png](https://upload-images.jianshu.io/upload_images/5367714-b36e75fda35b0566.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)



**Second step.第二步**
The latest version of the MinGW-W64 toolchain is `7.3.0`, but any version from `7.X` on should work.
MinGW-W64工具链的最新版本是`7.3.0`，但是任何版本的`7.X`应该都可以使用。

Choose the options for “posix” threads, and for “seh” exceptions handling, then install to the default location `c:\Program Files\mingw-w64\x86_64-7.3.0-posix-seh-rt_v5-rev2.`
选择“posix”线程的选项，并选择“seh”异常处理，然后安装到默认位置 `c:\Program Files\mingw-w64\x86_64-7.3.0-posix-seh-rt_v5-rev2`。

Add the `C:\Program Files\mingw-w64\x86_64-7.3.0-posix-seh-rt_v5-rev2\mingw64\bin `path to your System Path.
添加 `C:\Program Files\mingw-w64\x86_64-7.3.0-posix-seh-rt_v5-rev2\mingw64\bin` 系统路径的路径。




![image.png](https://upload-images.jianshu.io/upload_images/5367714-edf31ae00c069c2d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### CMake
Download and install CMake [https://cmake.org/download/](https://cmake.org/download/) to the default location. CMake installer will add CMake to your system path.
下载并安装CMake [https://cmake.org/download/](https://cmake.org/download/)到默认位置。CMake安装程序会将CMake添加到您的系统路径。


![image.png](https://upload-images.jianshu.io/upload_images/5367714-649332cebec2459c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

**add CMake to your system path**
**记得添加CMake环境变量**

#### OpenCV 4.0.0 and OpenCV Contrib Modules

1.Click to download https://github.com/opencv/opencv/archive/4.0.0.zip and extract it to `c:\opencv\`
1.点击下载https://github.com/opencv/opencv/archive/4.0.0.zip,并解压到`c:\opencv\`


2.Click to download https://github.com/opencv/opencv_contrib/archive/4.0.0.zip and extract it to`c:\opencv\`
2.点击下载https://github.com/opencv/opencv_contrib/archive/4.0.0.zip,并解压到`c:\opencv\`

![image.png](https://upload-images.jianshu.io/upload_images/5367714-3333bea9f09de684.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


The following commands should do everything to download and install OpenCV 4.0.0 on Windows:
以下命令应该可以在Windows上下载和安装OpenCV 4.0.0(进到%GOPATH%\src\gocv.io\x\gocv)：

    chdir %GOPATH%\src\gocv.io\x\gocv        
    win7_build_opencv.cmd

It will probably take at least 1 hour to download and build.
下载和构建可能至少需要1个小时。

Last, add `C:\opencv\build\install\x64\mingw\bin` to your System Path.
最后，添加 C:\opencv\build\install\x64\mingw\bin 到你的系统路径。

####Verifying the installation
验证安装


Change the current directory to the location of the GoCV repo:
将当前目录更改为GoCV仓库的位置：

    chdir %GOPATH%\src\gocv.io\x\gocv

Now you should be able to build or run any of the command examples:
现在您应该能够构建或运行任何命令示例：
   
     go run cmd\version\main.go

The version program should output the following:
版本程序应输出以下内容：
    
    gocv version: 0.18.0
    opencv lib version: 4.0.0

That’s it, now you are ready to use GoCV.
就是这样，现在你已经准备好使用GoCV了。

![image.png](https://upload-images.jianshu.io/upload_images/5367714-2804270bf0a35256.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
