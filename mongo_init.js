db.createUser(
        {
            user: "user",
            pwd: "pass",
            roles: [
                {
                    role: "readWrite",
                    db: "db-rifa"
                }
            ]
        }
);