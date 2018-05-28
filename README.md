 <div align="center">
     <h1>~ passgen ~</h1>
     <strong>Simple command line tool to generate random passwords</strong><br><br>
 </div>

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
