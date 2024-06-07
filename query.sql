create table genres (
	id serial primary key,
	genre_name varchar(255) not null
);

create table admins (
	id serial primary key,
	name varchar(255) not null,
	email varchar(255) not null,
	password varchar(255) not null,
	phone varchar(20)
);

create table books (
	id serial primary key,
	genre_id int not null,
	title varchar(255) not null,
	author varchar(255) not null,
	publisher varchar(255) not null,
	publish_year int not null,
	foreign key (genre_id) references genres(id)
);

create table users (
	id serial primary key,
	name varchar(255) not null,
	email varchar(255) not null,
	password varchar(255) not null,
	address text,
	phone varchar (20) not null,
	birth_date date
);

create table loans (
	id serial primary key,
	user_id int not null,
	book_id int not null,
	deadline_date timestamp,
	date_loan timestamp,
	date_return  timestamp,
	status_loan varchar(255),
	foreign key (user_id) references users(id),
	foreign key (book_id) references books(id)
);

create table book_request (
	id serial primary key,
	user_id int not null,
	approved_admin_id int not null,
	title varchar(255) not null,
	author varchar(255) not null,
	publisher varchar(255) not null,
	publish_year int not null,
	status_request varchar(255) not null,
	foreign key (user_id) references users(id),
	foreign key (approved_admin_id) references admins(id)
);


--menambahkan 5 data genre
INSERT INTO genres (genre_name) VALUES ('Science Fiction');
INSERT INTO genres (genre_name) VALUES ('Fantasy');
INSERT INTO genres (genre_name) VALUES ('Mystery');
INSERT INTO genres (genre_name) VALUES ('Romance');
INSERT INTO genres (genre_name) VALUES ('Non-Fiction');

--menambahkan 5 data buku untuk masing masing genre
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (1, 'Dune', 'Frank Herbert', 'Chilton Books', 1965);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (1, 'Neuromancer', 'William Gibson', 'Ace', 1984);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (1, 'Foundation', 'Isaac Asimov', 'Gnome Press', 1951);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (1, 'Snow Crash', 'Neal Stephenson', 'Bantam Books', 1992);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (1, 'Hyperion', 'Dan Simmons', 'Doubleday', 1989);

INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (2, 'The Hobbit', 'J.R.R. Tolkien', 'George Allen & Unwin', 1937);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (2, 'Harry Potter and the Philosopher Stone', 'J.K. Rowling', 'Bloomsbury', 1997);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (2, 'A Game of Thrones', 'George R.R. Martin', 'Bantam Books', 1996);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (2, 'The Name of the Wind', 'Patrick Rothfuss', 'DAW Books', 2007);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (2, 'The Way of Kings', 'Brandon Sanderson', 'Tor Books', 2010);

INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (3, 'The Girl with the Dragon Tattoo', 'Stieg Larsson', 'Norstedts Förlag', 2005);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (3, 'Gone Girl', 'Gillian Flynn', 'Crown Publishing Group', 2012);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (3, 'The Da Vinci Code', 'Dan Brown', 'Doubleday', 2003);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (3, 'In the Woods', 'Tana French', 'Viking Press', 2007);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (3, 'Big Little Lies', 'Liane Moriarty', 'Penguin Publishing Group', 2014);

INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (4, 'Pride and Prejudice', 'Jane Austen', 'T. Egerton, Whitehall', 1813);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (4, 'The Notebook', 'Nicholas Sparks', 'Warner Books', 1996);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (4, 'Me Before You', 'Jojo Moyes', 'Michael Joseph', 2012);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (4, 'Twilight', 'Stephenie Meyer', 'Little, Brown and Company', 2005);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (4, 'The Time Travelers Wife', 'Audrey Niffenegger', 'MacAdam/Cage', 2003);

INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (5, 'Sapiens: A Brief History of Humankind', 'Yuval Noah Harari', 'Harvill Secker', 2011);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (5, 'Educated', 'Tara Westover', 'Random House', 2018);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (5, 'Becoming', 'Michelle Obama', 'Crown Publishing Group', 2018);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (5, 'The Immortal Life of Henrietta Lacks', 'Rebecca Skloot', 'Crown Publishing Group', 2010);
INSERT INTO books (genre_id, title, author, publisher, publish_year) VALUES (5, 'The Wright Brothers', 'David McCullough', 'Simon & Schuster', 2015);

--menambahkan 5 user
INSERT INTO users (name, email, password, address, phone, birth_date) VALUES ('Alice', 'alice@example.com', 'password123', 'Bandung', '081265993766', '1990-01-01');
INSERT INTO users (name, email, password, address, phone, birth_date) VALUES ('Farah', 'farah@example.com', 'password124', 'Banda Aceh', '081265993777', '1985-02-14');
INSERT INTO users (name, email, password, address, phone, birth_date) VALUES ('Alza', 'alza@example.com', 'password125', 'Lampung', '085260652246', '1978-03-21');
INSERT INTO users (name, email, password, address, phone, birth_date) VALUES ('David', 'david@example.com', 'password126', 'Jakarta Selatan', '085358097560', '1992-04-18');
INSERT INTO users (name, email, password, address, phone, birth_date) VALUES ('Yasa Singgih', 'yasa@example.com', 'password127', 'Bogor', '081275798064', '1988-05-12');

