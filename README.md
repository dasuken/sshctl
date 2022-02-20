# sshctl
`sshctl` is a interactive cli for editing ssh config.

## How to install

```bash
$ go get gihub.com/dasuken/sshgen
```

**manually:**

Download binary from release page

## Usage
### add
from command
```bash
$ sshctl add
? select input format:  0: input by commandline (ex: ssh -i xxx.pem ec2-user@xxx)
? Enter ssh command:  ssh -i ~/key.pem ec2-user@11.178.212.233
? Enter unique label name as host:  host1

[OK] config was created!! if you use that config setting, $ ssh host1 

*if you want to use more options, please write directory. 
```

or, use interactive option

```bash
$ sshctl  add
 ? select input format:  1: input by interactive
 ? Enter Unique Name to represent specific endpoint: host2
 ? Enter HostName that represents the connection destination: 11.178.212.233
 ? Enter user name:  ec2-user
 ? Enter IdentityFile path:  /key.pem
 
[OK] config was created!! if you use that config setting, $ ssh host 

*if you want to use more options, please write directory. 
```

then, new settings have benn added to the `$HOME/.ssh/config`
```bash

HOST host1
	HostName 11.178.212.233
	User ec2-user
	identityFile /key.pem

HOST host2
	HostName 11.178.212.233
	User ec2-user
	identityFile /key.pem
``` 

### list

```bash
$ go run cmd/main.go list
0: host1 
1: host2 
```

## Roadmap
- [ ] Add Test
- [ ] UPDATE
- [ ] DELETE

# LICENCE

MIT licence