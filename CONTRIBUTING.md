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

You will need to run `source env.sh` or `env.cmd` depending on your OS one time before you can run the Go unit tests.

Then you can run:

```
go test .
go test ./contrib/.
```

If you are using the Intel CV SDK, you can run those tests using:

```
go test ./pvl/.
```
