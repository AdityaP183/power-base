CREATE TABLE
    heroes (
        id VARCHAR(36) PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        -- Appearance
        appearance_eye_color VARCHAR(100),
        appearance_gender VARCHAR(50),
        appearance_hair_color VARCHAR(100),
        appearance_height TEXT,
        appearance_race VARCHAR(100),
        appearance_weight TEXT,
        -- Biography
        biography_aliases TEXT,
        biography_alignment VARCHAR(50),
        biography_alter_egos TEXT,
        biography_first_appearance TEXT,
        biography_full_name VARCHAR(255),
        biography_place_of_birth TEXT,
        biography_publisher VARCHAR(100),
        -- Connections
        connections_group_affiliation TEXT,
        connections_relatives TEXT,
        -- Powerstats
        powerstats_combat VARCHAR(10),
        powerstats_durability VARCHAR(10),
        powerstats_intelligence VARCHAR(10),
        powerstats_power VARCHAR(10),
        powerstats_speed VARCHAR(10),
        powerstats_strength VARCHAR(10),
        -- Work
        work_base TEXT,
        work_occupation TEXT,
        -- Image
        image_url TEXT,
        response VARCHAR(50) DEFAULT 'success',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );
