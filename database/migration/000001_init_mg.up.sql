-- Users Creation
create table users (
    id int not null auto_increment,
    username varchar(30) unique,
    password text,
    role text,
    primary key(id)
);

INSERT INTO users (username, password, role) VALUES
(
  'librarian',
  SHA1('password'),
  'admin'
),
(
  'bookworm123',
  SHA1('reader123'),
  'user'
),
(
  'techfanatic',
  SHA1('cybersecurity'),
  'user'
);

-- Books Creation
create table books (
    bookid int auto_increment, 
    name text not null, 
    author text, 
    shelf text, 
    primary key(bookid)
);

-- Issued Craetion
create table issued (
    issueid int not null auto_increment, 
    bookid int unique, 
    id int, 
    primary key (issueid), 
    foreign key (bookid) references books(bookid), 
    foreign key (id) references users(id)
);

INSERT INTO books (name, author, shelf) VALUES
("To Kill a Mockingbird", "Harper Lee", "A1"),
("The Lord of the Rings", "J. R. R. Tolkien", "B3"),
("The Hitchhiker's Guide to the Galaxy", "Douglas Adams", "C2"),
("Pride and Prejudice", "Jane Austen", "D5"),
("One Hundred Years of Solitude", "Gabriel García Márquez", "E7"),
("Frankenstein", "Mary Shelley", "F1"),
("The Great Gatsby", "F. Scott Fitzgerald", "G4"),
("The Catcher in the Rye", "J. D. Salinger", "H6"),
("1984", "George Orwell", "I3"),
("The Adventures of Huckleberry Finn", "Mark Twain", "J8"),
("Hacking: The Art of Exploitation", "Jon Erickson", "K2"),
("Cryptography Engineering: Design Principles and Practical Applications", "Niels Ferguson", "L1"),
("Ghost in the Wires: My Adventures as the World\'s Most Wanted Hacker", "Kevin Mitnick", "M4"),
("The Hacker Playbook 2: Practical Guide To Penetration Testing", "Peter Kim", "N7"),
("Web Application Security: A Developer's Guide", "David Heffelfinger", "O9");

INSERT INTO issued (bookid, id) VALUES
  -- User 'librarian' (assumed ID 1) issued book with ID 2 (Fiction1)
  (2, 1),
  (10, 1),
  -- User 'bookworm123' (assumed ID 2) issued book with ID 5 (Classics1)
  (5, 2),
  (6, 2),
  (3, 2),
  -- User 'techfanatic' (assumed ID 3) issued book with ID 11 (Security1)
  (11, 3),
  (15, 3);
