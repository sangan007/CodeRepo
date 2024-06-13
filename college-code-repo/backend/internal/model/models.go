package model

type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    StudentID string `json:"student_id"`
    Password string `json:"password"`
    Role     string `json:"role"`
}

type Repository struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    OwnerID     int    `json:"owner_id"`
    Stars       int    `json:"stars"`
    Issues      int    `json:"issues"`
    PullRequests int   `json:"pull_requests"`
}

type File struct {
    ID           int    `json:"id"`
    Name         string `json:"name"`
    RepositoryID int    `json:"repository_id"`
}

type Submission struct {
    ID             int    `json:"id"`
    RepositoryID   int    `json:"repository_id"`
    UserID         int    `json:"user_id"`
    SubmissionDate string `json:"submission_date"`
}

type Evaluation struct {
    ID           int    `json:"id"`
    TeacherID    int    `json:"teacher_id"`
    SubmissionID int    `json:"submission_id"`
    Grade        int    `json:"grade"`
    Feedback     string `json:"feedback"`
}

type RepositoryPermission struct {
    ID             int    `json:"id"`
    UserID         int    `json:"user_id"`
    RepositoryID   int    `json:"repository_id"`
    PermissionLevel string `json:"permission_level"`
}
