create table wagers (
    id int AUTO_INCREMENT primary key,
    total_wager_value int not null ,
    odds int not null ,
    selling_percentage float not null,
    selling_price float(10,2) not null,
    current_selling_price float(10,2) not null,
    percentage_sold int,
    amount_sold int,
    placed_at timestamp,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
)

create table purchases (
    id int AUTO_INCREMENT primary key,
    wager_id int not null,
    buying_price float(10,2),
    bought_at timestamp,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
)

atler table purchases add foreign key (wager_id) references wagers(id);