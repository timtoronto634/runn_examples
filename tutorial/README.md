https://zenn.dev/katzumi/books/runn-tutorial


# Logs

## use-db.yml
```
% runn run ./init-db.yml --verbose
=== スキーマを定義する (./init-db.yml)
    --- テーブル作成 (crateTable) ... ok
    --- テーブル一覧を確認する (showTables) ... ok


1 scenario, 0 skipped, 0 failures
```

sqlite3
```
% sqlite3
SQLite version 3.43.2 2023-10-10 13:08:14
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.

sqlite> .open tutorial.db
sqlite> select * from article;
sqlite> 

sqlite> .schema article
CREATE TABLE article (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    article_type TEXT,
    body_letters_count INTEGER,
    body_updated_at TEXT,
    comments_count INTEGER,
    emoji TEXT,
    is_suspending_private INTEGER,
    liked_count INTEGER,
    path TEXT,
    pinned INTEGER,
    post_type TEXT,
    publication TEXT,
    published_at TEXT,
    slug TEXT,
    source_repo_updated_at TEXT,
    title TEXT
);

sqlite> INSERT INTO article (
(x1...>     article_type,
(x1...>     body_letters_count,
(x1...>     body_updated_at,
(x1...>     comments_count,
(x1...>     emoji,
(x1...>     is_suspending_private,
(x1...>     liked_count,
(x1...>     path,
(x1...>     pinned,
(x1...>     post_type,
(x1...>     publication,
(x1...>     published_at,
(x1...>     slug,
(x1...>     source_repo_updated_at,
(x1...>     title
(x1...> ) VALUES (
(x1...>     'blog',
(x1...>     1000,
(x1...>     '2024-10-13 15:30:00',
(x1...>     5,
(x1...>     '😊',
(x1...>     0,
(x1...>     10,
(x1...>     '/blog/my-first-post',
(x1...>     0,
(x1...>     'article',
(x1...>     'My Publication',
(x1...>     '2024-10-13 12:00:00',
(x1...>     'my-first-post',
(x1...>     '2024-10-13 11:30:00',
(x1...>     'My First Blog Post'
(x1...> );
sqlite> 
sqlite> 
sqlite> select * from article;
1|blog|1000|2024-10-13 15:30:00|5|😊|0|10|/blog/my-first-post|0|article|My Publication|2024-10-13 12:00:00|my-first-post|2024-10-13 11:30:00|My First Blog Post
sqlite> 
```