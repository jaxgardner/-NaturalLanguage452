package database

var InsertDepartments = `
INSERT INTO departments (name, head_of_department) VALUES
('Computer Science', NULL),
('Mathematics', NULL),
('Physics', NULL),
('Literature', NULL);
`

var InsertInstrutors = `
INSERT INTO instructors (first_name, last_name, email, phone_number, hire_date) VALUES
('Alice', 'Johnson', 'alice.johnson@university.edu', '123-456-7890', '2015-08-10'),
('Bob', 'Smith', 'bob.smith@university.edu', '123-456-7891', '2017-09-15'),
('Carol', 'Brown', 'carol.brown@university.edu', '123-456-7892', '2016-03-20'),
('David', 'Lee', 'david.lee@university.edu', '123-456-7893', '2018-01-12');
`

var InsertStudents = `
INSERT INTO students (first_name, last_name, email, date_of_birth, enrollment_date, major_id) VALUES
('Eve', 'Taylor', 'eve.taylor@university.edu', '2000-05-21', '2020-09-01', 1),
('Frank', 'Harris', 'frank.harris@university.edu', '1999-11-15', '2019-09-01', 2),
('Grace', 'Walker', 'grace.walker@university.edu', '2001-02-28', '2021-09-01', 3),
('Hank', 'Adams', 'hank.adams@university.edu', '2000-07-10', '2020-09-01', 4);
`

var InsertCourses = `
INSERT INTO courses (code, name, credits, department_id, instructor_id) VALUES
('CS101', 'Introduction to Programming', 4, 1, 1),
('CS102', 'Data Structures', 3, 1, 1),
('MATH101', 'Calculus I', 4, 2, 2),
('PHYS101', 'General Physics I', 4, 3, 3),
('LIT101', 'Introduction to Literature', 3, 4, 4);
`

var InsertEnrollments = `
INSERT INTO enrollments (student_id, course_id, enrollment_date, grade) VALUES
(1, 1, '2020-09-05', 'A'),
(1, 2, '2020-09-05', 'B+'),
(2, 3, '2019-09-10', 'B'),
(3, 4, '2021-09-15', 'A-'),
(4, 5, '2020-09-20', 'B+');
`
