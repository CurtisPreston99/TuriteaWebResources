
create table users(
	id bigint PRIMARY KEY,
  	user_name text not null,
  	password text not null,
  	role integer not null
);

create table media (
	uid bigint PRIMARY KEY,
	title text,
	url text not null,
	type integer,
);

create table articles(
	id bigint PRIMARY KEY,
  	summary text,
  	writenBy int,
  	FOREIGN KEY (writenBy) REFERENCES users(id)
);


create table pins(
	uid bigint PRIMARY KEY,
  	longitude float NOT NUll,
  	latitude float NOT NUll,
  	discription text,
	time bigint not null
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
  	feedback text,
	state boolean not null
);

create table pinlinkarticle(
	pin_id bigint references pins(uid),
	article_id bigint references articles(id),
	primary key (pin_id, article_id)
);

