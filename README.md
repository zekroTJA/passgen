 <div align="center">
     <h1>~ passgen ~</h1>
     <strong>Simple command line tool to generate random passwords</strong><br><br>
 </div>

 ---
 
 # Description
 
 This is a small binary to create random passwords in the console. The charset is the default ASCII-Charset from 0x20 (`Space`) to 0x7E (`~`). This tool only returns the raw password string as output, so you can use it for example like following:
```
$ echo "maxmustermann   $(passgen -l 64 -s 2)" >> userslist
```
 ---
 
 # Usage
 
 Help: 
 ```
 $ passgen --help
 ```
 
 Argument | Example | Description
 ---------|---------|-------------
 -l       | -l 64   | Character length of the password
 -n       | -n 10   | Number of generated passwords
 -s       | -s 3    | Strength of the password:<br>0 - `[a-z\d]`<br>1 - `[a-zA-Z\d]`<br>2 - `[\w\d]`<br>3 - `[\w\d!"§$%&/()=?-]`<br>4 - `.`
 -rx      | -rx [A-Za-z0-9] | Use regex to match characters of default charset to create password

---

# Build it for yourself

Just clone the repository with
```
$ git clone https://github.com/zekroTJA/passgen.git
```
or download the repository as zip.

You need to install the Go envoirement:
Download the installer [here](https://golang.org/dl/) or if you are using Linux with APT package manager:
```
# apt install -y golang-go
```

Then, build the repository with:
```
go build -o passgen passgen/src/github.com/zekrotja/passgen/main.go
```

Now, make the binary accessable in console:

For Linux:
```
# cp passgen /usr/bin/passgen
```

For Windows:
System Settings → System → Advanced System Settings → System Variables → Envoirement Variables → Add path of build to `PATH` variable
