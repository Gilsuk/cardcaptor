CREATE TABLE class (
    class INTEGER PRIMARY KEY,
    slug  TEXT    UNIQUE
                  NOT NULL,
    name  TEXT    UNIQUE
                  NOT NULL
);

CREATE TABLE rarity (
    rarity INTEGER PRIMARY KEY,
    slug   TEXT    UNIQUE
                   NOT NULL,
    name   TEXT    UNIQUE
                   NOT NULL
);

CREATE TABLE keyword (
    keyword INTEGER PRIMARY KEY,
    slug    TEXT    UNIQUE
                    NOT NULL,
    name    TEXT    UNIQUE
                    NOT NULL,
    ref     TEXT,
    text    TEXT
);

CREATE TABLE race (
    race INTEGER PRIMARY KEY,
    slug TEXT    UNIQUE
                 NOT NULL,
    name TEXT    UNIQUE
                 NOT NULL
);

CREATE TABLE type (
    type INTEGER PRIMARY KEY,
    slug TEXT    UNIQUE
                 NOT NULL,
    name TEXT    UNIQUE
                 NOT NULL
);

CREATE TABLE setgroup (
    setgroup INTEGER PRIMARY KEY,
    slug     TEXT    UNIQUE
                     NOT NULL,
    year     INTEGER,
    name     TEXT    UNIQUE
                     NOT NULL,
    standard INTEGER NOT NULL
                     DEFAULT (0) 
                     CHECK (0 OR 
                            1) 
);

CREATE INDEX setyear ON setgroup (
    year DESC
)
WHERE year IS NOT NULL;

CREATE INDEX setstandard ON setgroup (
    standard DESC
);

CREATE TABLE cardSet (
    cardSet       INTEGER PRIMARY KEY,
    name        TEXT    NOT NULL
                        UNIQUE,
    slug        TEXT    UNIQUE
                        NOT NULL,
    releasedate TEXT,
    type        TEXT,
    setgroup    INTEGER REFERENCES setgroup (setgroup) ON DELETE SET NULL
                                                       ON UPDATE CASCADE
);

CREATE INDEX setreleasedate ON cardset (
    releasedate DESC
)
WHERE releasedate IS NOT NULL AND 
      releasedate != '';

CREATE INDEX settype ON cardset (
    type DESC
)
WHERE type IS NOT NULL AND 
      type != '';

CREATE INDEX setsetgroup ON cardset (
    setgroup DESC
)
WHERE setgroup IS NOT NULL;

CREATE TABLE card (
    card        INTEGER PRIMARY KEY,
    slug        TEXT    UNIQUE
                        NOT NULL,
    class       INTEGER REFERENCES class (class) ON DELETE RESTRICT
                                                 ON UPDATE CASCADE
                        NOT NULL,
    type        INTEGER REFERENCES type (type) ON DELETE RESTRICT
                                               ON UPDATE CASCADE
                        NOT NULL,
    cardset     INTEGER REFERENCES cardset (cardset) ON DELETE SET NULL
                                                     ON UPDATE CASCADE,
    rarity      INTEGER REFERENCES rarity (rarity) ON DELETE RESTRICT
                                                   ON UPDATE CASCADE
                        NOT NULL,
    race        INTEGER REFERENCES race (race) ON DELETE RESTRICT
                                               ON UPDATE CASCADE,
    artist      TEXT,
    name        TEXT    NOT NULL,
    text        TEXT,
    flavor      TEXT,
    img         TEXT,
    cropimg     TEXT,
    cost        INTEGER,
    health      INTEGER,
    attack      INTEGER,
    armor       INTEGER,
    arena       INTEGER CHECK (0 OR 
                               1) 
                        NOT NULL
                        DEFAULT (0),
    collectable INTEGER CHECK (0 OR 
                               1) 
                        NOT NULL
                        DEFAULT (0) 
);

CREATE INDEX cardclass ON card (
    class
);

CREATE INDEX cardtype ON card (
    type
);

CREATE INDEX cardcardset ON card (
    cardset
)
WHERE cardset IS NOT NULL;

CREATE INDEX cardrarity ON card (
    rarity
);

CREATE INDEX cardrace ON card (
    race
)
WHERE race IS NOT NULL;

CREATE INDEX cardname ON card (
    name
);

CREATE INDEX cardcost ON card (
    cost
)
WHERE cost IS NOT NULL;

CREATE INDEX cardhealth ON card (
    health
)
WHERE health IS NOT NULL;

CREATE INDEX cardattack ON card (
    attack
)
WHERE attack IS NOT NULL;

CREATE INDEX cardarmor ON card (
    armor
)
WHERE armor IS NOT NULL;

CREATE INDEX cardarena ON card (
    arena DESC
);

CREATE INDEX cardcollectable ON card (
    collectable DESC
);

CREATE TABLE classes (
    card  INTEGER REFERENCES card (card) ON DELETE CASCADE
                                         ON UPDATE CASCADE,
    class INTEGER REFERENCES class (class) ON DELETE RESTRICT
                                           ON UPDATE CASCADE,
    PRIMARY KEY (
        card,
        class
    )
);

CREATE TABLE mechanism (
    card    INTEGER REFERENCES card (card) ON DELETE CASCADE
                                           ON UPDATE CASCADE,
    keyword INTEGER REFERENCES keyword (keyword) ON DELETE RESTRICT
                                                 ON UPDATE CASCADE,
    PRIMARY KEY (
        card,
        keyword
    )
);

CREATE TABLE family (
    child  INTEGER PRIMARY KEY,
    parent INTEGER NOT NULL
);

CREATE INDEX familyparent ON family (
    parent
);
