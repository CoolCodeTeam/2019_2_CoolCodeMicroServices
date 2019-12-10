DROP TABLE IF EXISTS messages CASCADE;
CREATE TABLE "messages"(
    ID BIGSERIAL NOT NULL PRIMARY KEY ,
    type SMALLINT NOT NULL, --IN ('TEXT','PHOTO','VOICE')
    body TEXT NOT NULL,
    fileID VARCHAR ,
    fileExtension VARCHAR ,
    chatID BIGINT NOT NULL,
    messageTime timestamp,
    likes BIGINT DEFAULT 0,T
    authorID BIGINT NOT NULL,
    hideForAuthor bool default false

);

