# How to contribute

Thank you for your interest in improving GoCV.

We would like your help to make this project better, so we appreciate any contributions. See if one of the following descriptions matches your situation:

### Newcomer to GoCV, to OpenCV, or to computer vision in general

We'd love to get your feedback on getting started with GoCV. Run into any difficulty, confusion, or anything else? You are not alone. We want to know about your experience, so we can help the next people. Please open a Github issue with your questions, or get in touch directly with us.

### Something in GoCV is not working as you expect

Please open a Github issue with your problem, and we will be happy to assist.

### Something you want/need from OpenCV does not appear to be in GoCV

We probably have not implemented it yet. Please take a look at our [ROADMAP.md](ROADMAP.md). Your pull request adding the functionality to GoCV would be greatly appreciated.

### You found some Python code on the Internet that performs some computer vision task, and you want to do it using GoCV

Please open a Github issue with your needs, and we can see what we can do.

## How to use our Github repository

The `master` branch of this repo will always have the latest released version of GoCV. All of the active development work for the next release will take place in the `dev` branch. GoCV will use semantic versioning and will create a tag/release for each release.

Here is how to contribute back some code or documentation:

- Fork repo
- Create a feature branch off of the `dev` branch
- Make some useful change
- Submit a pull request against the `dev` branch.
- Be kind

## How to add a function from OpenCV to GoCV

Here are a few basic guidelines on how to add a function from OpenCV to GoCV:

- Please open a Github issue. We want to help, and also make sure that there is no duplications of efforts. Sometimes what you need is already being worked on by someone else.
- Use the proper Go style naming `MissingFunction()` for the Go wrapper.
- Make any output parameters `Mat*` to indicate to developers that the underlying OpenCV data will be changed by the function.
- Use Go types when possible as parameters for example `image.Point` and then convert to the appropriate OpenCV struct. Also define a new type based on `int` and `const` values instead of just passing "magic numbers" as params. For example, the `VideoCaptureProperties` type used in `videoio.go`.
- Always add the function to the GoCV file named the same as the OpenCV module to which the function belongs.
- If the new function is in a module that is not yet implemented by GoCV, a new set of files for that module will need to be added.
- Always add a "smoke" test for the new function being added. We are not testing OpenCV itself, but just the GoCV wrapper, so all that is needed generally is just exercising the new function.
- If OpenCV has any default params for a function, we have been implementing 2 versions of the function since Go does not support overloading. For example, with a OpenCV function:

```c
opencv::xYZ(int p1, int p2, int p3=2, int p4=3);
```

We would define 2 functions in GoCV:

```go
// uses default param values
XYZ(p1, p2)

// sets each param
XYZWithParams(p2, p2, p3, p4)
```

## How to run tests

To run the tests:

```
go test .
go test ./contrib/.
```

If you want to run an individual test, you can provide a RegExp to the `-run` argument:
```
go test -run TestMat
```

If you are using Intel OpenVINO, you can run those tests using:

```
go test ./openvino/...
```

## Contributing workflow

This section provides a short description of one of many possible workflows you can follow to contribute to `CoCV`. This workflow is based on multiple [git remotes](https://git-scm.com/docs/git-remote) and it's by no means the only workflow you can use to contribute to `GoCV`. However, it's an option that might help you get started quickly without too much hassle as this workflow lets you work off the `gocv` repo directory path!

Assuming you have already forked the `gocv` repo, you need to add a new `git remote` which will point to your GitHub fork. Notice below that you **must** `cd` to `gocv` repo directory before you add the new `git remote`:

```shell
cd $GOPATH/src/gocv.io/x/gocv
git remote add gocv-fork https://github.com/YOUR_GH_HANDLE/gocv.git
```

Note, that in the command above we called our new `git remote`, **gocv-fork** for convenience so we can easily recognize it. You are free to choose any remote name of your liking.

You should now see your new `git remote` when running the command below:

```shell
git remote -v

gocv-fork	https://github.com/YOUR_GH_HANDLE/gocv.git (fetch)
gocv-fork	https://github.com/YOUR_GH_HANDLE/gocv.git (push)
origin	        https://github.com/hybridgroup/gocv (fetch)
origin	        https://github.com/hybridgroup/gocv (push)
```

Before you create a new branch from `dev` you should fetch the latests commits from the `dev` branch:

```shell
git fetch origin dev
```

You want the `dev` branch in your `gocv` fork to be in sync with the `dev` branch of `gocv`, so push the earlier fetched commits to your GitHub fork as shown below. Note, the `-f` force switch might not be needed:

```shell
git push gocv-fork dev -f
```

Create a new feature branch from `dev`:

```shell
git checkout -b new-feature
```

After you've made your changes you can run the tests using the `make` command listed below. Note, you're still working off the `gocv` project root directory, hence running the command below does not require complicated `$GOPATH` rewrites or whatnot:

```shell
make test
```

Once the tests have passed, commit your new code to the `new-feature` branch and push it to your fork running the command below:

```shell
git push gocv-fork new-feature
```

You can now open a new PR from `new-feature` branch in your forked repo against the `dev` branch of `gocv`.
