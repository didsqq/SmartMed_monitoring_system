CREATE TABLE doctors
(
    id serial primary key,
    fullname varchar(255) not null,
    date_of_birth date,
    email varchar(255) not null,
    specialization VARCHAR(100),
    password_hash varchar(255) not null
);

CREATE TABLE patients
(
    id serial primary key,
    doctor_id int,
    fullname varchar(255) not null,
    date_of_birth date,
    gender VARCHAR(10),
    address varchar(255),
    email varchar(255) not null,
    phone_number char(10) not null,
    password_hash varchar(255) not null,
    foreign key(doctor_id) references doctors(id) on delete cascade
);

CREATE TABLE analysis
(
    id serial primary key,
    patient_id int,
    pulse smallint, -- пульс
    respiratory_rate smallint, -- частота дыхательных движений
    oxygen_saturation decimal(5, 2), -- сатурация
    systolic_blood_pressure int, -- систолическое давление (Артериальное давление)
    diastolic_blood_pressure int, -- диастолическое давление (Артериальное давление)
    heart_rate INT, -- частота сердечных сокращений из ЭКГ
    analysis_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    foreign key(patient_id) references patients(id) on delete cascade
);
