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
