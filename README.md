# Cryptography_project
# Project Title: "Secure Password Manager using bcrypt and AES Encryption"

1.	Overview of The Project Goal
The goal of this project is to design and implement a secure password manager that allows users to safely store and retrieve their online account credentials. The system uses a master password for authentication and applied modern cryptographic techniques to ensure that all sored passwords remain encrypted , protected , and inaccessible to unauthorized users.
The project demonstrates how secure systems project confidential information such as: Website logins , Application credentials , Personal passwords. It is look like or mimics the core security design of real password managers like LastPass, 1Password…, but in a simplified, educational form.
2.	Problem Statement
In modern digital life , a single user may have a lot of online accounts, including :
•	Email
•	Social media
•	Banking apps
•	Any apps…
Because humans cannot remember many unique, strong passwords, they often:
•	Reuse the same weak password across different services
•	Store passwords in unsafe locations (notes apps, text files ,…)
•	Forget login credentials
•	Get exposed to credential-suffering attacks
This behavior leads to serious cybersecurity risk such as:
•	Account haked
•	Data theft
•	Unauthorized access
Therefore, users need a secure way to store their passwords that does not expose them to attackers.
3.	Solution Overview
This project solve =s the problem by creating a locally encrypted password vault , protected using :
•	Bcrypt hashing :
Used to securely store the master password. Even if the vault file is stolen, attackers cannot recover the master password.
•	PBKDF2 (Password-Based key Derivation Function):
Used to convert the master password into a strong, 256-bit AES key through key stretching.
•	AES Encryption (AES-GCM or AES-CBC): 
Used to encrypt each saved password entry. Only the correct master password can decrypt the data.
•	JSON vault storage: 
All encrypted entries are stored in an easy-to-read but secure JSON file.
This ensures that:
•	Passwords are never stored in plaintext
•	Even the developer cannot see user passwords
•	Attackers cannot decrypt the vault without the correct master password
4.	Motivation
This project is motivated by : 
•	Real world cybersecurity needs:
Password managers are a critical part of modern security. Understanding their design helps improve such as Authentication mechanisms, Key management, Encryption practices.
•	Hand on cryptography learning :
Instead of study on Cryptography subject , this project will show the implementation of what I learned like about how encryption protects real data, how hashing prevents password theft.
•	Practical usefulness:
The project has direct value such as : Users can securely store credentials, Student learn how encrypted vaults work.
5.	Related Cryptography Concepts
This project integrated multiple cryptographic mechanisms:
•	Bcrypt : Slow, adaptive password hashing algorithm.
•	Salt: Random value added before hashing.
•	PBKDF2: Converts a password into a strong AES key.
•	IV/Nonce: Ensures identical passwords do not produce identical ciphertexts.
•	JSON Storage: Stores encrypted entries.
