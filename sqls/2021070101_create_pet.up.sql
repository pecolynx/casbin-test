create table `pet` (
 `id` int auto_increment
,`version` int not null default 1
,`created_at` datetime not null default current_timestamp
,`updated_at` datetime not null default current_timestamp on update current_timestamp
,`name` varchar(20) character set ascii not null
,primary key(`id`)
,unique(`name`)
);
