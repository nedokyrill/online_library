create table if not exists "songs"
(
    "id"           serial primary key,
    "group_name"   varchar(255) not null,
    "song_name"    varchar(255) not null,
    "release_date" date not null default current_date,
    "text"         text         not null,
    "link"         text         not null
);

create index if not exists idx_songs_song_name on "songs" ("song_name");
