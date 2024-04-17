# Password Manager

Password Manager is a simple password management application built using Golang for the backend and SQLite3 for the database. It employs AES encryption to secure passwords stored in the database. While AES encryption enhances security, please note that this password manager is primarily a proof of concept and may not provide the highest level of security for managing passwords.

## Usage

### Installation

Clone the repository:

```bash
git clone https://github.com/hussainalboori/PasswordManager
```

### Running the Application

Navigate to the project directory and run the application:

```bash
cd PasswordManager
go run .
```

If the database file already exists, the application will use it; otherwise, it will create a new empty database file.

```bash
Database file 'data.db' already exists.
2024/04/17 11:21:08 Connect to our website through http://localhost:<port>
```
you can change the port by changing the variable named port in main.go

## Features

- **User Authentication:** Users can sign up, log in, and log out.
- **Password Management:** Users can add, view, update, and delete passwords.
- **Password Encryption:** Passwords are encrypted using AES encryption before being stored in the database.
- **Random Password Generation:** Users can generate random strong passwords if desired.
- **Simple Interface:** The application offers a user-friendly interface for easy navigation.

## Security Considerations

While AES encryption enhances security, it's essential to consider the following security measures:

- **Encryption Key Management:** Securely store and manage encryption keys to prevent unauthorized access to passwords.
- **Secure Password Practices:** Encourage users to create strong, unique passwords and enable multi-factor authentication where possible.
- **Regular Maintenance:** Keep the application updated with security patches and adhere to best practices to mitigate potential vulnerabilities.

## Disclaimer

This password manager is intended for educational and demonstration purposes only. It's not recommended for storing sensitive or critical passwords without undergoing thorough security evaluation and implementing additional security measures.

## License

This project is licensed under the [MIT License](LICENSE). Feel free to modify and use it according to your needs.