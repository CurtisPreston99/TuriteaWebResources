
create table users(
	id bigint PRIMARY KEY,
  	name text not null,
  	password text not null,
  	role integer not null，
	unique(name)
);

create table media (
	uid bigint PRIMARY KEY,
	title text,
	url text not null,
	type integer
);

create table articles(
	id bigint PRIMARY KEY,
  	summary text,
  	writenBy int,
  	FOREIGN KEY (writenBy) REFERENCES users(id)
);


create table pins(
	uid bigint PRIMARY KEY,
	owner bigint not null,
  	longitude float NOT NUll,
  	latitude float NOT NUll,
  	discription text,
	time bigint not null,
	tag_type integer not null,
	name text not null,
	foreign key (owner) references users(id)
);


create table subscription(
  	name text,
  	email text primary key
);


create table feedback (
	id SERIAL PRIMARY KEY,
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
