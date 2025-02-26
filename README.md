# habit-tracker
The Habit Tracker Web App will help users log daily habits, track their progress, and visualize streaks and completion rates. This project is designed to demonstrate both full-stack development and personal productivity enhancement, all while leveraging popular Golang libraries.

## Project Breakdown by Day
### Define Scope & Set Up Project

Goals:

    Define core features and the database schema.
    Set up project structure and initialize the Go module.

Tasks:

    Sketch a rough UI wireframe (habit list view, daily checkboxes, analytics dashboard).
    Define core features:
        CRUD for habits (Add a habit, update daily status, edit habit details, delete a habit).
        Track daily completion and streak counts.
        Display analytics (e.g., "3-day streak", "80% completion rate this week").
    Set up project folder structure:

    /habit-tracker
      ├── /templates        // HTML templates
      ├── /static           // CSS/JS files
      ├── main.go
      ├── models.go
      ├── routes.go
      └── database.go

    Run go mod init habit-tracker and install dependencies (Gin, GORM, godotenv).

📌 Deliverable: Project is initialized, and the scope is clearly defined.

### Set Up Database & Models

Goals:

    Use GORM with SQLite/PostgreSQL to create the database schema for tracking habits.

Tasks:

    Design models for the Habits table with fields such as:
        ID, Name, Description, Frequency (e.g., daily, weekly), CurrentStreak, LastCompletedAt, and TotalCompletions.
    Implement database connection logic in database.go using GORM.
    Write migration functions to create the necessary tables.
    Seed some test data to verify the schema works.

📌 Deliverable: A fully functional database schema set up with GORM.

### Implement Backend Logic (CRUD Operations)

Goals:

    Build the API endpoints for habit management using Gin.

Tasks:

    In routes.go, set up HTTP routes for:
        POST /habits → Create a new habit.
        GET /habits → Retrieve all habits.
        PUT /habits/:id → Update a habit (e.g., mark it as completed, update streaks).
        DELETE /habits/:id → Delete a habit.
    Implement handler functions in main.go or separate controllers.
    Use GORM functions for interacting with the database.
    Validate incoming data (e.g., non-empty habit names, valid frequencies).

📌 Deliverable: Fully functional backend CRUD operations for habit tracking.

### Build the UI & Connect Frontend to Backend

Goals:

    Create HTML templates and design a user-friendly interface using html/template.
    Connect the frontend to the backend API.

Tasks:

    Develop HTML templates for:
        Homepage: Display the list of habits with daily checkboxes and current streaks.
        Add/Edit Habit Page: Forms for adding and editing habits.
        Analytics/Dashboard: A section showing overall progress.
    Use simple CSS (or a lightweight framework like TailwindCSS) for styling.
    Set up routes to serve these templates and pass data from your backend to the UI.
    Implement AJAX or simple form submissions to interact with your API.

📌 Deliverable: A functional and aesthetically pleasing UI that dynamically displays habit data.

### Implement Analytics & Habit Tracking Logic

Goals:

    Enhance the backend to calculate analytics, such as streaks and completion percentages.
    Reflect these analytics in the UI.

Tasks:

    Add logic to update a habit’s streak when marked as completed.
    Write functions to calculate metrics (e.g., consecutive days, weekly completion rate).
    Display these analytics in the dashboard (e.g., “Current Streak: 4 days”, “Week Completion: 80%”).
    Test the calculations with different scenarios to ensure accuracy.

📌 Deliverable: Users can see real-time analytics on their habit progress.

### Testing, Debugging & Final UI Enhancements

Goals:

    Test the full functionality of the Habit Tracker.
    Refine both the backend and the UI for a smooth user experience.

Tasks:

    Manually test all CRUD operations and analytics calculations.
    Debug any issues found (e.g., incorrect streak calculations, UI misalignments).
    Optimize database queries if necessary.
    Finalize CSS styling and ensure the UI is responsive across devices.

📌 Deliverable: A bug-free, polished Habit Tracker with a smooth, responsive interface.

### Final Testing, Documentation & Deployment

Goals:

    Conduct final end-to-end testing.
    Polish documentation and prepare for deployment.
    Deploy the app using platforms like Heroku or Railway.

Tasks:

    Perform comprehensive testing of the entire app flow.
    Write or update documentation (README, API docs, deployment instructions).
    Set up environment variables using godotenv.
    Deploy the app on Railway, Heroku, or your preferred platform.
    Record a short demo or take screenshots for content creation.

📌 Deliverable: The Habit Tracker is fully deployed, documented, and ready to be showcased.
