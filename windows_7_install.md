#https://www.jianshu.com/p/2bf442857818

参考 ：https://gocv.io



#  1.安装
安装GoCV包：

    `go get -u -d gocv.io/x/gocv`

**要在Windows上使用GoCV，您必须构建并安装OpenCV 3.4.3。首先下载并安装MinGW-W64和CMake，如下所示。**

#### MinGW的-W64

从[https://sourceforge.net/projects/mingw-w64/?source=typ_redirect](https://sourceforge.net/projects/mingw-w64/?source=typ_redirect)下载并运行MinGW-W64编译器安装程序。

**第一步，选Files**
![image.png](https://upload-images.jianshu.io/upload_images/5367714-b36e75fda35b0566.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)



**第二步选
MinGW-W64工具链的最新版本是`7.3.0`，但是任何版本的`7.X`应该都可以使用。
选择“posix”线程的选项，并选择“seh”异常处理，然后安装到默认位置 `c:\Program Files\mingw-w64\x86_64-7.3.0-posix-seh-rt_v5-rev2`。添加 `C:\Program Files\mingw-w64\x86_64-7.3.0-posix-seh-rt_v5-rev2\mingw64\bin` 系统路径的路径。**

![image.png](https://upload-images.jianshu.io/upload_images/5367714-edf31ae00c069c2d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### CMake的

下载并安装CMake [https://cmake.org/download/](https://cmake.org/download/)到默认位置。CMake安装程序会将CMake添加到您的系统路径。


![image.png](https://upload-images.jianshu.io/upload_images/5367714-649332cebec2459c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

**记得添加CMake环境变量**

#### OpenCV 3.4.3和OpenCV Contrib模块
以下命令应该可以在Windows上下载和安装OpenCV 3.4.3：

    chdir %GOPATH%\src\gocv.io\x\gocv      #第一条执行的是： 切换到你的gopath下的\src\gocv.io\x\gocv
  
**win_build_opencv.cmd脚本有些命令windows 7并没有识别,我将脚本内容修改为下面脚本内容**

**%GOPATH%\src\gocv.io\x\gocv\win_build_opencv.cmd修改后的内容：**
```
echo echo是打印命令

echo 如果没存在就去创建了两个文件夹C:\opencv和C:\opencv\build

if not exist "C:\opencv" mkdir "C:\opencv"
if not exist "C:\opencv\build" mkdir "C:\opencv\build"


echo  上面删掉的内容可以自己离线下载并解压，后面附百度云链接

echo 切换到了C:\opencv\build
cd /D C:\opencv\build

echo 设置环境变量，如果设置过的话可以忽略

set PATH=%PATH%;C:\Program Files (x86)\CMake\bin;C:\mingw-w64\x86_64-6.3.0-posix-seh-rt_v5-rev1\mingw64\bin

echo 构建项目，路径需要核对好上面删除的下载两个压缩包解压的位置，不然构建会失败
cmake C:\opencv\opencv-3.4.3 -G "MinGW Makefiles" -BC:\opencv\build -DENABLE_CXX11=ON -DOPENCV_EXTRA_MODULES_PATH=C:\opencv\opencv_contrib-3.4.3\modules -DBUILD_SHARED_LIBS=ON -DWITH_IPP=OFF -DWITH_MSMF=OFF -DBUILD_EXAMPLES=OFF -DBUILD_TESTS=OFF -DBUILD_PERF_TESTS=OFF -DBUILD_opencv_java=OFF -DBUILD_opencv_python=OFF -DBUILD_opencv_python2=OFF -DBUILD_opencv_python3=OFF -DBUILD_DOCS=OFF -DENABLE_PRECOMPILED_HEADERS=OFF -DBUILD_opencv_saliency=OFF -DCPU_DISPATCH= -Wno-dev

mingw32-make -j%NUMBER_OF_PROCESSORS%
mingw32-make install

echo 移除删除文件夹，我把下面两条删了，根据自己喜好，硬盘太小的话就删掉
rmdir c:\opencv\opencv-3.4.3 /s /q
rmdir c:\opencv\opencv_contrib-3.4.3 /s /q


echo 切换到你设置的gopath\src\gocv.io\x\gocv
chdir /D %GOPATH%\src\gocv.io\x\gocv

```

修改完后执行

    win_build_opencv.cmd               #执行win_build_opencv.cmd脚本（脚本有问题,所以我大概把要下的给离线下载好并解压到了C:\openvc\下）


下载和构建可能至少需要1个小时，（猜想使用离线会快吧）。

重要，路径不能错！！！***最后，添加 C:\opencv\build\install\x64\mingw\bin 到你的系统路径。***

#### 验证安装
将当前目录更改为GoCV仓库的位置：
切换到你设置的GOPATH\src\gocv.io\x\gocv    下

    chdir %GOPATH%\src\gocv.io\x\gocv    
现在您应该能够构建或运行任何命令示例：

    go run cmd\version\main.go
版本程序应输出以下内容：

    gocv version: 0.17.0
    opencv lib version: 3.4.3
就是这样，现在你已经准备好使用GoCV了。

完美
![image.png](https://upload-images.jianshu.io/upload_images/5367714-90bceb34534b1ee1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


>百度云连接，正在上传中











