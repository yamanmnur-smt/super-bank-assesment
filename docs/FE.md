# Frontend Setup Guide

This guide explains how to set up the frontend locally and using Docker Compose.

## Prerequisites
- **Node.js** (version 18 or higher)
- **npm** or **yarn**
- **Next.js** (latest)
- **Zustand** (for state management)
- **Docker** and **Docker Compose**

---

## Local Setup

1. **Clone the Repository**:
    ```bash
    git clone git@github.com:yamanmnur/super-bank-assesment.git
    cd super-bank-assesment/frontend
    ```

2. **Install Dependencies**:
    ```bash
    npm install
    # or
    yarn install
    ```

3. **Set Up Environment Variables**:
    - Copy `.env.example` to `.env.local`:
      ```bash
      cp .env.example .env.local
      ```
    - Update the `.env.local` file with your local configuration.

4. **Start the Development Server**:
    ```bash
    npm run dev
    # or
    yarn dev
    ```

5. **Access the Application**:
    - Open your browser and navigate to `http://localhost:3000`.

---

## Docker Compose Setup

1. **Clone the Repository**:
    ```bash
    git clone git@github.com:yamanmnur/super-bank-assesment.git
    cd super-bank-assesment
    ```

2. **Set Up Environment Variables**:
    - Copy `.env.example` to `.env.local`:
      ```bash
      cd frontend
      cp .env.example .env.local
      ```
    - Update the `.env.local` file with your Docker configuration.

3. **Start Services**:
    ```bash
    docker-compose up --build -d
    ```

4. **Access the Application**:
    - Open your browser and navigate to `http://localhost:3000`.

---

# Frontend Structure

The frontend is built using **Next.js** and follows a modular structure. Below is an explanation of the structure:

### Folder Structure
```
src/
├── app/                # Application structure
│   ├── _components/    # Shared reusable UI components
│   ├── (auth)/         # Authentication-related pages and components
│   │   ├── login/      # Login page and related components
│   ├── (main)/         # Main application pages and components
│   │   ├── _components/ # Shared components for main pages
│   │   ├── dashboard/  # Dashboard page and related components
│   │   ├── customer/   # Customer page and related components
│   ├── api/            # API routes for server-side logic
│   │   ├── auth/       # Authentication-related API routes
│   │   ├── dashboard/  # Dashboard-related API routes
│   │   ├── customer/   # Customer-related API routes
├── assets/             # Static assets (images, fonts, etc.)
├── context/            # Context providers for global state management
```

### Explanation of Key Features

1. **State Management**:
    - Uses **Zustand** for managing global state.
    - Example: Authentication state, user data.

2. **JWT Token Handling**:
    - JWT tokens are stored in cookies using Next.js API routes.
    - The frontend fetches data from the backend via the Next.js API, ensuring secure token handling.

3. **API Services**:
    - Centralized API service functions are defined in the `services/` folder.
    - Example: `authService`, `userService`.

4. **Middleware**:
    - Middleware is used for request handling and authentication checks.

5. **Styling**:
    - Uses CSS modules and global styles for consistent design.

---

## Notes
- Replace `<repository-url>` and `<PORT>` with the actual values for your project.
- Ensure Docker is running before using Docker Compose.  
- For production, update the `.env.local` file with the appropriate backend API URL.  
- Use the existing Postman collection to test the backend APIs.  
