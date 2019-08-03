
create table users(
	id int PRIMARY KEY,
  	user_name text not null,
  	password text not null,
  	permissions text not null
);

create table videos(
	id int PRIMARY KEY,
  	name text,
  	uploadedBy int,
  	url text,
  	FOREIGN KEY (uploadedBy) REFERENCES users(id)
);

create table images(
	id int PRIMARY KEY,
  	name text,
  	uploadedBy int,
  	url text,
  	FOREIGN KEY (uploadedBy) REFERENCES users(id)
);

create table articles(
	id int PRIMARY KEY,
  	summary text,
  	writenBy int,
  	FOREIGN KEY (writenBy) REFERENCES users(id)
);


create table pins(
	id int PRIMARY KEY,
  	long float NOT NUll,
  	lat float NOT NUll,
  	discription text,
  	article int,
    FOREIGN KEY (article) REFERENCES articles(id)
);


create table subscription(
	id int PRIMARY KEY,
  	name text,
  	email text
);


create table feedback(
	id int PRIMARY KEY,
  	name text,
  	email text,
  	feedback text
);



