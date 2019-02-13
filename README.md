 <div align="center">
     <h1>~ passgen ~</h1>
     <strong>Simple command line tool to generate cryptographically random passwords</strong><br><br>
 </div>

 ---
 
 # Description
 
 This is a small binary to quickly create random passwords in the console. The unsed charset where the characters are picked from is the default ASCII-Charset from 0x20 (`Space`) to 0x7E (`~`). This tool only returns the raw password strings as output, so you can pipe them directly, for example, into a file:
```
$ echo "maxmustermann  $(passgen -l 64 -s 2)" | tee -a userslist
```
 ---
 
 # Usage
 
 Help: 
 ```
 $ passgen -h
 ```
 
```
Usage of ./passgen2:
  -l uint
        length in characters of the generated token (default 32)
  -n uint
        number of generated tokens (default 1)
  -rx string
        Regex for used characters for generating the token (matched on the whole charset string)
  -s uint
        strength of the generated token
          0 - [a-z\d]
          1 - [a-zA-Z\d]
          2 - [\w\d]
          3 - [\w\d!"§$%&/()=?-]
          4 - .
         (default 2)
  -sep string
        custom seperator for multiple tokens (default "\n")
  -v    display version information
```

---

# v2 vs. v1

Because passgen was one of my firts go projects, the code was kind of crap and. Because I am using this program myself very frequently, I had a lot of ideas how I could optimize this tool:

- For random number generation, now `crypto/rand` is used instead of `math/rand`  
  *This is quite slower, but way more "true" random than mathimatically randomly generated numbers.* 
- Flags are now parsed by golang standard flags `library`  
  *Which is way better and more handy than handling flags manually like before.*
- Asyncronous password generation  
  *The characters of the password are now generated asyncronously in seperate go routines, which is ~26 times faster than in v1.1, as you can see in my little test below:*  
  ```
  # version 1.1
  zekro@vmi211544:~$ time ./passgen11 -l 1000 -n 1000 &> /dev/null
  real    0m34.603s
  user    0m36.300s
  sys     0m0.832s

  # version 2.0
  zekro@vmi211544:~$ time ./passgen2 -l 1000 -n 1000 &> /dev/null
  real    0m1.325s
  user    0m2.396s
  sys     0m2.220s
  ```

---

# Get it

### Dowload

You can just download the binaries from [Releases](https://github.com/zekroTJA/passgen/releases) and install them on your system.

### Self-compile

Alternatively, you can compile the sources by yourself. Just clone the repository and chanche directory into the created directory:
```
$ git clone https://github.com/zekroTJA/passgen.git
$ cd passgen
```

Then you can build the binary with 
```
$ make
```

or install the binary directly with 
```
$ sudo make install
```
*This will only work on Linux or by using gitbash on Windows!*

The Makefile was only tested on Debian and Windows. If it fails, you can also build the binaries as following:

```
$ git clone ttps://github.com/zekroTJA/passgen.git gopath/src/github.com/zekroTJA/passgen
$ export GOPATH=$PWD/gopath
$ cd gopath/src/github.com/zekroTJA/passgen
$ go build -v -o passgen .
```

Now, you can put the binary where you want to access it from the terminal.

---

Copyright © 2019 zekro Development (Ringo Hoffmann).  
Covered by MIT Licence.