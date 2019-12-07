

create table action (
  id int not null,
  name varchar(30) not null
);


create table step (
  id int not null,
  title varchar(300) not null,
  priority int not null,
  actions varchar(300) not null,
  templates varchar(300) not null,
  next_steps varchar(300) not null,
  description varchar(300) not null,
  render_type varchar(300) not null,
  rule varchar(300) not null
);


create table template (
  id int not null,
  name varchar(30) not null,
  template varchar(3000) not null,
  template_type varchar(30) not null
);