--Tampilkan hasil pencarian data buku berdasarkan judul buku
select * from books b where title = 'Becoming';

--Tampilkan hasil pencarian data buku berdasarkan genre id
select * from books b where b.genre_id = 2;

-- Menambahkan data peminjaman buku
-- Data book_id dimulai dari 7
INSERT INTO loans (user_id, book_id, deadline_date, date_loan, status_loan)
VALUES (1, 7, '2024-06-21 00:00:00', '2024-06-07 00:00:00', 'borrowed');

INSERT INTO loans (user_id, book_id, deadline_date, date_loan, status_loan)
VALUES (1, 8, '2024-06-21 00:00:00', '2024-06-07 00:00:00', 'borrowed');

INSERT INTO loans (user_id, book_id, deadline_date, date_loan, status_loan)
VALUES (1, 13, '2024-06-21 00:00:00', '2024-06-07 00:00:00', 'borrowed');

INSERT INTO loans (user_id, book_id, deadline_date, date_loan, status_loan)
VALUES (2, 9, '2024-06-21 00:00:00', '2024-06-07 00:00:00', 'borrowed');

INSERT INTO loans (user_id, book_id, deadline_date, date_loan, status_loan)
VALUES (3, 7, '2024-06-21 00:00:00', '2024-06-07 00:00:00', 'borrowed');


--Mengubah status peminjaman buku yang dipinjam oleh user 1 dan buku dengan id 7 tadi menjadi “dikembalikan”
update loans 
set status_loan = 'returned', date_return = '2024-06-15 00:00:00'
where user_id = 1 and book_id = 7


-- Tampilkan data user beserta buku yang masih dipinjam/belum dikembalikan
select u.id  as user_id, u.name as user_name, b.id as book_id, b.title as book_title, l.id as loan_id, l.status_loan 
from loans l
join users u on l.user_id = u.id 
join books b on l.book_id = b.id
where l.status_loan = 'borrowed';

-- Tampilkan data buku yang statusnya telah dikembalikan oleh user.
select b.id as book_id, b.title as book_title, b.genre_id, b.author, b.publisher, b.publish_year, l.id as loan_id, l.status_loan 
from loans l
join books b on l.book_id = b.id
where l.status_loan = 'returned';

-- Tampilkan data buku yang belum pernah dipinjam oleh user
select  b.id as book_id, b.title, b.author, b.publisher, b.publish_year
from books b
left join loans l on b.id = l.book_id
where l.id is null;

-- Tampilkan data user beserta banyaknya buku yang pernah dia pinjam
select  u.id as user_id, u.name as user_name, u.email, u.phone, count(l.id) as total_books_borrowed
from users u
left join loans l ON u.id = l.user_id
group by u.id
order by u.id;

-- Tampilkan data user yang belum pernah meminjam buku
select u.id as user_id, u.name as user_name, u.email as user_email
from users u 
left join loans l on u.id = l.user_id 
where l.id is null;

-- Tampilkan data user yang sudah pernah meminjam buku
select u.id as user_id, u.name as user_name, u.email as user_email, count(l.id) as total_books_borrowed
from users u 
left join loans l on u.id = l.user_id 
where l.id is not null
group by u.id
order by u.id;

-- Tampilkan data user yang paling banyak meminjam buku
select u.id as user_id, u.name as user_name, u.email as user_email, count(l.id) as total_books_borrowed
from users u 
left join loans l on u.id = l.user_id 
where l.id is not null
group by u.id
order by total_books_borrowed desc 
limit 1;

-- Tampilkan data genre beserta banyaknya buku di masing-masing genre
SELECT genres.id AS genre_id, genres.genre_name, COUNT(books.id) AS total_books
FROM genres
LEFT JOIN books ON genres.id = books.genre_id
GROUP BY genres.id, genres.genre_name
ORDER BY total_books DESC;

-- Tampilkan genre yang paling banyak dipinjam bukunya oleh user
SELECT genres.id AS genre_id, genres.genre_name, COUNT(loans.id) AS total_loans
FROM genres
JOIN books ON genres.id = books.genre_id
JOIN loans ON books.id = loans.book_id
GROUP BY genres.id, genres.genre_name
ORDER BY total_loans DESC
LIMIT 1;

-- Tampilkan data user, beserta buku yang dipinjam dan sekaligus genre dari buku tersebut dalam satu query statement.
SELECT users.id AS user_id, users.name AS user_name, users.email AS user_email,
       books.id AS book_id, books.title AS book_title,
       genres.id AS genre_id, genres.genre_name AS genre_name
FROM users
JOIN loans ON users.id = loans.user_id
JOIN books ON loans.book_id = books.id
JOIN genres ON books.genre_id = genres.id;




