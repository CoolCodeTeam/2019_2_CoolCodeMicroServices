DROP TABLE IF EXISTS STICKERS_USERS CASCADE;
CREATE TABLE stickers_users(
    stickerpackID BIGINT,
    userID BIGINT,

    FOREIGN KEY(stickerpackID) references stickers(stickerpackID) on delete CASCADE,
    FOREIGN KEY(userID) references users(ID) on delete CASCADE
)