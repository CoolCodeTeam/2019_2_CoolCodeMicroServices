DROP TABLE IF EXISTS workspaces_users CASCADE;
CREATE TABLE workspaces_users
(
    isAdmin     BOOLEAN NOT NULL,
    workspaceID BIGINT  NOT NULL,
    userID      BIGINT  NOT NULL,

    FOREIGN KEY (workspaceID) REFERENCES workspaces (ID) ON DELETE CASCADE
);

