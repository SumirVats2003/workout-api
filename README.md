# Workout API

## Tech Stack Used
- Database: PostgreSQL
- Programming Language: Golang

## Overview

This document outlines the database schema for a fitness tracking application. The design supports basic user functionality such as exercise cataloging, workout logging, and tracking of workout plans. The schema is modular and can be extended as new features are introduced.

## Entity-Relationship Summary

The core schema includes the following relations:

1. `users` – Stores user account details.
2. `exercises` – Maintains a catalog of exercises.
3. `workouts` – Logs workout entries consisting of one or more exercises.
4. `plans` – Tracks planned or completed workout sessions.

## 1. `users` Table

Stores authentication and identity information for registered users.

| Column Name | Data Type | Constraints      | Description                    |
| ----------- | --------- | ---------------- | ------------------------------ |
| `userId`    | UUID      | PRIMARY KEY      | Unique identifier for the user |
| `name`      | TEXT      | NOT NULL         | Full name of the user          |
| `email`     | TEXT      | UNIQUE, NOT NULL | Email address (used for login) |
| `password`  | TEXT      | NOT NULL         | Hashed password string         |

**Notes:**

- Passwords must be securely hashed using a modern algorithm (e.g., bcrypt).
- Consider adding `created_at` and `updated_at` timestamps for auditing.

## 2. `exercises` Table

Contains metadata about each exercise available in the system.

| Column Name   | Data Type | Constraints | Description                               |
| ------------- | --------- | ----------- | ----------------------------------------- |
| `exerciseId`  | UUID      | PRIMARY KEY | Unique identifier for the exercise        |
| `name`        | TEXT      | NOT NULL    | Name of the exercise                      |
| `category`    | TEXT      |             | Exercise type (e.g., strength, cardio)    |
| `muscleGroup` | TEXT      |             | Targeted muscle group (e.g., chest, legs) |

**Notes:**

- Optionally, normalize `category` and `muscleGroup` into separate tables to ensure data consistency and enable filtering or grouping.

## 3. `workouts` Table

Stores individual workout records associated with exercises.

| Column Name  | Data Type | Constraints | Description                             |
| ------------ | --------- | ----------- | --------------------------------------- |
| `workoutId`  | UUID      | PRIMARY KEY | Unique identifier for the workout entry |
| `exerciseId` | UUID      | FOREIGN KEY | References `exercises(exerciseId)`      |
| `planId`     | UUID      | FOREIGN KEY | References `plans(planId)`              |
| `reps`       | INT       | NOT NULL    | Number of repetitions per set           |
| `sets`       | INT       | NOT NULL    | Number of sets                          |
| `weight`     | NUMERIC   |             | Weight used (in kg or lbs)              |

**Notes:**

- Consider tracking more granular details (e.g., rest time, tempo) in future iterations.
   
- Index `exerciseId` for faster queries and joins.

## 4. `plans` Table

Associates users with scheduled or completed workout routines.

| Column Name | Data Type   | Constraints               | Description                                |
| ----------- | ----------- | ------------------------- | ------------------------------------------ |
| `planId`    | UUID        | PRIMARY KEY               | Unique identifier for the plan             |
| `userId`    | UUID        | FOREIGN KEY               | References `users(userId)`                 |
| `timestamp` | TIMESTAMPTZ | DEFAULT CURRENT_TIMESTAMP | Date and time the plan was created or used |
| `status`    | STRING      | Completed/Planned         | Stores the completion status of the plan   |

## API Design

### **Authentication Endpoints**

- `POST /login`    
    - Authenticates a user and returns an access token upon successful login.

- `POST /signup`
    - Registers a new user account with required credentials.

### **Exercise Management**

- `POST /exercises/create`
    - Creates a new exercise entry in the system.

### **Workout Plan Management**

- `GET /user/:userId/plan/:planId`
    - Retrieves a specific workout plan associated with a user.
   
- `POST /user/:userId/plan/create`
    - Creates a new workout plan for the specified user.

- `POST /user/:userId/plan/:planId/schedule`
    - Schedules the specified workout plan for a given date or time.
   
- `PUT /user/:userId/plan/:planId`
    - Updates the details of an existing workout plan.
  
- `DELETE /user/:userId/plan/:planId`
    - Deletes a workout plan from the user's profile.

- `GET /user/:userId/all-plans`
    - Retrieves all workout plans associated with the user.

### **Reporting**

- `GET /user/:userId/report`
    - Generates a workout or performance report for the specified user.

## General Recommendations

- Use UUIDs for all primary keys to ensure uniqueness across distributed systems and for better security.
   
- Index all foreign key columns to improve performance during joins and lookups.
   
- Maintain a consistent naming convention for all tables and fields (e.g., snake_case).
   
- Consider adding auditing columns (`created_at`, `updated_at`, `deleted_at`) for record lifecycle management.
   
- Future enhancements may include user roles, exercise history, workout templates, and progress visualizations.
