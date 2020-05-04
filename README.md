# Energy Comparison

## Executable

The executable for the program is in the folder

```
./bin/energy-comparison
```

In order to execute the program, please load the `plans.json` file first with the following command

```
./bin/energy-comparison ./files/plans.json
```

## Commands

Once the program has loaded the plans, it will prompt the following output

```
Enter command:
```

You can then provide the commands and process them by pressing the key ENTER

```
price 1000
```

The output should be simliar to the following

```
2020/03/24 19:37:42 eon,variable,108.68
2020/03/24 19:37:42 edf,fixed,111.25
2020/03/24 19:37:42 ovo,standard,120.23
2020/03/24 19:37:42 bg,standing-charge,121.33
```

After that, the program will ask again for another command with the prompt


## Compiling

If the executable is not working, you might need to compile the source code.

The program was developed with GoLang v1.13.

Follow the instructions of how to install GoLang [here](https://golang.org/dl/).

Once is installed, copy the TAR file with the source code inside your GOPATH (it is usually in your home directory).  Replace the value of GOPATH with the right path for your environment.

```
cp uswitch.com.tar $GOPATH/src/
```

The go to that folder and extract the TAR file

```
cd $GOPATH/src
tar -xvf uswitch.com.tar
```

Once the file is expanded, go inside the folder `uswitch.com/energy-comparison`. This is the folder with the source code

```
cd uswitch.com/energy-comparison
```

Once in the root directory, execute the following command

```
go build -o ./bin/energy-comparison
```

The previous command will compile the source code and put the binary file in the specified folder `./bin`. You could use the binary file with the instructions given above.
