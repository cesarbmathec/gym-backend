-- SQL schema for database named 'gym_backend'
-- Table: clients
CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    identity_card VARCHAR(15) UNIQUE NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    weight DECIMAL(5,2),
    height DECIMAL(3,2),
    blood_type VARCHAR(5),
    medical_conditions TEXT,
    birth_date DATE,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: trainers
CREATE TABLE trainers (
    id_trainer SERIAL PRIMARY KEY,
    identity_card VARCHAR(15) UNIQUE NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    trainer_type VARCHAR(20) CHECK (trainer_type IN ('general', 'personal')) DEFAULT 'general',
    is_active BOOLEAN DEFAULT true
);

-- Table: subscriptions
CREATE TABLE subscriptions (
    id_subscription SERIAL PRIMARY KEY,
    client_id INT REFERENCES clients(id) ON DELETE CASCADE,
    training_type VARCHAR(20) CHECK (training_type IN ('gym_normal', 'crossfit')),
    personal_trainer_id INT REFERENCES trainers(id_trainer),
    start_date DATE NOT NULL,
    expiry_date DATE NOT NULL,
    monthly_fee DECIMAL(10,2) NOT NULL,
    status VARCHAR(20) CHECK (status IN ('active', 'overdue', 'restricted')) DEFAULT 'active',
    UNIQUE(client_id)
);

-- Table: payments
CREATE TABLE payments (
    id_payment SERIAL PRIMARY KEY,
    subscription_id INT REFERENCES subscriptions(id_subscription) ON DELETE CASCADE,
    amount DECIMAL(10,2) NOT NULL,
    payment_date TIMESTAMP NOT NULL,
    payment_method VARCHAR(20) CHECK (payment_method IN ('cash', 'zelle', 'usdt', 'airtm', 'mobile_payment', 'bank_transfer'))
);

-- Table: check_ins
CREATE TABLE check_ins (
    id_check_in SERIAL PRIMARY KEY,
    client_id INT REFERENCES clients(id) ON DELETE CASCADE,
    check_in_time TIMESTAMP NOT NULL,
    trainer_id INT REFERENCES trainers(id_trainer)
);

-- Table: crossfit_performance
CREATE TABLE crossfit_performance (
    id_performance SERIAL PRIMARY KEY,
    client_id INT REFERENCES clients(id) ON DELETE CASCADE,
    wod_name VARCHAR(100),
    reps INT,
    time_spent INTERVAL,
    performance_date DATE NOT NULL,
    score DECIMAL(6,2)
);

-- Table: bodybuilding_performance
CREATE TABLE bodybuilding_performance (
    id_performance SERIAL PRIMARY KEY,
    client_id INT REFERENCES clients(id) ON DELETE CASCADE,
    weight DECIMAL(5,2),
    body_fat_percentage DECIMAL(4,2),
    arm_measurement DECIMAL(4,2),
    measurement_date DATE NOT NULL
);

-- Table: chats
CREATE TABLE chats (
    id_chat SERIAL PRIMARY KEY,
    chat_name VARCHAR(100) NOT NULL,
    description TEXT,
    is_group_chat BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: chat_participants
CREATE TABLE chat_participants (
    id_participant SERIAL PRIMARY KEY,
    chat_id INT REFERENCES chats(id_chat) ON DELETE CASCADE,
    user_id INT NOT NULL,
    is_chat_admin BOOLEAN DEFAULT false,
    user_type VARCHAR(20) CHECK (user_type IN ('client', 'trainer'))
);

-- Table: chat_messages
CREATE TABLE chat_messages (
    id_message SERIAL PRIMARY KEY,
    chat_id INT REFERENCES chats(id_chat) ON DELETE CASCADE,
    user_id INT NOT NULL,
    user_type VARCHAR(20) CHECK (user_type IN ('client', 'trainer')),
    message TEXT NOT NULL,
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_read BOOLEAN DEFAULT false
);

-- Table: notifications
CREATE TABLE notifications (
    id_notification SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    message TEXT NOT NULL,
    notification_type VARCHAR(30) CHECK (notification_type IN ('event', 'competition', 'payment_due', 'overdue', 'payment_change')),
    client_id INT REFERENCES clients(id) ON DELETE SET NULL,
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_read BOOLEAN DEFAULT false
);

-- Performance indexes
CREATE INDEX idx_clients_identity_card ON clients(identity_card);
CREATE INDEX idx_subscriptions_status ON subscriptions(status);
CREATE INDEX idx_payments_date ON payments(payment_date);
CREATE INDEX idx_chat_messages_chat_time ON chat_messages(chat_id, sent_at);
