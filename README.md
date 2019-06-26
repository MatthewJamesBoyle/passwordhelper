#Password Helper
##Project Goal
Sick of your bank/credit card provider/other site asking for character 1,3,5,7,432 of your password? If you use a password manager, it can be really hard to distinguish which character is which and if (like me) you have found yourself locked out of your account, then password helper is for you!

##How it Works
Once installed, run `passwordhelper ${password} character1 character 2 character 3` 
in your terminal. It will output a comma separated list of selected characters. 

Example
`passwordhelper mypassword0000 1 4 6`
`> m,a,s`

##Installation 
###From Releases
Download the latest release binary, rename it to `passwordhelper` move it to `/usr/local/bin` (or anyone else on your path).
You should now be able to run passwordhelper.

####From Source
clone this repo and run `go build -o passwordhelper cmd/cli/main` move the built binary to `usr/local/bin`