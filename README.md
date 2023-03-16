# Golang Email DNS Verifier
Golang Email DNS Verifier is a simple command-line tool that can be used to verify email addresses and check the DNS records of the email domain.
It is built using the Golang programming language and relies on third-party packages to perform email and DNS validation.

# Features
Verify the syntax of an email address.
Check the DNS records of the email domain to ensure it exists.
Check the MX records of the email domain to ensure it can receive emails.
Check if the email address is a disposable email address.
Check if the email domain is a known temporary email domain.

# Installation
Clone this repository to your local machine:

```bash
git clone https://github.com/santhoshsivanva/Golang-Email-DNS-Verifier.git
```

# Install the required dependencies:

```go
cd Golang-Email-DNS-Verifier
go mod download
```

# Build the application:

```go
go build main.go
```

# Run the application:

```bash
./main
```

# Usage
To use Golang Email DNS Verifier, simply run the command-line tool and enter an email address to verify. The application will perform various checks on the email address and provide feedback on its validity.

```yaml
$ ./main
Enter an email address to verify: john.doe@example.com
Validating email address: john.doe@example.com
Syntax: valid
DNS: valid
MX: valid
Disposable: no
Temporary: no
```

# Credits
This project was developed by santhoshsivanva as part of a personal project to learn Golang and email verification techniques.


# License
This project is licensed under the MIT License. See the LICENSE file for details.
