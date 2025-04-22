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
frontend/
├── components/         # Reusable UI components
├── pages/              # Next.js pages (routes)
│   ├── api/            # Next.js API routes (server-side logic)
├── stores/             # Zustand state management
├── utils/              # Utility functions
├── styles/             # Global and component-specific styles
├── public/             # Static assets
├── hooks/              # Custom React hooks
├── services/           # API service functions
├── middleware/         # Middleware for handling requests
├── next.config.js      # Next.js configuration
├── package.json        # Project dependencies and scripts
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

4. **Custom Hooks**:
    - Reusable hooks for common logic, such as `useAuth`, `useFetch`.

5. **Middleware**:
    - Middleware is used for request handling and authentication checks.

6. **Styling**:
    - Uses CSS modules and global styles for consistent design.

---

## Notes
- Replace `<repository-url>` and `<PORT>` with the actual values for your project.
- Ensure Docker is running before using Docker Compose.  
- For production, update the `.env.local` file with the appropriate backend API URL.  
- Use the existing Postman collection to test the backend APIs.  
