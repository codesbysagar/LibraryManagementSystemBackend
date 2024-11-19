
# 📔LibraryManagementSystemBackend

A comprehensive **Library Management System** built in **Go**, designed to manage library operations efficiently. The system integrates with **MongoDB** for database management and offers functionalities like:

- **Member Management**: Add, update, and manage member details.
- **Book Inventory Management**: Track books, including multiple copies and availability.
- **Book Issuance & Returns**: Maintain records of issued books and due dates.
- **Search & Filter**: Query books and members using various criteria.

## Features
- RESTful API implementation for seamless client-server communication.
- Robust error handling and validation to ensure system reliability.
- Secure member data storage with password hashing.


## Tech Stack

- **Go**
- **MongoDB**
- **Postman** (API testing)
- **Git** (version control)

## API Reference
**Endpoint**
- `"POST /CreateMember"`: **Adds new member in database**
- `"POST /AddNewBook"`: **Adds new book in database**
- `"POST /IssueBook"`: **Issues the book**
- `"POST /ReturnBook"`: **Returns the book**
- `"POST /GetAllRecordMember"`: **Returns all the books Issued by a member**
- `"POST /BookIssueRecord"`: **Returns all the books Issued to all the member with their details**
### 1. **Create Member**
- **URL**: `/CreateMember`
- **Method**: `POST`
- **Description**: Registers a new library member.
- **Request Body**: example
    ```json
    {
      "fullname": "John Doe",
      "contact": 1234567890,
      "email": "john.doe@example.com",
      "password": "securepassword"
    }
    ```
- **Response**:
    ```json
    {
      "status": "OK",
      "message": "Member Added Successfully",
      "data": {
        "member_id": "unique-member-id"
      }
    }
    ```

---

### 2. **Add New Book**
- **URL**: `/AddNewBook`
- **Method**: `POST`
- **Description**: Adds a new book to the inventory.
- **Request Body**: example
    ```json
    {
      "title": "The Great Gatsby",
      "author": "F. Scott Fitzgerald",
      "quantity": 5,
      "genre": "Fiction"
    }
    ```
- **Response**:
    ```json
    {
      "status": "OK",
      "message": "Book Added Successfully",
      "data": {
        "book_id": "unique-book-id"
      }
    }
    ```

---

### 3. **Issue Book**
- **URL**: `/IssueBook`
- **Method**: `POST`
- **Description**: Issues a book to a member.
- **Request Body**: example
    ```json
    {
      "memberId": 101,
      "password": "securepassword",
      "bookId": 202
    }
    ```
- **Response**:
    ```json
    {
      "status": "OK",
      "message": "Book Issued Successfully",
      "data": {
        "record_id": 1
      }
    }
    ```

---

### 4. **Return Book**
- **URL**: `/ReturnBook`
- **Method**: `POST`
- **Description**: Returns a borrowed book.
- **Request Body**: example
    ```json
    {
      "memberId": 101111,
      "password": "securepassword",
      "recordId": 1234567
    }
    ```
- **Response**:
    ```json
    {
      "status": "OK",
      "message": "Book Returned Successfully",
      "data": {
        "key": "bookData"
      }
    }
    ```

---

### 5. **Get All Borrowed Records of a Member**
- **URL**: `/GetAllRecordMember`
- **Method**: `POST`
- **Description**: Retrieves all borrowed books for a specific member.
- **Request Body**: example
    ```json
    {
      "memberId": 101111
    }
    ```
- **Response**:
    ```json
    {
      "status": "OK",
      "message": "Got All Borrowed Books",
      "data": [
        {
          "bookId": 202,
          "title": "The Great Gatsby",
          "author": "F. Scott Fitzgerald",
          "genre": "Fiction",
          "issueDate": "2024-11-17",
          "dueDate": "2024-12-01"
        }
      ]
    }
    ```

---

### 6. **Book Issue Record**
- **URL**: `/BookIssueRecord`
- **Method**: `POST`
- **Description**: Retrieves a record of all members who borrowed a specific book.
- **Request Body**: example
    ```json
    {
      "bookId": 20222
    }
    ```
- **Response**:
    ```json
    {
      "status": "OK",
      "message": "Got All Issued Records",
      "data": [
        {
          "recordId": 1234567,
          "memberId": 101111,
          "fullname": "John Doe",
          "contact": 1234567890,
          "email": "john.doe@example.com",
          "issueDate": "2024-11-17",
          "dueDate": "2024-12-01"
        }
      ]
    }
    ```


## Screenshots

<img src="https://github.com/codesbysagar/LibraryManagementSystemBackend/blob/main/Diagrams/LMS-BookEntry.png" alt="LMS-BookEntry" width="300" height="424"/><img src="https://github.com/codesbysagar/LibraryManagementSystemBackend/blob/main/Diagrams/LMS-CreateNewMember.png" alt="LMS-CreateNewMember" width="300" height="424"/><img src="https://github.com/codesbysagar/LibraryManagementSystemBackend/blob/main/Diagrams/LMS-Database%20structure.png" alt="LMS-Database%20structure" width="300" height="424"/><img src="https://github.com/codesbysagar/LibraryManagementSystemBackend/blob/main/Diagrams/LMS-IssuaBook.png" alt="LMS-IssuaBook" width="300" height="424"/><img src="https://github.com/codesbysagar/LibraryManagementSystemBackend/blob/main/Diagrams/LMS-ReturnBook.png" alt="LMS-ReturnBook" width="300" height="424"/>



## License
MIT License

Copyright (c) 2024 Sagar Sharma

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.