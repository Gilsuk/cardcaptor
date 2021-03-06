CREATE TABLE class (
    class INTEGER PRIMARY KEY,
    slug  TEXT    UNIQUE
                  NOT NULL,
    name  TEXT    UNIQUE
                  NOT NULL,
    card  INTEGER
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
    type        TEXT
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
    durability       INTEGER,
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
    parent INTEGER,
    child  INTEGER,
    PRIMARY KEY (
        parent,
        child
    )
);

CREATE INDEX familychild ON family (
    child
);

CREATE TABLE groups (
    setgroup    INTEGER REFERENCES setgroup (setgroup) ON DELETE RESTRICT
                                                       ON UPDATE CASCADE,
    cardSet     INTEGER REFERENCES cardSet (cardSet) ON DELETE CASCADE
                                                       ON UPDATE CASCADE,
    PRIMARY KEY (
        setgroup,
        cardSet
    )
);

CREATE VIEW vCard AS
    SELECT card.card AS id,
           card.slug AS slug,
           card.name AS name,
           IFNULL(class.name, '') AS class,
           IFNULL(type.name, '') AS type,
           IFNULL(rarity.name, '') AS rarity,
           IFNULL(race.name, '') AS race,
           cost,
           health,
           attack,
           armor,
           durability,
           arena,
           collectable,
           IFNULL(standard, 0) AS standard,
           card.text AS text,
           flavor,
           artist,
           IFNULL(cardset.name, '') AS cardSet,
           img,
           cropImg
      FROM card
           LEFT OUTER JOIN
           CLASS ON card.class = class.class
           LEFT OUTER JOIN
           type ON card.type = type.type
           LEFT OUTER JOIN
           rarity ON card.rarity = rarity.rarity
           LEFT OUTER JOIN
           race ON card.race = race.race
           LEFT OUTER JOIN
           cardset ON card.cardset = cardset.cardSet
           LEFT OUTER JOIN
           (
               SELECT cardSet.cardSet,
                      1 AS standard
                 FROM cardSet
                      INNER JOIN
                      groups ON cardSet.cardSet = groups.cardSet
                      INNER JOIN
                      setGroup ON groups.setgroup = setGroup.setgroup
                WHERE setGroup.setGroup = 7
           )
           AS CS ON card.cardSet = CS.cardSet;