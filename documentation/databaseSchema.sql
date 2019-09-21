
create table users(
	uid bigint PRIMARY KEY,
  	name text not null,
  	password_hash text not null,
  	role integer not null,
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
	home_content bigint,
  	FOREIGN KEY (writenBy) REFERENCES users(uid)
);


create table pins(
	uid bigint PRIMARY KEY,
	owner bigint not null,
  	longitude float NOT NUll,
  	latitude float NOT NUll,
  	description text,
	time bigint not null,
	tag_type integer not null,
	name text not null,
	color varchar(7) not null,
	foreign key (owner) references users(uid)
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


insert into users
values(0,'admin','21232f297a57a5a743894a0e4a801fc3',2);
