# Laptop Coms Server 📡
This is a personal server written for my MacBook to automate tasks and run commands securely.

## Why not SSH?
Security, SSH allows your computer to be accessed remotely and if compromised, allows access to your entire computer undermining all the existing security measures I have implemented. Jailed SSH didn't really seem to be viable on MacOS.

## Why a Web Server?
Although there might be ore light weight alternatives, other than SSH, a web server seemed to be the most compatible especially with iOS Shortcuts.

## Why is it written in Go?
Performance. I wanted this to take the least amount of system resources without compromising security and Go was a perfect lightweight solution. This is also why I'm planning on migrating from TCP to UDP. QUIC is cool too, but it might take more resources due to encryption not to mention the lack of compatibility.

## Why no HTTPS?
Although in theory it should be more secure, it takes more system resources due to encryption, not to mention the additional latency on the initial handshake. For my use case, HTTPS does not really provide much benefit.

## Implementations 📻

### Phone Backup
The idea of this method is to backup my phone every ~24 hours to my Mac. Backing up any iPhone via iCloud is basically pointless because Apple only gives you 5 GB of free storage. Almost every day, I connect my phone to my docking station which is connected to my MacBook. Keeping this in mind, I have set an automation via the Shortcuts app to ping the server ever time it starts charging to backup. There is additional logic to prevent redundant pinging and logic on the server side to check the time and validate integrity.

## Running
IDK why you would want to run this but if you really do make sure you have Go Lang installed and this repo downloaded run `go run main.go` in the projects directory

## Installing
Again, IDK why you would want to install this but if you really do run `go install sethusenthil.com/main/coms`