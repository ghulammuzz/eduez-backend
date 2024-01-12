import sqlite3

# Membuat koneksi ke database (jika tidak ada, maka akan dibuat)
conn = sqlite3.connect('eduze.db')

# Membuat objek cursor
cursor = conn.cursor()

# Membuat tabel Course
cursor.execute('''
    CREATE TABLE IF NOT EXISTS course (
        id TEXT PRIMARY KEY,
        user_id TEXT,
        prompt TEXT,
        title TEXT,
        type_activity TEXT,
        theme_activity TEXT,
        desciption TEXT,
        is_done BOOLEAN DEFAULT 0,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        avg_rating FLOAT DEFAULT 0
    )
''')

# Membuat tabel Subtitle
cursor.execute('''
    CREATE TABLE IF NOT EXISTS subtitle (
        id TEXT PRIMARY KEY,
        topic TEXT,
        shortdesc TEXT,
        course_id TEXT,
        is_done BOOLEAN DEFAULT 0,
        FOREIGN KEY (course_id) REFERENCES course(id) ON DELETE CASCADE
    )
''')

# Membuat tabel Content
cursor.execute('''
    CREATE TABLE IF NOT EXISTS content (
        id TEXT PRIMARY KEY,
        opening TEXT,
        closing TEXT,
        subtitle_id TEXT,
        FOREIGN KEY (subtitle_id) REFERENCES subtitle(id) ON DELETE CASCADE
    )
''')

# Membuat tabel Steps
cursor.execute('''
    CREATE TABLE IF NOT EXISTS steps (
        id TEXT PRIMARY KEY,
        texts TEXT,
        content_id TEXT,
        FOREIGN KEY (content_id) REFERENCES content(id) ON DELETE CASCADE
    )
''')

# Menyimpan perubahan ke database
conn.commit()

# Menutup koneksi
conn.close()
